package khr_device_group

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_device_group

type Extension interface {
	CmdDispatchBase(commandBuffer core1_0.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	CmdSetDeviceMask(commandBuffer core1_0.CommandBuffer, deviceMask uint32)
	DeviceGroupPeerMemoryFeatures(device core1_0.Device, heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatures

	WithKHRSurface() ExtensionWithKHRSurface
	WithKHRSwapchain() ExtensionWithKHRSwapchain
}

type ExtensionWithKHRSurface interface {
	DeviceGroupPresentCapabilities(device core1_0.Device, outData *DeviceGroupPresentCapabilities) (common.VkResult, error)
	DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error)
	PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error)
}

type ExtensionWithKHRSwapchain interface {
	AcquireNextImage2(device core1_0.Device, o AcquireNextImageInfo) (int, common.VkResult, error)
}
