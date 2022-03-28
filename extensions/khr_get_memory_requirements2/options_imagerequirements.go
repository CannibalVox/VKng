package khr_get_memory_requirements2

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

type ImageRequirementsOptions struct {
	Image core1_0.Image

	common.HaveNext
}

func (o ImageRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageMemoryRequirementsInfo2KHR{})))
	}

	options := (*C.VkImageMemoryRequirementsInfo2KHR)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

func (o ImageRequirementsOptions) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	options := (*C.VkImageMemoryRequirementsInfo2KHR)(cPointer)
	return options.pNext, nil
}
