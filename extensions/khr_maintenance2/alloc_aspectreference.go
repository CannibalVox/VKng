package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type InputAttachmentAspectReference struct {
	Subpass              int
	InputAttachmentIndex int
	AspectMask           common.ImageAspectFlags
}

func (ref InputAttachmentAspectReference) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkInputAttachmentAspectReferenceKHR{})))
	}

	val := (*C.VkInputAttachmentAspectReferenceKHR)(preallocatedPointer)
	val.subpass = C.uint32_t(ref.Subpass)
	val.inputAttachmentIndex = C.uint32_t(ref.InputAttachmentIndex)
	val.aspectMask = C.VkImageAspectFlags(ref.AspectMask)

	return preallocatedPointer, nil
}
