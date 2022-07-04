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

type ExportMemoryAllocateOptions struct {
	HandleTypes khr_external_memory_capabilities.ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExportMemoryAllocateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportMemoryAllocateInfoKHR{})))
	}

	info := (*C.VkExportMemoryAllocateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlagsKHR(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExportMemoryAllocateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExportMemoryAllocateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
