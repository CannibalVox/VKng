package khr_device_group

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

type DeviceGroupPresentOptions struct {
	DeviceMasks []uint32
	Mode        DeviceGroupPresentModeFlags

	common.HaveNext
}

func (o DeviceGroupPresentOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceGroupPresentInfoKHR)
	}

	info := (*C.VkDeviceGroupPresentInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
	info.pNext = next
	info.mode = C.VkDeviceGroupPresentModeFlagBitsKHR(o.Mode)

	count := len(o.DeviceMasks)
	info.swapchainCount = C.uint32_t(count)
	info.pDeviceMasks = nil

	if count > 0 {
		masks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		maskSlice := ([]C.uint32_t)(unsafe.Slice(masks, count))

		for i := 0; i < count; i++ {
			maskSlice[i] = C.uint32_t(o.DeviceMasks[i])
		}
		info.pDeviceMasks = masks
	}

	return preallocatedPointer, nil
}

func (o DeviceGroupPresentOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceGroupPresentInfoKHR)(cDataPointer)
	return info.pNext, nil
}
