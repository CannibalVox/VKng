package ext_separate_stencil_usage

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

type ImageStencilUsageCreateOptions struct {
	StencilUsage common.ImageUsages

	common.HaveNext
}

func (o ImageStencilUsageCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageStencilUsageCreateInfoEXT{})))
	}

	info := (*C.VkImageStencilUsageCreateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT
	info.pNext = next
	info.stencilUsage = C.VkImageUsageFlags(o.StencilUsage)

	return preallocatedPointer, nil
}

func (o ImageStencilUsageCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkImageStencilUsageCreateInfoEXT)(cDataPointer)
	return info.pNext, nil
}
