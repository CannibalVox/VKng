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

type ExternalFenceOutData struct {
	ExportFromImportedHandleTypes ExternalFenceHandleTypes
	CompatibleHandleTypes         ExternalFenceHandleTypes
	ExternalFenceFeatures         ExternalFenceFeatures

	common.NextOutData
}

func (o *ExternalFenceOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalFencePropertiesKHR{})))
	}

	info := (*C.VkExternalFencePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalFenceOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalFencePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalFenceHandleTypes(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalFenceHandleTypes(info.compatibleHandleTypes)
	o.ExternalFenceFeatures = ExternalFenceFeatures(info.externalFenceFeatures)

	return info.pNext, nil
}
