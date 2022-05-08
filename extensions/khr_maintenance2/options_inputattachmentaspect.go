package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type RenderPassInputAttachmentAspectOptions struct {
	AspectReferences []InputAttachmentAspectReference

	common.HaveNext
}

func (o RenderPassInputAttachmentAspectOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassInputAttachmentAspectCreateInfoKHR{})))
	}

	createInfo := (*C.VkRenderPassInputAttachmentAspectCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.AspectReferences)
	if count < 1 {
		return nil, errors.New("options RenderPassInputAttachmentAspectOptions must include at least 1 entry in AspectReferences")
	}

	createInfo.aspectReferenceCount = C.uint32_t(count)
	references, err := common.AllocSlice[C.VkInputAttachmentAspectReference, InputAttachmentAspectReference](allocator, o.AspectReferences)
	if err != nil {
		return nil, err
	}
	createInfo.pAspectReferences = references

	return preallocatedPointer, nil
}

func (o RenderPassInputAttachmentAspectOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkRenderPassInputAttachmentAspectCreateInfoKHR)(cDataPointer)
	return createInfo.pNext, nil
}

var _ common.Options = RenderPassInputAttachmentAspectOptions{}
