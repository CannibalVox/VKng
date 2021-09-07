//go:build darwin
// +build darwin

package ext_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_macos.h"

VkResult vkCreateMacOSSurfaceMVK(VkInstance instance, const VkMacOSSurfaceCreateInfoMVK* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	PFN_vkCreateMacOSSurfaceMVK f = (PFN_vkCreateMacOSSurfaceMVK)(vkGetInstanceProcAddr((VkInstance)instance, "vkCreateMacOSSurfaceMVK");
	return f(instance, pCreateInfo, pAllocator, pSurface);
}
*/
import "C"

func CreateSurface(createInfo unsafe.Pointer, instance *objects.Instance) (*Surface, core.Result, error) {
	var surface C.VkSurfaceKHR
	instanceHandle := (C.VkInstance)(instance.Handle())

	res := core.Result(C.vkCreateMacOSSurfaceMVK(instance, createInfo, nil, &surface))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Surface{handle: surface, instance: instanceHandle}, res, nil
}
