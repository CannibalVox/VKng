package ext_surface

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SURFACE_EXTENSION_NAME

type Handle C.VkSurfaceKHR
type Surface struct {
	instance C.VkInstance
	handle   C.VkSurfaceKHR
}

func (s *Surface) Handle() Handle {
	return Handle(s.handle)
}

func (s *Surface) Destroy() {
	C.vkDestroySurfaceKHR(s.instance, s.handle, nil)
}

func (s *Surface) SupportsDevice(physicalDevice *core.PhysicalDevice, queueFamilyIndex int) (bool, VKng.Result, error) {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	res := VKng.Result(C.vkGetPhysicalDeviceSurfaceSupportKHR(deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent))

	return canPresent != C.VK_FALSE, res, res.ToError()
}

func (s *Surface) Capabilities(allocator cgoalloc.Allocator, device *core.PhysicalDevice) (*Capabilities, VKng.Result, error) {
	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	defer allocator.Free(capabilitiesPtr)

	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res := VKng.Result(C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR((C.VkPhysicalDevice)(unsafe.Pointer(device.Handle())), s.handle, cCapabilities))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Capabilities{
		MinImageCount: uint32(cCapabilities.minImageCount),
		MaxImageCount: uint32(cCapabilities.maxImageCount),
		CurrentExtent: VKng.Extent2D{
			Width:  uint32(cCapabilities.currentExtent.width),
			Height: uint32(cCapabilities.currentExtent.height),
		},
		MinImageExtent: VKng.Extent2D{
			Width:  uint32(cCapabilities.minImageExtent.width),
			Height: uint32(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: VKng.Extent2D{
			Width:  uint32(cCapabilities.maxImageExtent.width),
			Height: uint32(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: uint32(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransforms(cCapabilities.supportedTransforms),
		CurrentTransform:    SurfaceTransforms(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaModes(cCapabilities.supportedCompositeAlpha),
		SupportedImageUsage:     VKng.ImageUsages(cCapabilities.supportedUsageFlags),
	}, res, nil
}

func (s *Surface) Formats(allocator cgoalloc.Allocator, device *core.PhysicalDevice) ([]Format, VKng.Result, error) {
	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(formatCountPtr)

	formatCount := (*C.uint32_t)(formatCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := VKng.Result(C.vkGetPhysicalDeviceSurfaceFormatsKHR(deviceHandle, s.handle, formatCount, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, res, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))
	defer allocator.Free(formatsPtr)

	res = VKng.Result(C.vkGetPhysicalDeviceSurfaceFormatsKHR(deviceHandle, s.handle, formatCount, (*C.VkSurfaceFormatKHR)(formatsPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	formatSlice := ([]C.VkSurfaceFormatKHR)(unsafe.Slice((*C.VkSurfaceFormatKHR)(formatsPtr), count))
	var result []Format
	for i := 0; i < count; i++ {
		result = append(result, Format{
			Format:     VKng.DataFormat(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, res, nil
}

func (s *Surface) PresentModes(allocator cgoalloc.Allocator, device *core.PhysicalDevice) ([]PresentMode, VKng.Result, error) {
	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(modeCountPtr)

	modeCount := (*C.uint32_t)(modeCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := VKng.Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(deviceHandle, s.handle, modeCount, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, res, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	defer allocator.Free(modesPtr)

	presentModes := (*C.VkPresentModeKHR)(modesPtr)

	res = VKng.Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(deviceHandle, s.handle, modeCount, presentModes))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	presentModeSlice := ([]C.VkPresentModeKHR)(unsafe.Slice(presentModes, count))
	var result []PresentMode
	for i := 0; i < count; i++ {
		result = append(result, PresentMode(presentModeSlice[i]))
	}

	return result, res, nil
}
