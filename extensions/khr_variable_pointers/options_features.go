package khr_variable_pointers

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

type VariablePointersFeatureOptions struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.HaveNext
}

func (o VariablePointersFeatureOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeaturesKHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeaturesKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
	createInfo.pNext = next
	createInfo.variablePointersStorageBuffer = C.VkBool32(0)
	createInfo.variablePointers = C.VkBool32(0)

	if o.VariablePointersStorageBuffer {
		createInfo.variablePointersStorageBuffer = C.VkBool32(1)
	}
	if o.VariablePointers {
		createInfo.variablePointers = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o VariablePointersFeatureOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceVariablePointersFeaturesKHR)(cDataPointer)
	return createInfo.pNext, nil
}
