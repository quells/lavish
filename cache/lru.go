package cache

import (
	"github.com/dop251/goja"
	lru "github.com/hashicorp/golang-lru/v2"
)

type LRU struct {
	programCache *lru.Cache[uint64, *goja.Program]
	renderCache  *lru.Cache[renderHashKey, string]
}

func NewLRU(size int) LRU {
	programCache, err := lru.New[uint64, *goja.Program](size)
	if err != nil {
		panic(err)
	}
	renderCache, err := lru.New[renderHashKey, string](size)
	if err != nil {
		panic(err)
	}
	return LRU{
		programCache: programCache,
		renderCache:  renderCache,
	}
}

func (c LRU) AddProgram(k uint64, program *goja.Program) {
	c.programCache.Add(k, program)
}

func (c LRU) GetProgram(k uint64) *goja.Program {
	if cached, ok := c.programCache.Get(k); ok {
		return cached
	}
	return nil
}

func (c LRU) AddRender(program uint64, data uint64, rendered string) {
	c.renderCache.Add(renderHashKey{program, data}, rendered)
}

func (c LRU) GetRender(program uint64, data uint64) (rendered string, ok bool) {
	var cached string
	if cached, ok = c.renderCache.Get(renderHashKey{program, data}); ok {
		rendered = cached
		return
	}
	return
}
