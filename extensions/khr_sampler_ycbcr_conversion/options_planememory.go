package khr_sampler_ycbcr_conversion

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ImagePlaneMemoryRequirementsInfo struct {
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o ImagePlaneMemoryRequirementsInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImagePlaneMemoryRequirementsInfoKHR{})))
	}

	info := (*C.VkImagePlaneMemoryRequirementsInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR
	info.pNext = next
	info.planeAspect = C.VkImageAspectFlagBits(o.PlaneAspect)

	return preallocatedPointer, nil
}
