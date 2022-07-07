package khr_maintenance1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type CommandPoolTrimFlags int32

var commandPoolTrimFlagsMapping = common.NewFlagStringMapping[CommandPoolTrimFlags]()

func (f CommandPoolTrimFlags) Register(str string) {
	commandPoolTrimFlagsMapping.Register(f, str)
}
func (f CommandPoolTrimFlags) String() string {
	return commandPoolTrimFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_MAINTENANCE1_EXTENSION_NAME

	FormatFeatureTransferDst core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR
	FormatFeatureTransferSrc core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR

	ImageCreate2DArrayCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR

	VkErrorOutOfPoolMemory common.VkResult = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
)

func init() {
	FormatFeatureTransferDst.Register("Transfer Destination")
	FormatFeatureTransferSrc.Register("Transfer Source")

	ImageCreate2DArrayCompatible.Register("2D Array Compatible")

	VkErrorOutOfPoolMemory.Register("out of pool memory")
}
