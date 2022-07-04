package khr_get_physical_device_properties2

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_physical_device_properties2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type Extension interface {
	PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *DeviceFeatures) error
	PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format core1_0.DataFormat, out *FormatPropertiesOutData) error
	PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options ImageFormatOptions, out *ImageFormatPropertiesOutData) (common.VkResult, error)
	PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *MemoryPropertiesOutData) error
	PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *DevicePropertiesOutData) error
	PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyOutData) ([]*QueueFamilyOutData, error)
	PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options SparseImageFormatOptions, outDataFactory func() *SparseImageFormatPropertiesOutData) ([]*SparseImageFormatPropertiesOutData, error)
}
