package khr_sampler_mirror_clamp_to_edge

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
)

const (
	ExtensionName string = C.VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME

	SamplerAddressModeMirrorClampToEdge core1_0.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR
)

func init() {
	SamplerAddressModeMirrorClampToEdge.Register("Mirror Clamp To Edge")
}
