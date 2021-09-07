//go:build darwin
// +build darwin

package ext_surface_sdl2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_macos.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

func (o *CreationOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkMacOSSurfaceCreateInfoMVK)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkMacOSSurfaceCreateInfoMVK{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
	createInfo.flags = 0
	createInfo.pView = sdl.Metal_CreateView(o.Window)

	var next unsafe.Pointer

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), err
}
