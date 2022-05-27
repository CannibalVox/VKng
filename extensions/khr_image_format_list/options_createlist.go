package khr_image_format_list

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

type ImageFormatListCreateOptions struct {
	ViewFormats []common.DataFormat

	common.HaveNext
}

func (o ImageFormatListCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatListCreateInfoKHR{})))
	}

	info := (*C.VkImageFormatListCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
	info.pNext = next

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := unsafe.Slice(info.pViewFormats, count)

		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}

func (o ImageFormatListCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkImageFormatListCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
