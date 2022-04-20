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

type DeviceGroupBindSparseOptions struct {
	ResourceDeviceIndex int
	MemoryDeviceIndex   int

	common.HaveNext
}

func (o DeviceGroupBindSparseOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupBindSparseInfo{})))
	}

	createInfo := (*C.VkDeviceGroupBindSparseInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR
	createInfo.pNext = next
	createInfo.resourceDeviceIndex = C.uint32_t(o.ResourceDeviceIndex)
	createInfo.memoryDeviceIndex = C.uint32_t(o.MemoryDeviceIndex)

	return preallocatedPointer, nil
}

func (o DeviceGroupBindSparseOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkDeviceGroupBindSparseInfo)(cDataPointer)
	return createInfo.pNext, nil
}
