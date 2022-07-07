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

type ExportSemaphoreCreateInfo struct {
	HandleTypes khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o ExportSemaphoreCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportSemaphoreCreateInfoKHR{})))
	}

	info := (*C.VkExportSemaphoreCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalSemaphoreHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
