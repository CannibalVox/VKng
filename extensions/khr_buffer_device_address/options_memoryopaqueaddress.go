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

type MemoryOpaqueCaptureAddressAllocateOptions struct {
	OpaqueCaptureAddress uint64

	common.HaveNext
}

func (o MemoryOpaqueCaptureAddressAllocateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR{})))
	}

	info := (*C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO
	info.pNext = next
	info.opaqueCaptureAddress = C.uint64_t(o.OpaqueCaptureAddress)

	return preallocatedPointer, nil
}

func (o MemoryOpaqueCaptureAddressAllocateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR)(cDataPointer)
	return info.pNext, nil
}