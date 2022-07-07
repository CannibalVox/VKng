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

type SubpassDependency2 struct {
	SrcSubpass      int
	DstSubpass      int
	SrcStageMask    core1_0.PipelineStageFlags
	DstStageMask    core1_0.PipelineStageFlags
	SrcAccessMask   core1_0.AccessFlags
	DstAccessMask   core1_0.AccessFlags
	DependencyFlags core1_0.DependencyFlags
	ViewOffset      int

	common.NextOptions
}

func (o SubpassDependency2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDependency2KHR{})))
	}

	info := (*C.VkSubpassDependency2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
	info.pNext = next
	info.srcSubpass = C.uint32_t(o.SrcSubpass)
	info.dstSubpass = C.uint32_t(o.DstSubpass)
	info.srcStageMask = C.VkPipelineStageFlags(o.SrcStageMask)
	info.dstStageMask = C.VkPipelineStageFlags(o.DstStageMask)
	info.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	info.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	info.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)
	info.viewOffset = C.int32_t(o.ViewOffset)

	return preallocatedPointer, nil
}
