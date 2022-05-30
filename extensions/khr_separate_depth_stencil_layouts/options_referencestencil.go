package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type AttachmentReferenceStencilLayoutOptions struct {
	StencilLayout common.ImageLayout

	common.HaveNext
}

func (o AttachmentReferenceStencilLayoutOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReferenceStencilLayoutKHR{})))
	}

	info := (*C.VkAttachmentReferenceStencilLayoutKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR
	info.pNext = next
	info.stencilLayout = C.VkImageLayout(o.StencilLayout)

	return preallocatedPointer, nil
}

func (o AttachmentReferenceStencilLayoutOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAttachmentReferenceStencilLayoutKHR)(cDataPointer)
	return info.pNext, nil
}
