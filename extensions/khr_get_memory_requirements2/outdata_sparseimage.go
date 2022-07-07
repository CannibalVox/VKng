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

type SparseImageMemoryRequirements2 struct {
	MemoryRequirements core1_0.SparseImageMemoryRequirements

	common.NextOutData
}

func (o *SparseImageMemoryRequirements2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageMemoryRequirements2KHR{})))
	}

	outData := (*C.VkSparseImageMemoryRequirements2KHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageMemoryRequirements2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkSparseImageMemoryRequirements2KHR)(cDataPointer)
	o.MemoryRequirements.FormatProperties.Flags = core1_0.SparseImageFormatFlags(outData.memoryRequirements.formatProperties.flags)
	o.MemoryRequirements.FormatProperties.ImageGranularity = core1_0.Extent3D{
		Width:  int(outData.memoryRequirements.formatProperties.imageGranularity.width),
		Height: int(outData.memoryRequirements.formatProperties.imageGranularity.height),
		Depth:  int(outData.memoryRequirements.formatProperties.imageGranularity.depth),
	}
	o.MemoryRequirements.FormatProperties.AspectMask = core1_0.ImageAspectFlags(outData.memoryRequirements.formatProperties.aspectMask)
	o.MemoryRequirements.ImageMipTailSize = int(outData.memoryRequirements.imageMipTailSize)
	o.MemoryRequirements.ImageMipTailStride = int(outData.memoryRequirements.imageMipTailStride)
	o.MemoryRequirements.ImageMipTailOffset = int(outData.memoryRequirements.imageMipTailOffset)
	o.MemoryRequirements.ImageMipTailFirstLod = int(outData.memoryRequirements.imageMipTailFirstLod)

	return outData.pNext, nil
}
