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

type SparseImageFormatProperties2 struct {
	Properties core1_0.SparseImageFormatProperties
	common.NextOutData
}

func (o *SparseImageFormatProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageFormatProperties2KHR{})))
	}

	data := (*C.VkSparseImageFormatProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageFormatProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkSparseImageFormatProperties2KHR)(cDataPointer)

	o.Properties.AspectMask = core1_0.ImageAspectFlags(data.properties.aspectMask)
	o.Properties.Flags = core1_0.SparseImageFormatFlags(data.properties.flags)
	o.Properties.ImageGranularity = core1_0.Extent3D{
		Width:  int(data.properties.imageGranularity.width),
		Height: int(data.properties.imageGranularity.height),
		Depth:  int(data.properties.imageGranularity.depth),
	}

	return data.pNext, nil
}
