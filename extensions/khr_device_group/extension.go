package khr_device_group

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
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver        khr_device_group_driver.Driver
	withSurface   *VulkanExtensionWithKHRSurface
	withSwapchain *VulkanExtensionWithKHRSwapchain
}

func CreateExtensionFromDevice(device core1_0.Device, instance core1_0.Instance) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	surfaceInteraction := instance.IsInstanceExtensionActive(khr_surface.ExtensionName)
	swapchainInteraction := device.IsDeviceExtensionActive(khr_swapchain.ExtensionName)

	return CreateExtensionFromDriver(khr_device_group_driver.CreateDriverFromCore(device.Driver()), surfaceInteraction, swapchainInteraction)
}

func CreateExtensionFromDriver(driver khr_device_group_driver.Driver, khrSurfaceInteraction bool, khrSwapchainInteraction bool) *VulkanExtension {
	ext := &VulkanExtension{
		driver: driver,
	}

	if khrSurfaceInteraction {
		ext.withSurface = &VulkanExtensionWithKHRSurface{driver: driver}
	}

	if khrSwapchainInteraction {
		ext.withSwapchain = &VulkanExtensionWithKHRSwapchain{driver: driver}
	}

	return ext
}

func (v *VulkanExtension) WithKHRSurface() ExtensionWithKHRSurface {
	return v.withSurface
}

func (v *VulkanExtension) WithKHRSwapchain() ExtensionWithKHRSwapchain {
	return v.withSwapchain
}

func (v *VulkanExtension) CmdDispatchBase(commandBuffer core1_0.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	v.driver.VkCmdDispatchBaseKHR(commandBuffer.Handle(),
		driver.Uint32(baseGroupX),
		driver.Uint32(baseGroupY),
		driver.Uint32(baseGroupZ),
		driver.Uint32(groupCountX),
		driver.Uint32(groupCountY),
		driver.Uint32(groupCountZ))
}

func (v *VulkanExtension) CmdSetDeviceMask(commandBuffer core1_0.CommandBuffer, deviceMask uint32) {
	v.driver.VkCmdSetDeviceMaskKHR(commandBuffer.Handle(), driver.Uint32(deviceMask))
}

func (v *VulkanExtension) DeviceGroupPeerMemoryFeatures(device core1_0.Device, heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatures {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlagsKHR(0)))))

	v.driver.VkGetDeviceGroupPeerMemoryFeaturesKHR(
		device.Handle(),
		driver.Uint32(heapIndex),
		driver.Uint32(localDeviceIndex),
		driver.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return PeerMemoryFeatures(*featuresPtr)
}

type VulkanExtensionWithKHRSurface struct {
	driver khr_device_group_driver.Driver
}

func (v *VulkanExtensionWithKHRSurface) DeviceGroupPresentCapabilities(device core1_0.Device, outData *DeviceGroupPresentCapabilitiesOutData) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := v.driver.VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		(*khr_device_group_driver.VkDeviceGroupPresentCapabilitiesKHR)(optionPtr),
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

func (v *VulkanExtensionWithKHRSurface) DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error) {
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

func (v *VulkanExtensionWithKHRSurface) attemptGetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
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
	outRects := make([]core1_0.Rect2D, count)
	for i := 0; i < count; i++ {
		outRects[i].Offset.X = int(rectsSlice[i].offset.x)
		outRects[i].Offset.Y = int(rectsSlice[i].offset.y)
		outRects[i].Extent.Width = int(rectsSlice[i].extent.width)
		outRects[i].Extent.Height = int(rectsSlice[i].extent.height)
	}

	return outRects, res, nil
}

func (v *VulkanExtensionWithKHRSurface) PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	var outData []core1_0.Rect2D
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = v.attemptGetPhysicalDevicePresentRectangles(physicalDevice, surface)
	}
	return outData, result, err
}

type VulkanExtensionWithKHRSwapchain struct {
	driver khr_device_group_driver.Driver
}

func (v *VulkanExtensionWithKHRSwapchain) AcquireNextImage(device core1_0.Device, o AcquireNextImageOptions) (int, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return -1, core1_0.VKErrorUnknown, err
	}

	indexPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkAcquireNextImage2KHR(
		device.Handle(),
		(*khr_device_group_driver.VkAcquireNextImageInfoKHR)(optionPtr),
		indexPtr,
	)
	if err != nil {
		return -1, res, err
	}

	return int(*indexPtr), res, nil
}
