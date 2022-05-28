package khr_imageless_framebuffer

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

type RenderPassAttachmentBeginOptions struct {
	Attachments []core1_0.ImageView

	common.HaveNext
}

func (o RenderPassAttachmentBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassAttachmentBeginInfoKHR{})))
	}

	info := (*C.VkRenderPassAttachmentBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR
	info.pNext = next

	count := len(o.Attachments)
	info.attachmentCount = C.uint32_t(count)
	info.pAttachments = nil

	if count > 0 {
		info.pAttachments = (*C.VkImageView)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentSlice := unsafe.Slice(info.pAttachments, count)
		for i := 0; i < count; i++ {
			attachmentSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}
	}

	return preallocatedPointer, nil
}

func (o RenderPassAttachmentBeginOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkRenderPassAttachmentBeginInfoKHR)(cDataPointer)
	return info.pNext, nil
}
