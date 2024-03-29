package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ImageSwapchainCreateInfo struct {
	Swapchain khr_swapchain.Swapchain

	common.NextOptions
}

func (o ImageSwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSwapchainCreateInfoKHR{})))
	}

	info := (*C.VkImageSwapchainCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
	info.pNext = next
	info.swapchain = nil

	if o.Swapchain != nil {
		info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	}

	return preallocatedPointer, nil
}
