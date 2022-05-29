package khr_sampler_mirror_clamp_to_edge

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ExtensionName string = C.VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME

	SamplerAddressModeMirrorClampToEdge common.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR
)
