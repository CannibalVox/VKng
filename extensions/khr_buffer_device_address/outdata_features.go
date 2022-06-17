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

type PhysicalDeviceBufferAddressFeaturesOutData struct {
	BufferDeviceAddress              bool
	BufferDeviceAddressCaptureReplay bool
	BufferDeviceAddressMultiDevice   bool

	common.HaveNext
}

func (o *PhysicalDeviceBufferAddressFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceBufferAddressFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(cDataPointer)

	o.BufferDeviceAddress = info.bufferDeviceAddress != C.VkBool32(0)
	o.BufferDeviceAddressCaptureReplay = info.bufferDeviceAddressCaptureReplay != C.VkBool32(0)
	o.BufferDeviceAddressMultiDevice = info.bufferDeviceAddressMultiDevice != C.VkBool32(0)

	return info.pNext, nil
}
