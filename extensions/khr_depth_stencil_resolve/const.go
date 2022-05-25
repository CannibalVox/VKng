package khr_depth_stencil_resolve

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ResolveModeFlags int32

var resolveModeFlagsMapping = common.NewFlagStringMapping[ResolveModeFlags]()

func (f ResolveModeFlags) Register(str string) {
	resolveModeFlagsMapping.Register(f, str)
}

func (f ResolveModeFlags) String() string {
	return resolveModeFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME

	ResolveModeAverage    ResolveModeFlags = C.VK_RESOLVE_MODE_AVERAGE_BIT_KHR
	ResolveModeMax        ResolveModeFlags = C.VK_RESOLVE_MODE_MAX_BIT_KHR
	ResolveModeMin        ResolveModeFlags = C.VK_RESOLVE_MODE_MIN_BIT_KHR
	ResolveModeNone       ResolveModeFlags = C.VK_RESOLVE_MODE_NONE_KHR
	ResolveModeSampleZero ResolveModeFlags = C.VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR
)

func init() {
	ResolveModeAverage.Register("Average")
	ResolveModeMax.Register("Max")
	ResolveModeMin.Register("Min")
	ResolveModeNone.Register("None")
	ResolveModeSampleZero.Register("Sample Zero")
}
