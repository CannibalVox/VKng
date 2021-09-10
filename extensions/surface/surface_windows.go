//go:build windows
// +build windows

package ext_surface

/*
#include <stdlib.h>
#include <windows.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_win32.h"

VkResult cgoCreateWin32SurfaceKHR(PFN_vkCreateWin32SurfaceKHR fn, VkInstance instance, VkWin32SurfaceCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return fn(instance, pCreateInfo, pAllocator, pSurface);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func CreateSurface(allocator cgoalloc.Allocator, surfacePtr unsafe.Pointer, instance *resource.Instance) (*Surface, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	return buildSurface(arena, instance, (C.VkSurfaceKHR)(surfacePtr)), loader.VKSuccess, nil
}
