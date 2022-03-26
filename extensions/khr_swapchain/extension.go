package khr_swapchain

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_swapchain_driver.Driver
}

type Extension interface {
	CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options CreateOptions) (Swapchain, common.VkResult, error)
	PresentToQueue(queue core1_0.Queue, o PresentOptions) (common.VkResult, error)
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: khr_swapchain_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_swapchain_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options CreateOptions) (Swapchain, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var swapchain khr_swapchain_driver.VkSwapchainKHR

	res, err := l.driver.VkCreateSwapchainKHR(device.Handle(), (*khr_swapchain_driver.VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return nil, res, err
	}

	coreDriver := device.Driver()
	newSwapchain := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(swapchain), func() interface{} {
		return &vulkanSwapchain{
			handle:            swapchain,
			device:            device.Handle(),
			driver:            l.driver,
			minimumAPIVersion: device.APIVersion(),
			coreDriver:        coreDriver,
		}
	}).(*vulkanSwapchain)
	return newSwapchain, res, nil
}

func (s *VulkanExtension) PresentToQueue(queue core1_0.Queue, o PresentOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*khr_swapchain_driver.VkPresentInfoKHR)(createInfo)
	res, err := s.driver.VkQueuePresentKHR(queue.Handle(), createInfoPtr)
	popErr := common.PopulateOutData(o, createInfo)

	if popErr != nil {
		return core1_0.VKErrorUnknown, popErr
	} else if err != nil {
		return res, err
	}

	return res, res.ToError()
}
