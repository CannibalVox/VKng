package khr_external_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ExportSemaphoreOptions struct {
	HandleTypes khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypes

	common.HaveNext
}

func (o ExportSemaphoreOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportSemaphoreCreateInfoKHR{})))
	}

	info := (*C.VkExportSemaphoreCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalSemaphoreHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExportSemaphoreOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExportSemaphoreCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
