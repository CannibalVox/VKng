package khr_portability_subset

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_beta.h"
*/
import "C"

const (
	ExtensionName string = C.VK_KHR_PORTABILITY_SUBSET_EXTENSION_NAME
)
