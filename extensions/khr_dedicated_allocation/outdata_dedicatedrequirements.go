package khr_dedicated_allocation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type MemoryDedicatedRequirements struct {
	PrefersDedicatedAllocation  bool
	RequiresDedicatedAllocation bool

	common.NextOutData
}

func (o *MemoryDedicatedRequirements) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedRequirementsKHR{})))
	}

	outData := (*C.VkMemoryDedicatedRequirementsKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryDedicatedRequirements) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryDedicatedRequirementsKHR)(cDataPointer)
	o.RequiresDedicatedAllocation = driver.VkBool32(outData.requiresDedicatedAllocation) != driver.VkBool32(0)
	o.PrefersDedicatedAllocation = driver.VkBool32(outData.prefersDedicatedAllocation) != driver.VkBool32(0)

	return outData.pNext, nil
}
