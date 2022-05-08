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

type PhysicalDeviceExternalImageFormatOptions struct {
	HandleType ExternalMemoryHandleTypes

	common.HaveNext
}

func (o PhysicalDeviceExternalImageFormatOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalImageFormatInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalImageFormatInfoKHR)(preallocatedPointer)

	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalMemoryHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

func (o PhysicalDeviceExternalImageFormatOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceExternalImageFormatInfoKHR)(cDataPointer)
	return info.pNext, nil
}
