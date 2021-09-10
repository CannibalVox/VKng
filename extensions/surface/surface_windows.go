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

func CreateSurface(allocator cgoalloc.Allocator, createInfo unsafe.Pointer, instance *resource.Instance) (*Surface, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	var surface C.VkSurfaceKHR
	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))

	createSurfaceFunc := (C.PFN_vkCreateWin32SurfaceKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkCreateWin32SurfaceKHR"))))

	res := loader.VkResult(C.cgoCreateWin32SurfaceKHR(createSurfaceFunc, instanceHandle, (*C.VkWin32SurfaceCreateInfoKHR)(createInfo), nil, &surface))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return buildSurface(arena, instance, surface), res, nil
}
