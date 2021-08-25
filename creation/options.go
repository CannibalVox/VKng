package creation

import (
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type Options interface {
	AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error)
}