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

type FramebufferAttachmentImageOptions struct {
	Flags      common.ImageCreateFlags
	Usage      common.ImageUsages
	Width      int
	Height     int
	LayerCount int

	ViewFormats []common.DataFormat

	common.HaveNext
}

func (o FramebufferAttachmentImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentImageInfoKHR{})))
	}

	info := (*C.VkFramebufferAttachmentImageInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
	info.pNext = next
	info.flags = C.VkImageCreateFlags(o.Flags)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.width = C.uint32_t(o.Width)
	info.height = C.uint32_t(o.Height)
	info.layerCount = C.uint32_t(o.LayerCount)

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := ([]C.VkFormat)(unsafe.Slice(info.pViewFormats, count))
		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}

func (o FramebufferAttachmentImageOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkFramebufferAttachmentImageInfoKHR)(cDataPointer)
	return info.pNext, nil
}
