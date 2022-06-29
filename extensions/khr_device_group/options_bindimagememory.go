package khr_device_group

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

type BindImageMemoryDeviceGroupOptions struct {
	DeviceIndices            []int
	SplitInstanceBindRegions []core1_0.Rect2D

	common.HaveNext
}

func (o BindImageMemoryDeviceGroupOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemoryDeviceGroupInfoKHR{})))
	}

	info := (*C.VkBindImageMemoryDeviceGroupInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR
	info.pNext = next

	count := len(o.DeviceIndices)
	info.deviceIndexCount = C.uint32_t(count)
	info.pDeviceIndices = nil
	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.DeviceIndices[i])
		}

		info.pDeviceIndices = indices
	}

	count = len(o.SplitInstanceBindRegions)
	info.splitInstanceBindRegionCount = C.uint32_t(count)
	info.pSplitInstanceBindRegions = nil
	if count > 0 {
		regions := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		regionSlice := ([]C.VkRect2D)(unsafe.Slice(regions, count))

		for i := 0; i < count; i++ {
			regionSlice[i].offset.x = C.int32_t(o.SplitInstanceBindRegions[i].Offset.X)
			regionSlice[i].offset.y = C.int32_t(o.SplitInstanceBindRegions[i].Offset.Y)
			regionSlice[i].extent.width = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Width)
			regionSlice[i].extent.height = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Height)
		}

		info.pSplitInstanceBindRegions = regions
	}

	return preallocatedPointer, nil
}

func (o BindImageMemoryDeviceGroupOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkBindImageMemoryDeviceGroupInfoKHR)(cDataPointer)
	return info.pNext, nil
}
