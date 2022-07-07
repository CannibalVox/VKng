package khr_get_physical_device_properties2

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

type PhysicalDeviceSparseImageFormatInfo2 struct {
	Format  core1_0.Format
	Type    core1_0.ImageType
	Samples core1_0.SampleCountFlags
	Usage   core1_0.ImageUsageFlags
	Tiling  core1_0.ImageTiling

	common.NextOptions
}

func (o PhysicalDeviceSparseImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSparseImageFormatInfo2KHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceSparseImageFormatInfo2KHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR
	createInfo.pNext = next
	createInfo.format = C.VkFormat(o.Format)
	createInfo._type = C.VkImageType(o.Type)
	createInfo.samples = C.VkSampleCountFlagBits(o.Samples)
	createInfo.usage = C.VkImageUsageFlags(o.Usage)
	createInfo.tiling = C.VkImageTiling(o.Tiling)

	return preallocatedPointer, nil
}
