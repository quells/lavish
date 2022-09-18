package cache

import (
	"github.com/dop251/goja"
	lru "github.com/hashicorp/golang-lru"
)

type LRU struct {
	cache *lru.Cache
}

func NewLRU(size int) LRU {
	cache, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return LRU{cache: cache}
}

func (c LRU) AddProgram(k uint64, program *goja.Program) {
	c.cache.Add(k, program)
}

func (c LRU) GetProgram(k uint64) *goja.Program {
	if cached, ok := c.cache.Get(k); ok {
		return cached.(*goja.Program)
	}
	return nil
}

func (c LRU) AddRender(program uint64, data uint64, rendered string) {
	c.cache.Add(renderHashKey{program, data}, rendered)
}

func (c LRU) GetRender(program uint64, data uint64) (rendered string, ok bool) {
	var cached interface{}
	if cached, ok = c.cache.Get(renderHashKey{program, data}); ok {
		rendered = cached.(string)
		return
	}
	return
}
