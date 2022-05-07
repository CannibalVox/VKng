package khr_get_memory_requirements2

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

type BufferMemoryRequirementsOptions struct {
	Buffer core1_0.Buffer

	common.HaveNext
}

func (o BufferMemoryRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferMemoryRequirementsInfo2KHR{})))
	}

	options := (*C.VkBufferMemoryRequirementsInfo2KHR)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
	options.pNext = next
	options.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}

func (o BufferMemoryRequirementsOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	options := (*C.VkBufferMemoryRequirementsInfo2KHR)(cDataPointer)
	return options.pNext, nil
}
