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
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver Driver
}

type Extension interface {
	CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options *CreationOptions) (Swapchain, common.VkResult, error)
	PresentToQueue(queue core1_0.Queue, o *PresentOptions) (common.VkResult, error)
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options *CreationOptions) (Swapchain, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var swapchain VkSwapchainKHR

	res, err := l.driver.VkCreateSwapchainKHR(device.Handle(), (*VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSwapchain{
		handle:            swapchain,
		device:            device.Handle(),
		driver:            l.driver,
		minimumAPIVersion: device.APIVersion(),
	}, res, nil
}

func (s *VulkanExtension) PresentToQueue(queue core1_0.Queue, o *PresentOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*VkPresentInfoKHR)(createInfo)
	res, err := s.driver.VkQueuePresentKHR(queue.Handle(), createInfoPtr)
	popErr := common.PopulateOutData(o, createInfo)

	if popErr != nil {
		return core1_0.VKErrorUnknown, popErr
	} else if err != nil {
		return res, err
	}

	return res, res.ToError()
}
