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

type ExternalSemaphoreProperties struct {
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         ExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatureFlags

	common.NextOutData
}

func (o *ExternalSemaphoreProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalSemaphorePropertiesKHR{})))
	}

	info := (*C.VkExternalSemaphorePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalSemaphoreProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalSemaphorePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalSemaphoreHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalSemaphoreFeatures = ExternalSemaphoreFeatureFlags(info.externalSemaphoreFeatures)

	return info.pNext, nil
}
