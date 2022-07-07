package khr_external_fence_capabilities

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

type ExternalFenceProperties struct {
	ExportFromImportedHandleTypes ExternalFenceHandleTypeFlags
	CompatibleHandleTypes         ExternalFenceHandleTypeFlags
	ExternalFenceFeatures         ExternalFenceFeatureFlags

	common.NextOutData
}

func (o *ExternalFenceProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalFencePropertiesKHR{})))
	}

	info := (*C.VkExternalFencePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalFenceProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalFencePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalFenceHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalFenceHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalFenceFeatures = ExternalFenceFeatureFlags(info.externalFenceFeatures)

	return info.pNext, nil
}
