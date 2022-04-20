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

type DeviceGroupSwapchainCreateOptions struct {
	Modes DeviceGroupPresentModeFlags

	common.HaveNext
}

func (o DeviceGroupSwapchainCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceGroupSwapchainCreateInfoKHR)
	}

	info := (*C.VkDeviceGroupSwapchainCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
	info.pNext = next
	info.modes = C.VkDeviceGroupPresentModeFlagsKHR(o.Modes)

	return preallocatedPointer, nil
}

func (o DeviceGroupSwapchainCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceGroupSwapchainCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
