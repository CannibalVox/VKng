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

type AttachmentReference2 struct {
	Attachment int
	Layout     core1_0.ImageLayout
	AspectMask core1_0.ImageAspectFlags

	common.NextOptions
}

func (o AttachmentReference2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReference2KHR{})))
	}

	info := (*C.VkAttachmentReference2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
	info.pNext = next
	info.attachment = C.uint32_t(o.Attachment)
	info.layout = C.VkImageLayout(o.Layout)
	info.aspectMask = C.VkImageAspectFlags(o.AspectMask)

	return preallocatedPointer, nil
}
