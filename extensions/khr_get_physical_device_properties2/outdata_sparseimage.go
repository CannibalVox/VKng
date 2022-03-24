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

type SparseImageFormatPropertiesOutData struct {
	SparseImageFormatProperties core1_0.SparseImageFormatProperties
	common.HaveNext
}

func (o *SparseImageFormatPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageFormatProperties2KHR{})))
	}

	data := (*C.VkSparseImageFormatProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageFormatPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	data := (*C.VkSparseImageFormatProperties2KHR)(cDataPointer)

	o.SparseImageFormatProperties.AspectMask = common.ImageAspectFlags(data.properties.aspectMask)
	o.SparseImageFormatProperties.Flags = common.SparseImageFormatFlags(data.properties.flags)
	o.SparseImageFormatProperties.ImageGranularity = common.Extent3D{
		Width:  int(data.properties.imageGranularity.width),
		Height: int(data.properties.imageGranularity.height),
		Depth:  int(data.properties.imageGranularity.depth),
	}

	return data.pNext, nil
}
