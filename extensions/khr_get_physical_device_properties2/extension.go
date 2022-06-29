package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	ext_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver ext_driver.Driver
}

func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: ext_driver.CreateDriverFromCore(instance.Driver()),
	}
}

func CreateExtensionFromDriver(driver ext_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *DeviceFeaturesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceFeatures2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceFeatures2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format core1_0.DataFormat, out *FormatPropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice.Handle(), driver.VkFormat(format), (*ext_driver.VkFormatProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options ImageFormatOptions, out *ImageFormatPropertiesOutData) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := e.driver.VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceImageFormatInfo2KHR)(optionData), (*ext_driver.VkImageFormatProperties2KHR)(outData))
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(out, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (e *VulkanExtension) PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *MemoryPropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceMemoryProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *DevicePropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyOutData) ([]*QueueFamilyOutData, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*QueueFamilyOutData, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &QueueFamilyOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkQueueFamilyProperties2KHR, *QueueFamilyOutData](arena, out)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), outDataCountPtr, (*ext_driver.VkQueueFamilyProperties2KHR)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2KHR, *QueueFamilyOutData](out, unsafe.Pointer(outData))
	return out, err
}

func (e *VulkanExtension) PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options SparseImageFormatOptions, outDataFactory func() *SparseImageFormatPropertiesOutData) ([]*SparseImageFormatPropertiesOutData, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	optionData, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR)(optionData), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*SparseImageFormatPropertiesOutData, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &SparseImageFormatPropertiesOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkSparseImageFormatProperties2KHR, *SparseImageFormatPropertiesOutData](arena, out)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR)(optionData), outDataCountPtr, (*ext_driver.VkSparseImageFormatProperties2KHR)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2KHR, *SparseImageFormatPropertiesOutData](out, unsafe.Pointer(outData))

	return out, err
}

var _ Extension = &VulkanExtension{}
