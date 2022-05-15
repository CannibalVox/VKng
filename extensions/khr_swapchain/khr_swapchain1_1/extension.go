package khr_swapchain1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_device_group_driver "github.com/CannibalVox/VKng/extensions/khr_device_group/driver"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver khr_swapchain_driver.Driver
}

func CreateExtensionFromDriver(driver khr_swapchain_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (v *VulkanExtension) AcquireNextImage(device core1_0.Device, o AcquireNextImageOptions) (int, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return -1, core1_0.VKErrorUnknown, err
	}

	indexPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkAcquireNextImage2KHR(
		device.Handle(),
		(*khr_swapchain_driver.VkAcquireNextImageInfoKHR)(optionPtr),
		indexPtr,
	)
	if err != nil {
		return -1, res, err
	}

	return int(*indexPtr), res, nil
}

func (v *VulkanExtension) DeviceGroupPresentCapabilities(device core1_0.Device, outData *DeviceGroupPresentCapabilitiesOutData) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := v.driver.VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		(*khr_swapchain_driver.VkDeviceGroupPresentCapabilitiesKHR)(optionPtr),
	)
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(outData, optionPtr)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (v *VulkanExtension) DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	flagsPtr := (*khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR)(arena.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupPresentModeFlagsKHR(0)))))

	res, err := v.driver.VkGetDeviceGroupSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		flagsPtr,
	)
	if err != nil {
		return 0, res, err
	}

	return DeviceGroupPresentModeFlags(*flagsPtr), res, nil
}

func (v *VulkanExtension) attemptGetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]common.Rect2D, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		countPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*countPtr)
	if count == 0 {
		return nil, res, nil
	}

	rectsPtr := arena.Malloc(count * C.sizeof_struct_VkRect2D)
	res, err = v.driver.VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		(*driver.Uint32)(countPtr),
		(*driver.VkRect2D)(rectsPtr),
	)
	if res != core1_0.VKSuccess {
		return nil, res, err
	}

	rectsSlice := ([]C.VkRect2D)(unsafe.Slice((*C.VkRect2D)(rectsPtr), count))
	outRects := make([]common.Rect2D, count)
	for i := 0; i < count; i++ {
		outRects[i].Offset.X = int(rectsSlice[i].offset.x)
		outRects[i].Offset.Y = int(rectsSlice[i].offset.y)
		outRects[i].Extent.Width = int(rectsSlice[i].extent.width)
		outRects[i].Extent.Height = int(rectsSlice[i].extent.height)
	}

	return outRects, res, nil
}

func (v *VulkanExtension) PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]common.Rect2D, common.VkResult, error) {
	var outData []common.Rect2D
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = v.attemptGetPhysicalDevicePresentRectangles(physicalDevice, surface)
	}
	return outData, result, err
}
