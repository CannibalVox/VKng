package khr_external_memory_capabilities

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

type PhysicalDeviceExternalBufferInfo struct {
	Flags      core1_0.BufferCreateFlags
	Usage      core1_0.BufferUsageFlags
	HandleType ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalBufferInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalBufferInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalBufferInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR
	info.pNext = next
	info.flags = (C.VkBufferCreateFlags)(o.Flags)
	info.usage = (C.VkBufferUsageFlags)(o.Usage)
	info.handleType = (C.VkExternalMemoryHandleTypeFlagBits)(o.HandleType)

	return preallocatedPointer, nil
}
