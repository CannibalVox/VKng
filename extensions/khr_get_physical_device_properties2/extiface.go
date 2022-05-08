package khr_get_physical_device_properties2

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_physical_device_properties2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type Extension interface {
	PhysicalDeviceFeatures(physicalDevice core1_0.PhysicalDevice, out *DeviceFeaturesOutData) error
	PhysicalDeviceFormatProperties(physicalDevice core1_0.PhysicalDevice, format common.DataFormat, out *FormatPropertiesOutData) error
	PhysicalDeviceImageFormatProperties(physicalDevice core1_0.PhysicalDevice, options ImageFormatOptions, out *ImageFormatPropertiesOutData) (common.VkResult, error)
	PhysicalDeviceMemoryProperties(physicalDevice core1_0.PhysicalDevice, out *MemoryPropertiesOutData) error
	PhysicalDeviceProperties(physicalDevice core1_0.PhysicalDevice, out *DevicePropertiesOutData) error
	PhysicalDeviceQueueFamilyProperties(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyOutData) ([]*QueueFamilyOutData, error)
	PhysicalDeviceSparseImageFormatProperties(physicalDevice core1_0.PhysicalDevice, options SparseImageFormatOptions, outDataFactory func() *SparseImageFormatPropertiesOutData) ([]*SparseImageFormatPropertiesOutData, error)
}
