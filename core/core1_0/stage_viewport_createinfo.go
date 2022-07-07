package core1_0

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

type PipelineViewportStateCreateFlags uint32

var pipelineViewportStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineViewportStateCreateFlags]()

func (f PipelineViewportStateCreateFlags) Register(str string) {
	pipelineViewportStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineViewportStateCreateFlags) String() string {
	return pipelineViewportStateCreateFlagsMapping.FlagsToString(f)
}

////

type PipelineViewportStateCreateInfo struct {
	Viewports []Viewport
	Scissors  []Rect2D
	Flags     PipelineViewportStateCreateFlags

	common.NextOptions
}

func (o PipelineViewportStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineViewportStateCreateInfo)
	}
	createInfo := (*C.VkPipelineViewportStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineViewportStateCreateFlags(o.Flags)
	createInfo.pNext = next

	viewportCount := len(o.Viewports)
	scissorsCount := len(o.Scissors)

	createInfo.viewportCount = C.uint(viewportCount)
	if viewportCount > 0 {
		viewportPtr := (*C.VkViewport)(allocator.Malloc(viewportCount * C.sizeof_struct_VkViewport))
		viewportSlice := ([]C.VkViewport)(unsafe.Slice(viewportPtr, viewportCount))
		for i := 0; i < viewportCount; i++ {
			viewportSlice[i].x = C.float(o.Viewports[i].X)
			viewportSlice[i].y = C.float(o.Viewports[i].Y)
			viewportSlice[i].width = C.float(o.Viewports[i].Width)
			viewportSlice[i].height = C.float(o.Viewports[i].Height)
			viewportSlice[i].minDepth = C.float(o.Viewports[i].MinDepth)
			viewportSlice[i].maxDepth = C.float(o.Viewports[i].MaxDepth)
		}
		createInfo.pViewports = viewportPtr
	}

	createInfo.scissorCount = C.uint(scissorsCount)
	if scissorsCount > 0 {
		scissorPtr := (*C.VkRect2D)(allocator.Malloc(scissorsCount * C.sizeof_struct_VkRect2D))
		scissorSlice := ([]C.VkRect2D)(unsafe.Slice(scissorPtr, scissorsCount))
		for i := 0; i < scissorsCount; i++ {
			scissorSlice[i].offset.x = C.int32_t(o.Scissors[i].Offset.X)
			scissorSlice[i].offset.y = C.int32_t(o.Scissors[i].Offset.Y)
			scissorSlice[i].extent.width = C.uint32_t(o.Scissors[i].Extent.Width)
			scissorSlice[i].extent.height = C.uint32_t(o.Scissors[i].Extent.Height)
		}
		createInfo.pScissors = scissorPtr
	}

	return preallocatedPointer, nil
}
