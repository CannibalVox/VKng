package khr_imageless_framebuffer

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

type PhysicalDeviceImagelessFramebufferFeaturesOutData struct {
	ImagelessFramebuffer bool

	common.HaveNext
}

func (o *PhysicalDeviceImagelessFramebufferFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceImagelessFramebufferFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(cDataPointer)

	o.ImagelessFramebuffer = info.imagelessFramebuffer != C.VkBool32(0)

	return info.pNext, nil
}
