package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions struct {
	SeparateDepthStencilLayouts bool

	common.HaveNext
}

func (o PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
	info.pNext = next
	info.separateDepthStencilLayouts = C.VkBool32(0)

	if o.SeparateDepthStencilLayouts {
		info.separateDepthStencilLayouts = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
