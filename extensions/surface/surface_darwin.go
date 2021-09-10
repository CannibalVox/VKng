//go:build darwin
// +build darwin

package ext_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_macos.h"

VkResult cgoCreateMacOSSurfaceMVK(PFN_vkCreateMacOSSurfaceMVK fn, VkInstance instance, VkMacOSSurfaceCreateInfoMVK* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	return fn(instance, pCreateInfo, pAllocator, pSurface);
}
*/
import "C"
import "github.com/CannibalVox/cgoalloc"

func CreateSurface(allocator cgoalloc.Allocator, createInfo unsafe.Pointer, instance *objects.Instance) (*Surface, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	var surface C.VkSurfaceKHR
	instanceHandle := (C.VkInstance)(instance.Handle())

	createSurfaceFunc := (C.PFN_vkCreateMacOSSurfaceMVK)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkCreateMacOSSurfaceMVK"))))

	res := core.Result(C.cgoCreateMacOSSurfaceMVK(createSurfaceFunc, instance, createInfo, nil, &surface))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return buildSurface(arena, instance, surface), res, nil
}
