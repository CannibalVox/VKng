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

type BindBufferMemoryDeviceGroupOptions struct {
	DeviceIndices []int

	common.HaveNext
}

func (o BindBufferMemoryDeviceGroupOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindBufferMemoryDeviceGroupInfoKHR{})))
	}

	info := (*C.VkBindBufferMemoryDeviceGroupInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR
	info.pNext = next

	count := len(o.DeviceIndices)
	info.deviceIndexCount = C.uint32_t(count)
	info.pDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.DeviceIndices[i])
		}

		info.pDeviceIndices = indices
	}

	return preallocatedPointer, nil
}

func (o BindBufferMemoryDeviceGroupOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkBindBufferMemoryDeviceGroupInfoKHR)(cDataPointer)
	return info.pNext, nil
}
