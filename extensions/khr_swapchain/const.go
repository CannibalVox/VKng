package khr_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type SwapchainCreateFlags int32

var swapchainCreateFlagsMapping = common.NewFlagStringMapping[SwapchainCreateFlags]()

func (f SwapchainCreateFlags) Register(str string) {
	swapchainCreateFlagsMapping.Register(f, str)
}
func (f SwapchainCreateFlags) String() string {
	return swapchainCreateFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

	ObjectTypeSwapchain core1_0.ObjectType = C.VK_OBJECT_TYPE_SWAPCHAIN_KHR

	ImageLayoutPresentSrc core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR

	VKErrorOutOfDate common.VkResult = C.VK_ERROR_OUT_OF_DATE_KHR
	VKSuboptimal     common.VkResult = C.VK_SUBOPTIMAL_KHR
)

func init() {
	ObjectTypeSwapchain.Register("Swapchain")

	ImageLayoutPresentSrc.Register("Present Src")

	VKErrorOutOfDate.Register("out of date")
	VKSuboptimal.Register("Suboptimal")
}
