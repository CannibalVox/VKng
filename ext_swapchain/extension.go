package ext_swapchain

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

