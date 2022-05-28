package khr_imageless_framebuffer

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

type FramebufferAttachmentsCreateOptions struct {
	AttachmentImageInfos []FramebufferAttachmentImageOptions

	common.HaveNext
}

func (o FramebufferAttachmentsCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentsCreateInfoKHR{})))
	}

	info := (*C.VkFramebufferAttachmentsCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR
	info.pNext = next

	count := len(o.AttachmentImageInfos)
	info.attachmentImageInfoCount = C.uint32_t(count)
	info.pAttachmentImageInfos = nil

	if count > 0 {
		infosPtr, err := common.AllocOptionSlice[C.VkFramebufferAttachmentImageInfoKHR, FramebufferAttachmentImageOptions](allocator, o.AttachmentImageInfos)
		if err != nil {
			return nil, err
		}

		info.pAttachmentImageInfos = infosPtr
	}

	return preallocatedPointer, nil
}

func (o FramebufferAttachmentsCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkFramebufferAttachmentsCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
