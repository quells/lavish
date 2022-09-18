// Package preact10 is the public interface for the Preact v10.x render engine.
package preact10

import "github.com/quells/lavish/internal/preact10"

type renderEngine struct {
	preact10.RenderEngine
}

// A RenderEngine bundled with Preact v10.11.0.
var RenderEngine renderEngine
