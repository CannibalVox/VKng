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

type PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData struct {
	SeparateDepthStencilLayouts bool

	common.HaveNext
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(cDataPointer)

	o.SeparateDepthStencilLayouts = info.separateDepthStencilLayouts != C.VkBool32(0)

	return info.pNext, nil
}
