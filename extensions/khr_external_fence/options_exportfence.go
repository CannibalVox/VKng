package khr_external_fence

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ExportFenceCreateInfo struct {
	HandleTypes khr_external_fence_capabilities.ExternalFenceHandleTypeFlags

	common.NextOptions
}

func (o ExportFenceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportFenceCreateInfoKHR{})))
	}

	info := (*C.VkExportFenceCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalFenceHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
