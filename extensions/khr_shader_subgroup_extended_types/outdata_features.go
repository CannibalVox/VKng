package khr_shader_subgroup_extended_types

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

type PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOutData struct {
	ShaderSubgroupExtendedTypes bool

	common.HaveNext
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(cDataPointer)

	o.ShaderSubgroupExtendedTypes = info.shaderSubgroupExtendedTypes != C.VkBool32(0)

	return info.pNext, nil
}
