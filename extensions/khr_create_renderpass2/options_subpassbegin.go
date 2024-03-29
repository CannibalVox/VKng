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
	"unsafe"
)

type SubpassBeginInfo struct {
	Contents core1_0.SubpassContents

	common.NextOptions
}

func (o SubpassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassBeginInfoKHR{})))
	}

	info := (*C.VkSubpassBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
	info.pNext = next
	info.contents = C.VkSubpassContents(o.Contents)

	return preallocatedPointer, nil
}
