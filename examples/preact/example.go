package main

import (
	_ "embed"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/quells/lavish"
	"github.com/quells/lavish/engine/preact10"
	"hash/fnv"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	addr = flag.String("addr", ":3000", "Address to serve HTTP on")

	//go:embed view/app.jsx
	appJSX string
	//go:embed view/index.jsx
	indexJSX string
)

type ExampleData struct {
	Items []string
}

func (d ExampleData) Hash64() uint64 {
	h := fnv.New64a()
	for _, item := range d.Items {
		_, _ = h.Write([]byte(item))
	}
	return binary.BigEndian.Uint64(h.Sum(nil))
}

func main() {
	renderer := preact10.RenderEngine
	b := lavish.NewBundle(
		renderer,
		lavish.ComponentJSX("app.jsx", appJSX, renderer.GetJSXOptions()),
	)

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		data := ExampleData{Items: strings.Split("Hello from Lavish", " ")}
		rendered, err := b.RenderJSX("index.jsx", indexJSX, data)
		elapsed := time.Since(start)

		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err.Error())
			return
		}

		rw.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprint(rw, rendered)

		log.Printf("rendered %d bytes in %s", len(rendered), elapsed)
	})

	log.Printf("Listening at %s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		panic(err)
	}
}
