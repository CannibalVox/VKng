package khr_device_group_creation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/core1_0"

const (
	ExtensionName string = C.VK_KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME

	MaxDeviceGroupSize int = C.VK_MAX_DEVICE_GROUP_SIZE_KHR

	MemoryHeapMultiInstance core1_0.MemoryHeapFlags = C.VK_MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR
)

func init() {
	MemoryHeapMultiInstance.Register("Multi-Instance")
}
