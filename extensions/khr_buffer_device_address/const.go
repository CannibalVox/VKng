package khr_buffer_device_address

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
)

const (
	ExtensionName string = C.VK_KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME

	BufferCreateDeviceAddressCaptureReplay core1_0.BufferCreateFlags = C.VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR

	BufferUsageShaderDeviceAddress core1_0.BufferUsages = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR

	MemoryAllocateDeviceAddress              core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR
	MemoryAllocateDeviceAddressCaptureReplay core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR

	VkErrorInvalidOpaqueCaptureAddress common.VkResult = C.VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
)

func init() {
	BufferCreateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	BufferUsageShaderDeviceAddress.Register("Shader Device Address")

	MemoryAllocateDeviceAddress.Register("Device Address")
	MemoryAllocateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	VkErrorInvalidOpaqueCaptureAddress.Register("invalid opaque capture address")
}
