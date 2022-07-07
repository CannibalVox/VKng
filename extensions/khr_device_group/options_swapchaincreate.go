package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceGroupSwapchainCreateInfo struct {
	Modes DeviceGroupPresentModeFlags

	common.NextOptions
}

func (o DeviceGroupSwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceGroupSwapchainCreateInfoKHR)
	}

	info := (*C.VkDeviceGroupSwapchainCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
	info.pNext = next
	info.modes = C.VkDeviceGroupPresentModeFlagsKHR(o.Modes)

	return preallocatedPointer, nil
}
