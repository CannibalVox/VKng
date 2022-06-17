package khr_buffer_device_address

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

type PhysicalDeviceBufferAddressFeaturesOptions struct {
	BufferDeviceAddress              bool
	BufferDeviceAddressCaptureReplay bool
	BufferDeviceAddressMultiDevice   bool

	common.HaveNext
}

func (o PhysicalDeviceBufferAddressFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR
	info.pNext = next
	info.bufferDeviceAddress = C.VkBool32(0)
	info.bufferDeviceAddressCaptureReplay = C.VkBool32(0)
	info.bufferDeviceAddressMultiDevice = C.VkBool32(0)

	if o.BufferDeviceAddress {
		info.bufferDeviceAddress = C.VkBool32(1)
	}

	if o.BufferDeviceAddressCaptureReplay {
		info.bufferDeviceAddressCaptureReplay = C.VkBool32(1)
	}

	if o.BufferDeviceAddressMultiDevice {
		info.bufferDeviceAddressMultiDevice = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceBufferAddressFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
