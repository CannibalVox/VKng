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

type DescriptorSetLayoutBindingFlagsCreateOptions struct {
	BindingFlags []DescriptorBindingFlags

	common.HaveNext
}

func (o DescriptorSetLayoutBindingFlagsCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT{})))
	}

	info := (*C.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
	info.pNext = next

	count := len(o.BindingFlags)
	info.bindingCount = C.uint32_t(count)
	info.pBindingFlags = nil

	if count > 0 {
		info.pBindingFlags = (*C.VkDescriptorBindingFlags)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkDescriptorBindingFlags(0)))))
		flagSlice := unsafe.Slice(info.pBindingFlags, count)

		for i := 0; i < count; i++ {
			flagSlice[i] = C.VkDescriptorBindingFlags(o.BindingFlags[i])
		}
	}

	return preallocatedPointer, nil
}

func (o DescriptorSetLayoutBindingFlagsCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT)(cDataPointer)
	return info.pNext, nil
}
