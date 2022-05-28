package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ExtensionName string = C.VK_KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME

	FramebufferCreateImageless common.FramebufferCreateFlags = C.VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR
)
