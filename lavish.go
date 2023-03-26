package lavish

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
)

type TransformError struct {
	Name     string
	Errors   []api.Message
	Warnings []api.Message
}

func (err TransformError) Error() string {
	b := new(strings.Builder)
	_, _ = fmt.Fprintf(b, "failed to transform %s: ", err.Name)

	idx := 0
	for _, e := range err.Errors {
		if idx > 0 {
			_, _ = fmt.Fprintf(b, ", ")
		}
		if loc := e.Location; loc != nil {
			_, _ = fmt.Fprintf(b, "[%s %d:%d] ", err.Name, loc.Line, loc.Column)
		}
		_, _ = fmt.Fprint(b, e.Text)
		idx++
	}
	for _, w := range err.Warnings {
		if idx > 0 {
			_, _ = fmt.Fprintf(b, ", ")
		}
		if loc := w.Location; loc != nil {
			_, _ = fmt.Fprintf(b, "[%s %d:%d] ", err.Name, loc.Line, loc.Column)
		}
		_, _ = fmt.Fprint(b, w.Text)
		idx++
	}

	return b.String()
}

// CompileJSX to a goja Program.
func CompileJSX(name, jsx string, options api.TransformOptions) (program *goja.Program, err error) {
	// Transpile JSX to vanilla JS
	transformed := api.Transform(jsx, options)
	if len(transformed.Errors) > 0 {
		err = TransformError{
			Name:     name,
			Errors:   transformed.Errors,
			Warnings: transformed.Warnings,
		}
		return
	}

	const isStrict = false
	program, err = goja.Compile(name, string(transformed.Code), isStrict)
	if err != nil {
		err = fmt.Errorf("failed to compile jsx: %w", err)
		return
	}

	return
}

// A Loader loads a module into a runtime environment.
type Loader interface {
	// Load module into runtime environment.
	Load(*goja.Runtime) error
}

type LoaderFunc func(*goja.Runtime) error

func (f LoaderFunc) Load(vm *goja.Runtime) error {
	return f(vm)
}

func ComponentJSX(name, source string, options api.TransformOptions) Loader {
	program, compileErr := CompileJSX(name, source, options)

	return LoaderFunc(func(vm *goja.Runtime) (err error) {
		if compileErr != nil {
			return fmt.Errorf("failed to compile component %q: %w", name, compileErr)
		}

		if _, err = vm.RunProgram(program); err != nil {
			return fmt.Errorf("failed to run compiled component %q: %w", name, err)
		}

		return
	})
}

type RenderEngine interface {
	Loader

	// GetJSXOptions to transpile to vanilla JS.
	GetJSXOptions() api.TransformOptions

	// GetRenderFunction after it has been loaded into the runtime environment.
	// This function should take exactly 1 argument and produce a string.
	GetRenderFunction(*goja.Runtime) (goja.Callable, error)
}
