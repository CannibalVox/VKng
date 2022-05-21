package khr_buffer_device_address

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

type DeviceMemoryOpaqueAddressOptions struct {
	Memory core1_0.DeviceMemory

	common.HaveNext
}

func (o DeviceMemoryOpaqueAddressOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR{})))
	}

	info := (*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
	info.pNext = next
	info.memory = C.VkDeviceMemory(unsafe.Pointer(o.Memory.Handle()))

	return preallocatedPointer, nil
}

func (o DeviceMemoryOpaqueAddressOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(cDataPointer)
	return info.pNext, nil
}
