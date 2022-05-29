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

type PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOptions struct {
	ShaderSubgroupExtendedTypes bool

	common.HaveNext
}

func (o PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR
	info.pNext = next
	info.shaderSubgroupExtendedTypes = C.VkBool32(0)

	if o.ShaderSubgroupExtendedTypes {
		info.shaderSubgroupExtendedTypes = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceShaderSubgroupExtendedTypesFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
