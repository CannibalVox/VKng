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

type DeviceMemoryOpaqueAddressInfo struct {
	Memory core1_0.DeviceMemory

	common.NextOptions
}

func (o DeviceMemoryOpaqueAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR{})))
	}

	info := (*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
	info.pNext = next
	info.memory = C.VkDeviceMemory(unsafe.Pointer(o.Memory.Handle()))

	return preallocatedPointer, nil
}
