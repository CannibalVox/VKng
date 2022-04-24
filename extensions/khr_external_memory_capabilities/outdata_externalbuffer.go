package khr_external_memory_capabilities

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

type ExternalBufferOutData struct {
	ExternalMemoryProperties ExternalMemoryProperties

	common.HaveNext
}

func (o *ExternalBufferOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalBufferPropertiesKHR{})))
	}
	info := (*C.VkExternalBufferPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalBufferOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalBufferPropertiesKHR)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, nil
}
