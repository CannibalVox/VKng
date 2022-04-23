package khr_multiview

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ExtensionName string = C.VK_KHR_MULTIVIEW_EXTENSION_NAME

	DependencyViewLocal common.DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR
)

func init() {
	DependencyViewLocal.Register("View Local")
}
