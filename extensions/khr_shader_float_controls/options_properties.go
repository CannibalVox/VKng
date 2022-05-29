package khr_shader_float_controls

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

type PhysicalDeviceFloatControlsOutData struct {
	DenormBehaviorIndependence ShaderFloatControlsIndependence
	RoundingMoundIndependence  ShaderFloatControlsIndependence

	ShaderSignedZeroInfNanPreserveFloat16 bool
	ShaderSignedZeroInfNanPreserveFloat32 bool
	ShaderSignedZeroInfNanPreserveFloat64 bool
	ShaderDenormPreserveFloat16           bool
	ShaderDenormPreserveFloat32           bool
	ShaderDenormPreserveFloat64           bool
	ShaderDenormFlushToZeroFloat16        bool
	ShaderDenormFlushToZeroFloat32        bool
	ShaderDenormFlushToZeroFloat64        bool
	ShaderRoundingModeRTEFloat16          bool
	ShaderRoundingModeRTEFloat32          bool
	ShaderRoundingModeRTEFloat64          bool
	ShaderRoundingModeRTZFloat16          bool
	ShaderRoundingModeRTZFloat32          bool
	ShaderRoundingModeRTZFloat64          bool

	common.HaveNext
}

func (o *PhysicalDeviceFloatControlsOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFloatControlsPropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceFloatControlsPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFloatControlsOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceFloatControlsPropertiesKHR)(cDataPointer)

	o.DenormBehaviorIndependence = ShaderFloatControlsIndependence(info.denormBehaviorIndependence)
	o.RoundingMoundIndependence = ShaderFloatControlsIndependence(info.roundingModeIndependence)
	o.ShaderSignedZeroInfNanPreserveFloat16 = info.shaderSignedZeroInfNanPreserveFloat16 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat32 = info.shaderSignedZeroInfNanPreserveFloat32 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat64 = info.shaderSignedZeroInfNanPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat16 = info.shaderDenormPreserveFloat16 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat32 = info.shaderDenormPreserveFloat32 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat64 = info.shaderDenormPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat16 = info.shaderDenormFlushToZeroFloat16 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat32 = info.shaderDenormFlushToZeroFloat32 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat64 = info.shaderDenormFlushToZeroFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat16 = info.shaderRoundingModeRTEFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat32 = info.shaderRoundingModeRTEFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat64 = info.shaderRoundingModeRTEFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat16 = info.shaderRoundingModeRTZFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat32 = info.shaderRoundingModeRTZFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat64 = info.shaderRoundingModeRTZFloat64 != C.VkBool32(0)

	return info.pNext, nil
}
