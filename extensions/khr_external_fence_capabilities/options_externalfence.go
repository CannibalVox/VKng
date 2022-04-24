package khr_external_fence_capabilities

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

type ExternalFencePropertiesOptions struct {
	HandleType ExternalFenceHandleTypes

	common.HaveNext
}

func (o ExternalFencePropertiesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalFenceInfoKHR{})))
	}
	info := (*C.VkPhysicalDeviceExternalFenceInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalFenceHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

func (o ExternalFencePropertiesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceExternalFenceInfoKHR)(cDataPointer)
	return info.pNext, nil
}
