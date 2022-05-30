package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ExtensionName string = C.VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME

	ImageLayoutDepthAttachmentOptimal   common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR
	ImageLayoutDepthReadOnlyOptimal     common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR
	ImageLayoutStencilAttachmentOptimal common.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR
	ImageLayoutStencilReadOnlyOptimal   common.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR
)

func init() {
	ImageLayoutDepthAttachmentOptimal.Register("Depth Attachment Optimal")
	ImageLayoutDepthReadOnlyOptimal.Register("Depth Read-Only Optimal")
	ImageLayoutStencilAttachmentOptimal.Register("Stencil Attachment Optimal")
	ImageLayoutStencilReadOnlyOptimal.Register("Stencil Read-Only Optimal")
}
