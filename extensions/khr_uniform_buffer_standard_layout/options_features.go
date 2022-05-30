package khr_uniform_buffer_standard_layout

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

type PhysicalDeviceUniformBufferStandardLayoutFeaturesOptions struct {
	UniformBufferStandardLayout bool

	common.HaveNext
}

func (o PhysicalDeviceUniformBufferStandardLayoutFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR
	info.pNext = next
	info.uniformBufferStandardLayout = C.VkBool32(0)

	if o.UniformBufferStandardLayout {
		info.uniformBufferStandardLayout = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceUniformBufferStandardLayoutFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
