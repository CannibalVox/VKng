package khr_shader_atomic_int64

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

type PhysicalDeviceShaderAtomicInt64FeaturesOptions struct {
	ShaderBufferInt64Atomics bool
	ShaderSharedInt64Atomics bool

	common.HaveNext
}

func (o PhysicalDeviceShaderAtomicInt64FeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR
	info.pNext = next
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)

	if o.ShaderBufferInt64Atomics {
		info.shaderBufferInt64Atomics = C.VkBool32(1)
	}

	if o.ShaderSharedInt64Atomics {
		info.shaderSharedInt64Atomics = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceShaderAtomicInt64FeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
