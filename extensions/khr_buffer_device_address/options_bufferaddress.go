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
	"github.com/cockroachdb/errors"
	"unsafe"
)

type BufferDeviceAddressInfo struct {
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o BufferDeviceAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferDeviceAddressInfoKHR{})))
	}

	if o.Buffer == nil {
		return nil, errors.New("khr_buffer_device_address.DeviceMemoryAddressOptions.Buffer cannot be nil")
	}

	info := (*C.VkBufferDeviceAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
	info.pNext = next
	info.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}
