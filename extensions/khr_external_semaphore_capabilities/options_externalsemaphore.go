package khr_external_semaphore_capabilities

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

type PhysicalDeviceExternalSemaphoreInfo struct {
	HandleType ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalSemaphoreInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalSemaphoreInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalSemaphoreInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}
