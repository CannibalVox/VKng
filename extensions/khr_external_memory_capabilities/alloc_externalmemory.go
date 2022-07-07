package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ExternalMemoryProperties struct {
	ExternalMemoryFeatures        ExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         ExternalMemoryHandleTypeFlags
}

func (o ExternalMemoryProperties) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryPropertiesKHR{})))
	}

	info := (*C.VkExternalMemoryPropertiesKHR)(preallocatedPointer)
	info.externalMemoryFeatures = C.VkExternalMemoryFeatureFlags(o.ExternalMemoryFeatures)
	info.exportFromImportedHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.ExportFromImportedHandleTypes)
	info.compatibleHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.CompatibleHandleTypes)

	return preallocatedPointer, nil
}

func (o *ExternalMemoryProperties) PopulateOutData(cDataPointer unsafe.Pointer) error {
	info := (*C.VkExternalMemoryPropertiesKHR)(cDataPointer)
	o.ExternalMemoryFeatures = ExternalMemoryFeatureFlags(info.externalMemoryFeatures)
	o.ExportFromImportedHandleTypes = ExternalMemoryHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalMemoryHandleTypeFlags(info.compatibleHandleTypes)

	return nil
}
