package khr_get_physical_device_properties2

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

type FormatPropertiesOutData struct {
	FormatProperties core1_0.FormatProperties
	common.HaveNext
}

func (o *FormatPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFormatProperties2KHR{})))
	}

	data := (*C.VkFormatProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *FormatPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkFormatProperties2KHR)(cDataPointer)
	o.FormatProperties.LinearTilingFeatures = common.FormatFeatures(data.formatProperties.linearTilingFeatures)
	o.FormatProperties.OptimalTilingFeatures = common.FormatFeatures(data.formatProperties.optimalTilingFeatures)
	o.FormatProperties.BufferFeatures = common.FormatFeatures(data.formatProperties.bufferFeatures)

	return data.pNext, nil
}