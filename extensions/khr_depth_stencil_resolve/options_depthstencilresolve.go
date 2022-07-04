package khr_depth_stencil_resolve

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_create_renderpass2"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SubpassDescriptionDepthStencilResolveOptions struct {
	DepthResolveMode              ResolveModeFlags
	StencilResolveMode            ResolveModeFlags
	DepthStencilResolveAttachment *khr_create_renderpass2.AttachmentReferenceOptions

	common.NextOptions
}

func (o SubpassDescriptionDepthStencilResolveOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescriptionDepthStencilResolveKHR{})))
	}

	info := (*C.VkSubpassDescriptionDepthStencilResolveKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR
	info.pNext = next
	info.depthResolveMode = C.VkResolveModeFlagBits(o.DepthResolveMode)
	info.stencilResolveMode = C.VkResolveModeFlagBits(o.StencilResolveMode)
	info.pDepthStencilResolveAttachment = nil

	if o.DepthStencilResolveAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.DepthStencilResolveAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilResolveAttachment = (*C.VkAttachmentReference2KHR)(attachment)
	}

	return preallocatedPointer, nil
}
