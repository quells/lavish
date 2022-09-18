package lavish

import (
	"encoding/binary"
	"github.com/dop251/goja"
	"hash/fnv"
)

type ProgramCache interface {
	AddProgram(k uint64, v *goja.Program)
	GetProgram(k uint64) (cached *goja.Program)
}

func hash(src string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(src)) // fnv.sum64a.Write cannot error
	return binary.BigEndian.Uint64(h.Sum(nil))
}

type Hashable interface {
	Hash64() uint64
}

type RenderCache interface {
	AddRender(program uint64, data uint64, v string)
	GetRender(program uint64, data uint64) (cached string, ok bool)
}
