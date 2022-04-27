package khr_external_memory

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_MEMORY_EXTENSION_NAME

	QueueFamilyExternal int = C.VK_QUEUE_FAMILY_EXTERNAL_KHR

	VkErrorInvalidExternalHandle common.VkResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
)

func init() {
	VkErrorInvalidExternalHandle.Register("invalid external handle")
}
