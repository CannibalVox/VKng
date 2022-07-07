package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type AttachmentDescriptionStencilLayout struct {
	StencilInitialLayout core1_0.ImageLayout
	StencilFinalLayout   core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescriptionStencilLayout) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
