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

type DeviceGroupRenderPassBeginOptions struct {
	DeviceMask        uint32
	DeviceRenderAreas []core1_0.Rect2D

	common.NextOptions
}

func (o DeviceGroupRenderPassBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupRenderPassBeginInfoKHR{})))
	}

	info := (*C.VkDeviceGroupRenderPassBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR
	info.pNext = next
	info.deviceMask = C.uint32_t(o.DeviceMask)

	count := len(o.DeviceRenderAreas)
	info.deviceRenderAreaCount = C.uint32_t(count)
	info.pDeviceRenderAreas = nil

	if count > 0 {
		areas := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		areaSlice := ([]C.VkRect2D)(unsafe.Slice(areas, count))

		for i := 0; i < count; i++ {
			areaSlice[i].offset.x = C.int32_t(o.DeviceRenderAreas[i].Offset.X)
			areaSlice[i].offset.y = C.int32_t(o.DeviceRenderAreas[i].Offset.Y)
			areaSlice[i].extent.width = C.uint32_t(o.DeviceRenderAreas[i].Extent.Width)
			areaSlice[i].extent.height = C.uint32_t(o.DeviceRenderAreas[i].Extent.Height)
		}

		info.pDeviceRenderAreas = areas
	}

	return preallocatedPointer, nil
}
