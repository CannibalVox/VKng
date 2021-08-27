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
	"github.com/CannibalVox/VKng/objects"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SURFACE_EXTENSION_NAME

type Handle C.VkSurfaceKHR
type Surface struct {
	instance C.VkInstance
	handle C.VkSurfaceKHR
}

func (s *Surface) Handle() Handle {
	return Handle(s.handle)
}

func (s *Surface) Destroy() {
	C.vkDestroySurfaceKHR(s.instance, s.handle, nil)
}

func (s *Surface) SupportsDevice(physicalDevice *objects.PhysicalDevice, queueFamilyIndex int) bool {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	C.vkGetPhysicalDeviceSurfaceSupportKHR(deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent)

	return canPresent != C.VK_FALSE
}

func (s *Surface) Capabilities(allocator cgoalloc.Allocator, device *objects.PhysicalDevice) (*Capabilities, error) {
	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	defer allocator.Free(capabilitiesPtr)

	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res := C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR((C.VkPhysicalDevice)(unsafe.Pointer(device.Handle())), s.handle, cCapabilities)
	err := VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Capabilities{
		MinImageCount: uint32(cCapabilities.minImageCount),
		MaxImageCount: uint32(cCapabilities.maxImageCount),
		CurrentExtent: VKng.Extent2D{
			Width: uint32(cCapabilities.currentExtent.width),
			Height: uint32(cCapabilities.currentExtent.height),
		},
		MinImageExtent: VKng.Extent2D{
			Width: uint32(cCapabilities.minImageExtent.width),
			Height: uint32(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: VKng.Extent2D{
			Width: uint32(cCapabilities.maxImageExtent.width),
			Height: uint32(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: uint32(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransforms(cCapabilities.supportedTransforms),
		CurrentTransform: SurfaceTransforms(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaModes(cCapabilities.supportedCompositeAlpha),
		SupportedImageUsage:     VKng.ImageUsage(cCapabilities.supportedUsageFlags),
	}, nil
}

func (s *Surface) Formats(allocator cgoalloc.Allocator, device *objects.PhysicalDevice) ([]*SurfaceFormat, error) {
	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(formatCountPtr)

	formatCount := (*C.uint32_t)(formatCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := C.vkGetPhysicalDeviceSurfaceFormatsKHR(deviceHandle, s.handle, formatCount, nil)
	err := VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))
	defer allocator.Free(formatsPtr)

	res = C.vkGetPhysicalDeviceSurfaceFormatsKHR(deviceHandle, s.handle, formatCount, (*C.VkSurfaceFormatKHR)(formatsPtr))
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	formatSlice := ([]C.VkSurfaceFormatKHR)(unsafe.Slice((*C.VkSurfaceFormatKHR)(formatsPtr), count))
	var result []*SurfaceFormat
	for i := 0; i < count; i++ {
		result = append(result, &SurfaceFormat{
			Format: VKng.ColorFormat(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, nil
}

func (s *Surface) PresentModes(allocator cgoalloc.Allocator, device *objects.PhysicalDevice) ([]PresentMode, error) {
	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(modeCountPtr)

	modeCount := (*C.uint32_t)(modeCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := C.vkGetPhysicalDeviceSurfacePresentModesKHR(deviceHandle, s.handle, modeCount, nil)
	err := VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	defer allocator.Free(modesPtr)
	
	presentModes := (*C.VkPresentModeKHR)(modesPtr)

	res = C.vkGetPhysicalDeviceSurfacePresentModesKHR(deviceHandle, s.handle, modeCount, presentModes)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	presentModeSlice := ([]C.VkPresentModeKHR)(unsafe.Slice(presentModes, count))
	var result []PresentMode
	for i := 0; i < count; i++ {
		result = append(result, PresentMode(presentModeSlice[i]))
	}

	return result, nil
}
