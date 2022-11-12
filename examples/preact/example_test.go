package main

import (
	"github.com/quells/lavish"
	"github.com/quells/lavish/cache"
	"github.com/quells/lavish/engine/preact10"
	"strings"
	"testing"
)

const expect = `<html lang="en">` +
	`<head><meta charset="UTF-8" />` +
	`<meta http-equiv="X-UA-Compatible" content="IE=edge" />` +
	`<meta name="viewport" content="width=device-width, initial-scale=1.0" />` +
	`<title>Preact Example</title>` +
	`</head>` +
	`<body>` +
	`<h1>Example</h1>` +
	`<h4>for Preact 10</h4>` +
	`<ul>` +
	`<li>Hello</li>` +
	`<li>from</li>` +
	`<li>Lavish</li>` +
	`</ul>` +
	`</body>` +
	`</html>`

func TestRender(t *testing.T) {
	renderer := preact10.RenderEngine
	data := ExampleData{Items: strings.Split("Hello from Lavish", " ")}
	bundle := lavish.NewBundle(
		renderer,
		lavish.ComponentJSX("app.jsx", appJSX, renderer.GetJSXOptions()),
	)

	got, err := bundle.RenderJSX("index.jsx", indexJSX, data)
	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Fatalf("got unexpected output: %s", got)
	}
}

func BenchmarkRender(b *testing.B) {
	renderer := preact10.RenderEngine
	data := ExampleData{Items: strings.Split("Hello from Lavish", " ")}
	bundle := lavish.NewBundle(
		renderer,
		lavish.ComponentJSX("app.jsx", appJSX, renderer.GetJSXOptions()),
	)

	tests := []struct {
		name   string
		bundle lavish.Bundle
	}{
		{
			name:   "uncached",
			bundle: bundle,
		},
		{
			name:   "program cache",
			bundle: bundle.WithProgramCache(cache.NewLRU(1)),
		},
		{
			name:   "render cache",
			bundle: bundle.WithProgramCache(cache.NewLRU(1)).WithRenderCache(cache.NewTwoQueue(16)),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				got, err := tt.bundle.RenderJSX("index.jsx", indexJSX, data)
				if err != nil {
					b.Fatal(err)
				}
				if got != expect {
					b.Fatalf("got unexpected output: %s", got)
				}
			}
		})
	}
}
