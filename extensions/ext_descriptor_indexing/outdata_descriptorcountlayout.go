package ext_descriptor_indexing

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

type DescriptorSetVariableDescriptorCountLayoutSupportOutData struct {
	MaxVariableDescriptorCount int

	common.HaveNext
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupportOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT{})))
	}

	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupportOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(cDataPointer)

	o.MaxVariableDescriptorCount = int(info.maxVariableDescriptorCount)

	return info.pNext, nil
}
