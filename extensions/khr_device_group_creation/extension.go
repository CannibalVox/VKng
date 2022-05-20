package khr_device_group_creation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_device_group_creation_driver "github.com/CannibalVox/VKng/extensions/khr_device_group_creation/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver khr_device_group_creation_driver.Driver
}

func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_device_group_creation_driver.CreateDriverFromCore(instance.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_device_group_creation_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) attemptEnumeratePhysicalDeviceGroups(instance core1_0.Instance, outDataFactory func() *DeviceGroupOutData) ([]*DeviceGroupOutData, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := e.driver.VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		countPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*countPtr)
	if count == 0 {
		return nil, core1_0.VKSuccess, nil
	}

	outDataSlice := make([]*DeviceGroupOutData, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &DeviceGroupOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkPhysicalDeviceGroupPropertiesKHR, *DeviceGroupOutData](arena, outDataSlice)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	res, err = e.driver.VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		countPtr,
		(*khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Pointer(outData)),
	)
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupPropertiesKHR, *DeviceGroupOutData](outDataSlice, unsafe.Pointer(outData), instance)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (e *VulkanExtension) PhysicalDeviceGroups(instance core1_0.Instance, outDataFactory func() *DeviceGroupOutData) ([]*DeviceGroupOutData, common.VkResult, error) {
	var outData []*DeviceGroupOutData
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = e.attemptEnumeratePhysicalDeviceGroups(instance, outDataFactory)
	}
	return outData, result, err
}
