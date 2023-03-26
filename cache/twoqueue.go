package cache

import (
	"github.com/dop251/goja"
	lru "github.com/hashicorp/golang-lru/v2"
)

type TwoQueue struct {
	programCache *lru.TwoQueueCache[uint64, *goja.Program]
	renderCache  *lru.TwoQueueCache[renderHashKey, string]
}

func NewTwoQueue(size int) TwoQueue {
	programCache, err := lru.New2Q[uint64, *goja.Program](size)
	if err != nil {
		panic(err)
	}
	renderCache, err := lru.New2Q[renderHashKey, string](size)
	if err != nil {
		panic(err)
	}
	return TwoQueue{
		programCache: programCache,
		renderCache:  renderCache,
	}
}

func (c TwoQueue) AddProgram(k uint64, program *goja.Program) {
	c.programCache.Add(k, program)
}

func (c TwoQueue) GetProgram(k uint64) *goja.Program {
	if cached, ok := c.programCache.Get(k); ok {
		return cached
	}
	return nil
}

func (c TwoQueue) AddRender(program uint64, data uint64, rendered string) {
	c.renderCache.Add(renderHashKey{program, data}, rendered)
}

func (c TwoQueue) GetRender(program uint64, data uint64) (rendered string, ok bool) {
	var cached string
	if cached, ok = c.renderCache.Get(renderHashKey{program, data}); ok {
		rendered = cached
		return
	}
	return
}
