package khr_device_group

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

type MemoryAllocateFlagsOptions struct {
	Flags      MemoryAllocateFlags
	DeviceMask uint32

	common.NextOptions
}

func (o MemoryAllocateFlagsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryAllocateFlagsInfoKHR{})))
	}

	createInfo := (*C.VkMemoryAllocateFlagsInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR
	createInfo.pNext = next
	createInfo.flags = C.VkMemoryAllocateFlags(o.Flags)
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
