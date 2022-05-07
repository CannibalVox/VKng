package khr_device_group_creation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DeviceGroupOutData struct {
	PhysicalDevices  []core1_0.PhysicalDevice
	SubsetAllocation bool

	common.HaveNext
}

func (o *DeviceGroupOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceGroupPropertiesKHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceGroupPropertiesKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceGroupOutData) PopulateOutData(cPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo := (*C.VkPhysicalDeviceGroupPropertiesKHR)(cPointer)
	o.SubsetAllocation = createInfo.subsetAllocation != C.VkBool32(0)

	instance, ok := common.OfType[core1_0.Instance](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an Instance passed to populate helpers")
	}

	count := int(createInfo.physicalDeviceCount)
	o.PhysicalDevices = make([]core1_0.PhysicalDevice, count)

	propertiesUnsafe := arena.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	for i := 0; i < count; i++ {
		handle := driver.VkPhysicalDevice(unsafe.Pointer(createInfo.physicalDevices[i]))
		instance.Driver().VkGetPhysicalDeviceProperties(handle, (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties core1_0.PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, err
		}

		version := instance.APIVersion().Min(properties.APIVersion)

		o.PhysicalDevices[i] = core.CreatePhysicalDevice(instance.Driver(), instance.Handle(), handle, version)
	}

	return createInfo.pNext, nil
}
