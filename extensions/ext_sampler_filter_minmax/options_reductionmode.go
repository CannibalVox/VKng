package ext_sampler_filter_minmax

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SamplerReductionModeCreateInfo struct {
	ReductionMode SamplerReductionMode

	common.NextOptions
}

func (o SamplerReductionModeCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerReductionModeCreateInfoEXT{})))
	}

	info := (*C.VkSamplerReductionModeCreateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT
	info.pNext = next
	info.reductionMode = C.VkSamplerReductionModeEXT(o.ReductionMode)

	return preallocatedPointer, nil
}
