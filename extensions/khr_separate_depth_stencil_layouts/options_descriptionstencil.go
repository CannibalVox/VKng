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

type AttachmentDescriptionStencilLayoutOptions struct {
	StencilInitialLayout common.ImageLayout
	StencilFinalLayout   common.ImageLayout

	common.HaveNext
}

func (o AttachmentDescriptionStencilLayoutOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescriptionStencilLayoutKHR{})))
	}

	info := (*C.VkAttachmentDescriptionStencilLayoutKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR
	info.pNext = next
	info.stencilInitialLayout = C.VkImageLayout(o.StencilInitialLayout)
	info.stencilFinalLayout = C.VkImageLayout(o.StencilFinalLayout)

	return preallocatedPointer, nil
}

func (o AttachmentDescriptionStencilLayoutOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAttachmentDescriptionStencilLayoutKHR)(cDataPointer)
	return info.pNext, nil
}
