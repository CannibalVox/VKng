package ext_separate_stencil_usage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	ExtensionName string = C.VK_EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME
)
