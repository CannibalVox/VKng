package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type ColorBlendAttachment struct {
	BlendEnabled bool

	SrcColor     VKng.BlendFactor
	DstColor     VKng.BlendFactor
	ColorBlendOp VKng.BlendOp

	SrcAlpha     VKng.BlendFactor
	DstAlpha     VKng.BlendFactor
	AlphaBlendOp VKng.BlendOp

	WriteMask VKng.ColorComponents
}

type ColorBlendOptions struct {
	LogicOpEnabled bool
	LogicOp        VKng.LogicOp

	BlendConstants [4]float32
	Attachments    []ColorBlendAttachment

	Next VKng.Options
}

func (o *ColorBlendOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineColorBlendStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineColorBlendStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.logicOpEnable = C.VK_FALSE

	if o.LogicOpEnabled {
		createInfo.logicOpEnable = C.VK_TRUE
	}

	createInfo.logicOp = C.VkLogicOp(o.LogicOp)

	for i := 0; i < 4; i++ {
		createInfo.blendConstants[i] = C.float(o.BlendConstants[i])
	}

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)
	createInfo.pAttachments = nil

	if attachmentCount > 0 {
		attachmentsPtr := (*C.VkPipelineColorBlendAttachmentState)(allocator.Malloc(attachmentCount * C.sizeof_struct_VkPipelineColorBlendAttachmentState))
		attachmentSlice := ([]C.VkPipelineColorBlendAttachmentState)(unsafe.Slice(attachmentsPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			attachmentSlice[i].blendEnable = C.VK_FALSE
			if o.Attachments[i].BlendEnabled {
				attachmentSlice[i].blendEnable = C.VK_TRUE
			}

			attachmentSlice[i].srcColorBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcColor)
			attachmentSlice[i].dstColorBlendFactor = C.VkBlendFactor(o.Attachments[i].DstColor)
			attachmentSlice[i].colorBlendOp = C.VkBlendOp(o.Attachments[i].ColorBlendOp)
			attachmentSlice[i].srcAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcAlpha)
			attachmentSlice[i].dstAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].DstAlpha)
			attachmentSlice[i].alphaBlendOp = C.VkBlendOp(o.Attachments[i].AlphaBlendOp)
			attachmentSlice[i].colorWriteMask = C.VkColorComponentFlags(o.Attachments[i].WriteMask)
		}

		createInfo.pAttachments = attachmentsPtr
	}

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
