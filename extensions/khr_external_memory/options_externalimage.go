package khr_external_memory

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ExternalMemoryImageOptions struct {
	HandleTypes khr_external_memory_capabilities.ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExternalMemoryImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryImageCreateInfoKHR{})))
	}

	info := (*C.VkExternalMemoryImageCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlagsKHR(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExternalMemoryImageOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalMemoryImageCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
