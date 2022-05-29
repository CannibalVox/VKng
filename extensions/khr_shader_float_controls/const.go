package khr_shader_float_controls

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ShaderFloatControlsIndependence int32

var shaderFloatControlsIndependenceMapping = make(map[ShaderFloatControlsIndependence]string)

func (e ShaderFloatControlsIndependence) Register(str string) {
	shaderFloatControlsIndependenceMapping[e] = str
}

func (e ShaderFloatControlsIndependence) String() string {
	return shaderFloatControlsIndependenceMapping[e]
}

////

const (
	ExtensionName string = C.VK_KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME

	ShaderFloatControlsIndependence32BitOnly ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR
	ShaderFloatControlsIndependenceAll       ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR
	ShaderFloatControlsIndependenceNone      ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR
)

func init() {
	ShaderFloatControlsIndependenceAll.Register("All")
	ShaderFloatControlsIndependenceNone.Register("None")
	ShaderFloatControlsIndependence32BitOnly.Register("32-Bit Only")
}
