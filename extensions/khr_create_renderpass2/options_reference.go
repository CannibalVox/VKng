package khr_create_renderpass2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type AttachmentReferenceOptions struct {
	Attachment int
	Layout     common.ImageLayout
	AspectMask common.ImageAspectFlags

	common.HaveNext
}

func (o AttachmentReferenceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o AttachmentReferenceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAttachmentReference2KHR)(cDataPointer)
	return info.pNext, nil
}
