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

type SamplerYcbcrConversionCreateOptions struct {
	Format                      core1_0.DataFormat
	YcbcrModel                  SamplerYcbcrModelConversion
	YcbcrRange                  SamplerYcbcrRange
	Components                  core1_0.ComponentMapping
	ChromaOffsetX               ChromaLocation
	ChromaOffsetY               ChromaLocation
	ChromaFilter                core1_0.Filter
	ForceExplicitReconstruction bool

	common.NextOptions
}

func (o SamplerYcbcrConversionCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionCreateInfoKHR{})))
	}

	info := (*C.VkSamplerYcbcrConversionCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info.ycbcrModel = C.VkSamplerYcbcrModelConversion(o.YcbcrModel)
	info.ycbcrRange = C.VkSamplerYcbcrRange(o.YcbcrRange)
	info.components.r = C.VkComponentSwizzle(o.Components.R)
	info.components.g = C.VkComponentSwizzle(o.Components.G)
	info.components.b = C.VkComponentSwizzle(o.Components.B)
	info.components.a = C.VkComponentSwizzle(o.Components.A)
	info.xChromaOffset = C.VkChromaLocation(o.ChromaOffsetX)
	info.yChromaOffset = C.VkChromaLocation(o.ChromaOffsetY)
	info.chromaFilter = C.VkFilter(o.ChromaFilter)
	info.forceExplicitReconstruction = C.VkBool32(0)

	if o.ForceExplicitReconstruction {
		info.forceExplicitReconstruction = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
