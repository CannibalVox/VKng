package khr_descriptor_update_template

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

type DescriptorTemplateEntry struct {
	DstBinding      int
	DstArrayElement int
	DescriptorCount int

	DescriptorType common.DescriptorType

	Offset int
	Stride int
}

func (e DescriptorTemplateEntry) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateEntryKHR{})))
	}

	entry := (*C.VkDescriptorUpdateTemplateEntryKHR)(preallocatedPointer)
	entry.dstBinding = C.uint32_t(e.DstBinding)
	entry.dstArrayElement = C.uint32_t(e.DstArrayElement)
	entry.descriptorCount = C.uint32_t(e.DescriptorCount)
	entry.descriptorType = C.VkDescriptorType(e.DescriptorType)
	entry.offset = C.size_t(e.Offset)
	entry.stride = C.size_t(e.Stride)

	return preallocatedPointer, nil
}
