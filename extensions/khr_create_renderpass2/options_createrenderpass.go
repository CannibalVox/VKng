package khr_create_renderpass2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type RenderPassCreateInfo2 struct {
	Flags core1_0.RenderPassCreateFlags

	Attachments  []AttachmentDescription2
	Subpasses    []SubpassDescription2
	Dependencies []SubpassDependency2

	CorrelatedViewMasks []uint32

	common.NextOptions
}

func (o RenderPassCreateInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo2KHR{})))
	}

	info := (*C.VkRenderPassCreateInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
	info.pNext = next
	info.flags = C.VkRenderPassCreateFlags(o.Flags)

	attachmentCount := len(o.Attachments)
	subpassCount := len(o.Subpasses)
	dependencyCount := len(o.Dependencies)
	viewMaskCount := len(o.CorrelatedViewMasks)

	info.attachmentCount = C.uint32_t(attachmentCount)
	info.pAttachments = nil
	info.subpassCount = C.uint32_t(subpassCount)
	info.pSubpasses = nil
	info.dependencyCount = C.uint32_t(dependencyCount)
	info.pDependencies = nil
	info.correlatedViewMaskCount = C.uint32_t(viewMaskCount)
	info.pCorrelatedViewMasks = nil

	var err error
	if attachmentCount > 0 {
		info.pAttachments, err = common.AllocOptionSlice[C.VkAttachmentDescription2KHR, AttachmentDescription2](allocator, o.Attachments)
		if err != nil {
			return nil, err
		}
	}

	if subpassCount > 0 {
		info.pSubpasses, err = common.AllocOptionSlice[C.VkSubpassDescription2KHR, SubpassDescription2](allocator, o.Subpasses)
		if err != nil {
			return nil, err
		}
	}

	if dependencyCount > 0 {
		info.pDependencies, err = common.AllocOptionSlice[C.VkSubpassDependency2KHR, SubpassDependency2](allocator, o.Dependencies)
		if err != nil {
			return nil, err
		}
	}

	if viewMaskCount > 0 {
		viewMaskPtr := (*C.uint32_t)(allocator.Malloc(viewMaskCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := unsafe.Slice(viewMaskPtr, viewMaskCount)
		for i := 0; i < viewMaskCount; i++ {
			viewMaskSlice[i] = C.uint32_t(o.CorrelatedViewMasks[i])
		}
		info.pCorrelatedViewMasks = viewMaskPtr
	}

	return preallocatedPointer, nil
}
