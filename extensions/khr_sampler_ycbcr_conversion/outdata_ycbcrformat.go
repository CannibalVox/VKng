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

type SamplerYcbcrConversionImageFormatProperties struct {
	CombinedImageSamplerDescriptorCount int

	common.NextOutData
}

func (o *SamplerYcbcrConversionImageFormatProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionImageFormatPropertiesKHR{})))
	}

	info := (*C.VkSamplerYcbcrConversionImageFormatPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SamplerYcbcrConversionImageFormatProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSamplerYcbcrConversionImageFormatPropertiesKHR)(cDataPointer)

	o.CombinedImageSamplerDescriptorCount = int(info.combinedImageSamplerDescriptorCount)

	return info.pNext, nil
}
