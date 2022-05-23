package khr_sampler_ycbcr_conversion

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

type PhysicalDeviceSamplerYcbcrFeaturesOptions struct {
	SamplerYcbcrConversion bool

	common.HaveNext
}

func (o PhysicalDeviceSamplerYcbcrFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
	info.pNext = next
	info.samplerYcbcrConversion = C.VkBool32(0)

	if o.SamplerYcbcrConversion {
		info.samplerYcbcrConversion = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceSamplerYcbcrFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}