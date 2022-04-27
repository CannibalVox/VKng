package khr_external_semaphore_capabilities

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

type ExternalSemaphoreOutData struct {
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypes
	CompatibleHandleTypes         ExternalSemaphoreHandleTypes
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatures

	common.HaveNext
}

func (o *ExternalSemaphoreOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalSemaphorePropertiesKHR{})))
	}

	info := (*C.VkExternalSemaphorePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalSemaphoreOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalSemaphorePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypes(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalSemaphoreHandleTypes(info.compatibleHandleTypes)
	o.ExternalSemaphoreFeatures = ExternalSemaphoreFeatures(info.externalSemaphoreFeatures)

	return info.pNext, nil
}
