// +build windows

package ext_surface

/*
#include <stdlib.h>
#include <windows.h>
#include "vulkan/vulkan.h"
#include "vulkan/vulkan_win32.h"

VkResult vkCreateWin32SurfaceKHR(VkInstance instance, const VkWin32SurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface) {
	PFN_vkCreateWin32SurfaceKHR f = (PFN_vkCreateWin32SurfaceKHR)vkGetInstanceProcAddr((VkInstance)instance, "vkCreateWin32SurfaceKHR");
	return f(instance, pCreateInfo, pAllocator, pSurface);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/objects"
	"unsafe"
)

func CreateSurface(createInfo unsafe.Pointer, instance *objects.Instance) (*Surface, error) {
	var surface C.VkSurfaceKHR
	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))

	res := C.vkCreateWin32SurfaceKHR(instanceHandle, (*C.VkWin32SurfaceCreateInfoKHR)(createInfo), nil, &surface)
	err := VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Surface{handle:surface, instance:instanceHandle}, nil
}