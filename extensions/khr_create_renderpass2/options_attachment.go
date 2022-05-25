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

type AttachmentDescriptionOptions struct {
	Flags          common.AttachmentDescriptionFlags
	Format         common.DataFormat
	Samples        common.SampleCounts
	LoadOp         common.AttachmentLoadOp
	StoreOp        common.AttachmentStoreOp
	StencilLoadOp  common.AttachmentLoadOp
	StencilStoreOp common.AttachmentStoreOp
	InitialLayout  common.ImageLayout
	FinalLayout    common.ImageLayout

	common.HaveNext
}

func (o AttachmentDescriptionOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescription2KHR{})))
	}

	info := (*C.VkAttachmentDescription2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
	info.pNext = next
	info.flags = C.VkAttachmentDescriptionFlags(o.Flags)
	info.format = C.VkFormat(o.Format)
	info.samples = C.VkSampleCountFlagBits(o.Samples)
	info.loadOp = C.VkAttachmentLoadOp(o.LoadOp)
	info.storeOp = C.VkAttachmentStoreOp(o.StoreOp)
	info.stencilLoadOp = C.VkAttachmentLoadOp(o.StencilLoadOp)
	info.stencilStoreOp = C.VkAttachmentStoreOp(o.StencilStoreOp)
	info.initialLayout = C.VkImageLayout(o.InitialLayout)
	info.finalLayout = C.VkImageLayout(o.FinalLayout)

	return preallocatedPointer, nil
}

func (o AttachmentDescriptionOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAttachmentDescription2KHR)(cDataPointer)
	return info.pNext, nil
}
