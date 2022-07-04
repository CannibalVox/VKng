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

type PhysicalDeviceSamplerFilterMinmaxOutData struct {
	FilterMinmaxSingleComponentFormats bool
	FilterMinmaxImageComponentMapping  bool

	common.NextOutData
}

func (o *PhysicalDeviceSamplerFilterMinmaxOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT{})))
	}

	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerFilterMinmaxOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(cDataPointer)

	o.FilterMinmaxSingleComponentFormats = info.filterMinmaxSingleComponentFormats != C.VkBool32(0)
	o.FilterMinmaxImageComponentMapping = info.filterMinmaxImageComponentMapping != C.VkBool32(0)

	return info.pNext, nil
}
