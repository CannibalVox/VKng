package khr_surface

//go:generate mockgen -source surface.go -destination ./mocks/surface.go -package mock_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Capabilities struct {
	MinImageCount int
	MaxImageCount int

	CurrentExtent  common.Extent2D
	MinImageExtent common.Extent2D
	MaxImageExtent common.Extent2D

	MaxImageArrayLayers int
	SupportedTransforms SurfaceTransforms
	CurrentTransform    SurfaceTransforms

	SupportedCompositeAlpha CompositeAlphaModes
	SupportedImageUsage     common.ImageUsages
}

type Format struct {
	Format     common.DataFormat
	ColorSpace ColorSpace
}

type vulkanSurface struct {
	instance driver.VkInstance
	handle   VkSurfaceKHR
	driver   Driver
}

type Surface interface {
	Handle() VkSurfaceKHR
	Destroy(callbacks *driver.AllocationCallbacks)
	SupportsDevice(physicalDevice core1_0.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error)
	Capabilities(device core1_0.PhysicalDevice) (*Capabilities, common.VkResult, error)
	Formats(device core1_0.PhysicalDevice) ([]Format, common.VkResult, error)
	PresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error)
}

func CreateSurface(surfacePtr unsafe.Pointer, instance core1_0.Instance, driver Driver) (Surface, common.VkResult, error) {
	return &vulkanSurface{
		handle:   (VkSurfaceKHR)(surfacePtr),
		instance: instance.Handle(),
		driver:   driver,
	}, core1_0.VKSuccess, nil
}

func (s *vulkanSurface) Handle() VkSurfaceKHR {
	return s.handle
}

func (s *vulkanSurface) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySurfaceKHR(s.instance, s.handle, callbacks.Handle())
}

func (s *vulkanSurface) SupportsDevice(physicalDevice core1_0.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error) {
	var canPresent driver.VkBool32

	res, err := s.driver.VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice.Handle(), driver.Uint32(queueFamilyIndex), s.handle, &canPresent)

	return canPresent != C.VK_FALSE, res, err
}

func (s *vulkanSurface) Capabilities(device core1_0.PhysicalDevice) (*Capabilities, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	cCapabilities := (*VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfaceCapabilitiesKHR(device.Handle(), s.handle, cCapabilities)
	if err != nil {
		return nil, res, err
	}

	return &Capabilities{
		MinImageCount: int(cCapabilities.minImageCount),
		MaxImageCount: int(cCapabilities.maxImageCount),
		CurrentExtent: common.Extent2D{
			Width:  int(cCapabilities.currentExtent.width),
			Height: int(cCapabilities.currentExtent.height),
		},
		MinImageExtent: common.Extent2D{
			Width:  int(cCapabilities.minImageExtent.width),
			Height: int(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: common.Extent2D{
			Width:  int(cCapabilities.maxImageExtent.width),
			Height: int(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: int(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransforms(cCapabilities.supportedTransforms),
		CurrentTransform:    SurfaceTransforms(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaModes(cCapabilities.supportedCompositeAlpha),
		SupportedImageUsage:     common.ImageUsages(cCapabilities.supportedUsageFlags),
	}, res, nil
}

func (s *vulkanSurface) attemptFormats(device core1_0.PhysicalDevice) ([]Format, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	formatCount := (*driver.Uint32)(formatCountPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), s.handle, formatCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, res, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))

	res, err = s.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), s.handle, formatCount, (*VkSurfaceFormatKHR)(formatsPtr))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	formatSlice := ([]VkSurfaceFormatKHR)(unsafe.Slice((*VkSurfaceFormatKHR)(formatsPtr), count))
	var result []Format
	for i := 0; i < count; i++ {
		result = append(result, Format{
			Format:     common.DataFormat(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, res, nil
}

func (s *vulkanSurface) Formats(device core1_0.PhysicalDevice) ([]Format, common.VkResult, error) {
	var formats []Format
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		formats, result, err = s.attemptFormats(device)
	}

	return formats, result, err
}

func (s *vulkanSurface) attemptPresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	modeCount := (*driver.Uint32)(modeCountPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), s.handle, modeCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, res, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	presentModes := (*VkPresentModeKHR)(modesPtr)

	res, err = s.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), s.handle, modeCount, presentModes)
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	presentModeSlice := ([]VkPresentModeKHR)(unsafe.Slice(presentModes, count))
	var result []PresentMode
	for i := 0; i < count; i++ {
		result = append(result, PresentMode(presentModeSlice[i]))
	}

	return result, res, nil
}

func (s *vulkanSurface) PresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	var presentModes []PresentMode
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		presentModes, result, err = s.attemptPresentModes(device)
	}

	return presentModes, result, err
}
