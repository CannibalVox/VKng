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

type MemoryRequirementsOutData struct {
	MemoryRequirements core1_0.MemoryRequirements
	common.NextOutData
}

func (o *MemoryRequirementsOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryRequirements2KHR{})))
	}

	outData := (*C.VkMemoryRequirements2KHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryRequirementsOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryRequirements2KHR)(cDataPointer)
	o.MemoryRequirements.Size = int(outData.memoryRequirements.size)
	o.MemoryRequirements.Alignment = int(outData.memoryRequirements.alignment)
	o.MemoryRequirements.MemoryType = uint32(outData.memoryRequirements.memoryTypeBits)

	return outData.pNext, nil
}
