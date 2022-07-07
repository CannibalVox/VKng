package khr_device_group_creation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DeviceGroupDeviceCreateInfo struct {
	PhysicalDevices []core1_0.PhysicalDevice

	common.NextOptions
}

func (o DeviceGroupDeviceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupDeviceCreateInfoKHR{})))
	}

	if len(o.PhysicalDevices) < 1 {
		return nil, errors.New("must include at least one physical device in DeviceGroupDeviceCreateInfo")
	}

	createInfo := (*C.VkDeviceGroupDeviceCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.PhysicalDevices)
	createInfo.physicalDeviceCount = C.uint32_t(count)
	physicalDevicesPtr := (*C.VkPhysicalDevice)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkPhysicalDevice{}))))
	physicalDevicesSlice := ([]C.VkPhysicalDevice)(unsafe.Slice(physicalDevicesPtr, count))

	for i := 0; i < count; i++ {
		physicalDevicesSlice[i] = C.VkPhysicalDevice(unsafe.Pointer(o.PhysicalDevices[i].Handle()))
	}
	createInfo.pPhysicalDevices = physicalDevicesPtr
	return preallocatedPointer, nil
}
