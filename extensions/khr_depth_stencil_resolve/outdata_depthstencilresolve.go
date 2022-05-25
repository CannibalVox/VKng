package khr_depth_stencil_resolve

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

type PhysicalDeviceDepthStencilResolveOutData struct {
	SupportedDepthResolveModes   ResolveModeFlags
	SupportedStencilResolveModes ResolveModeFlags
	IndependentResolveNone       bool
	IndependentResolve           bool

	common.HaveNext
}

func (o *PhysicalDeviceDepthStencilResolveOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDepthStencilResolveOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR)(cDataPointer)
	o.SupportedStencilResolveModes = ResolveModeFlags(info.supportedStencilResolveModes)
	o.SupportedDepthResolveModes = ResolveModeFlags(info.supportedDepthResolveModes)
	o.IndependentResolveNone = info.independentResolveNone != C.VkBool32(0)
	o.IndependentResolve = info.independentResolve != C.VkBool32(0)

	return info.pNext, nil
}
