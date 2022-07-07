package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ImageViewUsageCreateInfo struct {
	Usage core1_0.ImageUsageFlags

	common.NextOptions
}

func (o ImageViewUsageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageViewUsageCreateInfoKHR{})))
	}

	createInfo := (*C.VkImageViewUsageCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.usage = C.VkImageUsageFlags(o.Usage)

	return preallocatedPointer, nil
}
