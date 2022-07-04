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

type BindImageMemorySwapchainOptions struct {
	Swapchain  khr_swapchain.Swapchain
	ImageIndex int

	common.NextOptions
}

func (o BindImageMemorySwapchainOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemorySwapchainInfoKHR{})))
	}

	info := (*C.VkBindImageMemorySwapchainInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
	info.pNext = next
	info.swapchain = nil
	info.imageIndex = C.uint32_t(o.ImageIndex)

	if o.Swapchain != nil {
		info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	}

	return preallocatedPointer, nil
}
