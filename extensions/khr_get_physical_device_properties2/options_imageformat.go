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

type PhysicalDeviceImageFormatInfo2 struct {
	Format core1_0.Format
	Type   core1_0.ImageType
	Tiling core1_0.ImageTiling
	Usage  core1_0.ImageUsageFlags
	Flags  core1_0.ImageCreateFlags

	common.NextOptions
}

func (o PhysicalDeviceImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImageFormatInfo2KHR{})))
	}
	info := (*C.VkPhysicalDeviceImageFormatInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info._type = C.VkImageType(o.Type)
	info.tiling = C.VkImageTiling(o.Tiling)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.flags = C.VkImageCreateFlags(o.Flags)

	return preallocatedPointer, nil
}
