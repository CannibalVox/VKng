package core1_0

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

type FramebufferCreateInfo struct {
	Attachments []ImageView
	Flags       FramebufferCreateFlags

	Width  int
	Height int
	Layers uint32

	RenderPass RenderPass

	common.NextOptions
}

func (o FramebufferCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkFramebufferCreateInfo)
	}
	createInfo := (*C.VkFramebufferCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	createInfo.flags = C.VkFramebufferCreateFlags(o.Flags)
	createInfo.pNext = next

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)
	if attachmentCount > 0 {
		attachmentsPtr := (*C.VkImageView)(allocator.Malloc(attachmentCount * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentsSlice := ([]C.VkImageView)(unsafe.Slice(attachmentsPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			attachmentsSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}

		createInfo.pAttachments = attachmentsPtr
	}

	createInfo.width = C.uint32_t(o.Width)
	createInfo.height = C.uint32_t(o.Height)
	createInfo.layers = C.uint32_t(o.Layers)

	return unsafe.Pointer(createInfo), nil
}
