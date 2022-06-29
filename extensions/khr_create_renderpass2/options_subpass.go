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
	"github.com/cockroachdb/errors"
	"unsafe"
)

type SubpassDescriptionOptions struct {
	Flags                  core1_0.SubPassDescriptionFlags
	PipelineBindPoint      core1_0.PipelineBindPoint
	ViewMask               uint32
	InputAttachments       []AttachmentReferenceOptions
	ColorAttachments       []AttachmentReferenceOptions
	ResolveAttachments     []AttachmentReferenceOptions
	DepthStencilAttachment *AttachmentReferenceOptions
	PreserveAttachments    []int

	common.HaveNext
}

func (o SubpassDescriptionOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescription2KHR{})))
	}

	info := (*C.VkSubpassDescription2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
	info.pNext = next
	info.flags = C.VkSubpassDescriptionFlags(o.Flags)
	info.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	info.viewMask = C.uint32_t(o.ViewMask)

	inputAttachmentCount := len(o.InputAttachments)
	colorAttachmentCount := len(o.ColorAttachments)
	resolveAttachmentCount := len(o.ResolveAttachments)
	preserveAttachmentCount := len(o.PreserveAttachments)

	if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
		return nil, errors.Newf("in this subpass, %d color attachments are defined, but %d resolve attachments are defined- they should be equal", colorAttachmentCount, resolveAttachmentCount)
	}

	info.inputAttachmentCount = C.uint32_t(inputAttachmentCount)
	info.pInputAttachments = nil
	info.colorAttachmentCount = C.uint32_t(colorAttachmentCount)
	info.pColorAttachments = nil
	info.pResolveAttachments = nil
	info.pDepthStencilAttachment = nil
	info.preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
	info.pPreserveAttachments = nil

	var err error
	if inputAttachmentCount > 0 {
		info.pInputAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReferenceOptions](allocator, o.InputAttachments)
		if err != nil {
			return nil, err
		}
	}

	if colorAttachmentCount > 0 {
		info.pColorAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReferenceOptions](allocator, o.ColorAttachments)
		if err != nil {
			return nil, err
		}

		info.pResolveAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReferenceOptions](allocator, o.ResolveAttachments)
		if err != nil {
			return nil, err
		}
	}

	if o.DepthStencilAttachment != nil {
		depthStencilPtr, err := common.AllocOptions(allocator, o.DepthStencilAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilAttachment = (*C.VkAttachmentReference2KHR)(depthStencilPtr)
	}

	if preserveAttachmentCount > 0 {
		attachmentsPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		attachmentsSlice := unsafe.Slice(attachmentsPtr, preserveAttachmentCount)
		for i := 0; i < preserveAttachmentCount; i++ {
			attachmentsSlice[i] = C.uint32_t(o.PreserveAttachments[i])
		}
		info.pPreserveAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}

func (o SubpassDescriptionOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSubpassDescription2KHR)(cDataPointer)
	return info.pNext, nil
}
