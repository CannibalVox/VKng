package VKng

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Allocator interface {
	Malloc(uint) unsafe.Pointer
	Free(unsafe.Pointer)
	CString(string) unsafe.Pointer
	CBytes([]byte) unsafe.Pointer
	GoString(unsafe.Pointer) string
	GoStringN(unsafe.Pointer, int) string
	GoBytes(unsafe.Pointer, int) []byte
}

type DefaultAllocator struct {}

func (a *DefaultAllocator) Malloc(size uint) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func (a *DefaultAllocator) Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func (a *DefaultAllocator) CString(str string) unsafe.Pointer {
	return unsafe.Pointer(C.CString(str))
}

func (a *DefaultAllocator) CBytes(b []byte) unsafe.Pointer {
	return C.CBytes(b)
}

func (a *DefaultAllocator) GoString(ptr unsafe.Pointer) string {
	return C.GoString((*C.char)(ptr))
}

func (a *DefaultAllocator) GoStringN(ptr unsafe.Pointer, l int) string {
	return C.GoStringN((*C.char)(ptr), C.int(l))
}

func (a *DefaultAllocator) GoBytes(b unsafe.Pointer, l int) []byte {
	return C.GoBytes(b, C.int(l))
}
