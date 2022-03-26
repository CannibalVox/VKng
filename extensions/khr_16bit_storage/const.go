package khr_16bit_storage

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	ExtensionName string = C.VK_KHR_16BIT_STORAGE_EXTENSION_NAME
)

type VkPhysicalDevice16BitStorageFeaturesKHR C.VkPhysicalDevice16BitStorageFeaturesKHR
