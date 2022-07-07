package khr_swapchain1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
)

//go:generate mockgen -source extiface.go -destination ../mocks/extension1_1.go -package mock_swapchain -mock_names Extension=MockExtension1_1

type Extension interface {
	AcquireNextImage2(device core1_0.Device, o AcquireNextImageInfo) (int, common.VkResult, error)
	DeviceGroupPresentCapabilities(device core1_0.Device, outData *DeviceGroupPresentCapabilities) (common.VkResult, error)
	DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error)
	PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error)
}
