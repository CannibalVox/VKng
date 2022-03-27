package khr_maintenance3

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

type DescriptorSetLayoutSupportOutData struct {
	Supported bool

	common.HaveNext
}

func (o *DescriptorSetLayoutSupportOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutSupportKHR{})))
	}

	outData := (*C.VkDescriptorSetLayoutSupportKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetLayoutSupportOutData) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	outData := (*C.VkDescriptorSetLayoutSupportKHR)(cPointer)
	o.Supported = outData.supported != C.VkBool32(0)

	return outData.pNext, nil
}
