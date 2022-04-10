package khr_dedicated_allocation

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

type DedicatedAllocationOptions struct {
	Image  core1_0.Image
	Buffer core1_0.Buffer

	common.HaveNext
}

func (o DedicatedAllocationOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image != nil && o.Buffer != nil {
		return nil, errors.New("both Image and Buffer fields are set in DedicatedAllocationOptions- only one must be set")
	} else if o.Image == nil && o.Buffer == nil {
		return nil, errors.New("neither Image nor Buffer fields are set in DedicatedAllocationOptions- one must be set")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedAllocateInfo{})))
	}

	createInfo := (*C.VkMemoryDedicatedAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR
	createInfo.pNext = next
	createInfo.image = nil
	createInfo.buffer = nil

	if o.Image != nil {
		createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	} else if o.Buffer != nil {
		createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	}

	return preallocatedPointer, nil
}

func (o DedicatedAllocationOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkMemoryDedicatedAllocateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
