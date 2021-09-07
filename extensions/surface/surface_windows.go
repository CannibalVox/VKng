//go:build windows
// +build windows

package ext_surface

/*
#include <stdlib.h>
#include <windows.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_win32.h"

VkResult vkCreateWin32SurfaceKHR(VkInstance instance, const VkWin32SurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	PFN_vkCreateWin32SurfaceKHR f = (PFN_vkCreateWin32SurfaceKHR)vkGetInstanceProcAddr((VkInstance)instance, "vkCreateWin32SurfaceKHR");
	return f(instance, pCreateInfo, pAllocator, pSurface);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resource"
	"unsafe"
)

func CreateSurface(createInfo unsafe.Pointer, instance *resource.Instance) (*Surface, core.Result, error) {
	var surface C.VkSurfaceKHR
	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))

	res := core.Result(C.vkCreateWin32SurfaceKHR(instanceHandle, (*C.VkWin32SurfaceCreateInfoKHR)(createInfo), nil, &surface))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Surface{handle: surface, instance: instanceHandle}, res, nil
}
