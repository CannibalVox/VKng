package khr_get_physical_device_properties2

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_physical_device_properties2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type Extension interface {
	PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceFeatures2) error
	PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format core1_0.Format, out *FormatProperties2) error
	PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error)
	PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceMemoryProperties2) error
	PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceProperties2) error
	PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error)
	PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)
}
