package preact10

//go:generate npm install
//go:generate go run github.com/evanw/esbuild/cmd/esbuild --bundle --minify --outfile=preact.bundle.js preact-glue.js

import (
	_ "embed"
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
)

var (
	//go:embed preact.bundle.js
	preactBundleJS string

	preactProgram *goja.Program
)

func init() {
	preactProgram = goja.MustCompile("preact.bundle.js", preactBundleJS, false)
}

type RenderEngine struct{}

// Load the Preact render bundle into the runtime.
func (_ RenderEngine) Load(vm *goja.Runtime) error {
	_, err := vm.RunProgram(preactProgram)
	return err
}

func (_ RenderEngine) GetJSXOptions() api.TransformOptions {
	return api.TransformOptions{
		Loader:      api.LoaderJSX,
		JSXFactory:  "h",
		JSXFragment: "Fragment",
	}
}

func (_ RenderEngine) GetRenderFunction(vm *goja.Runtime) (goja.Callable, error) {
	render, ok := goja.AssertFunction(vm.Get("_preact10_render"))
	if !ok {
		return nil, fmt.Errorf("failed to find _preact10_render function in global scope")
	}
	return render, nil
}
