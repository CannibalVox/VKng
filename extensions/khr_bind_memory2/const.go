package khr_bind_memory2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/core1_0"

const (
	ExtensionName string = C.VK_KHR_BIND_MEMORY_2_EXTENSION_NAME

	ImageCreateAlias core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_ALIAS_BIT_KHR
)

func init() {
	ImageCreateAlias.Register("Alias")
}
