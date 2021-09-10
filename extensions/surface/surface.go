package ext_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

VkResult cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return fn(physicalDevice, surface, pSurfaceCapabilities);
}

VkResult cgoGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fn, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return fn(physicalDevice, queueFamilyIndex, surface, pSupported);
}

void cgoDestroySurfaceKHR(PFN_vkDestroySurfaceKHR fn, VkInstance instance, VkSurfaceKHR surface, VkAllocationCallbacks* pAllocator) {
	fn(instance, surface, pAllocator);
}

VkResult cgoGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fn,VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return fn(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}

VkResult cgoGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return fn(physicalDevice, surface, pPresentModeCount, pPresentModes);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SURFACE_EXTENSION_NAME

type Handle C.VkSurfaceKHR
type Surface struct {
	instance C.VkInstance
	handle   C.VkSurfaceKHR

	physicalSurfaceCapabilitiesFunc C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	physicalSurfaceSupportFunc      C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR
	surfaceFormatsFunc              C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR
	presentModesFunc                C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR
	destroyFunc                     C.PFN_vkDestroySurfaceKHR
}

func (s *Surface) Handle() Handle {
	return Handle(s.handle)
}

func (s *Surface) Destroy() {
	C.cgoDestroySurfaceKHR(s.destroyFunc, s.instance, s.handle, nil)
}

func (s *Surface) SupportsDevice(physicalDevice *resource.PhysicalDevice, queueFamilyIndex int) (bool, loader.VkResult, error) {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	res := loader.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(s.physicalSurfaceSupportFunc, deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent))

	return canPresent != C.VK_FALSE, res, res.ToError()
}

func (s *Surface) Capabilities(allocator cgoalloc.Allocator, device *resource.PhysicalDevice) (*Capabilities, loader.VkResult, error) {
	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	defer allocator.Free(capabilitiesPtr)

	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res := loader.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(s.physicalSurfaceCapabilitiesFunc, (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle())), s.handle, cCapabilities))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Capabilities{
		MinImageCount: uint32(cCapabilities.minImageCount),
		MaxImageCount: uint32(cCapabilities.maxImageCount),
		CurrentExtent: core.Extent2D{
			Width:  uint32(cCapabilities.currentExtent.width),
			Height: uint32(cCapabilities.currentExtent.height),
		},
		MinImageExtent: core.Extent2D{
			Width:  uint32(cCapabilities.minImageExtent.width),
			Height: uint32(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: core.Extent2D{
			Width:  uint32(cCapabilities.maxImageExtent.width),
			Height: uint32(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: uint32(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransforms(cCapabilities.supportedTransforms),
		CurrentTransform:    SurfaceTransforms(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaModes(cCapabilities.supportedCompositeAlpha),
		SupportedImageUsage:     core.ImageUsages(cCapabilities.supportedUsageFlags),
	}, res, nil
}

func (s *Surface) Formats(allocator cgoalloc.Allocator, device *resource.PhysicalDevice) ([]Format, loader.VkResult, error) {
	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(formatCountPtr)

	formatCount := (*C.uint32_t)(formatCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := loader.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(s.surfaceFormatsFunc, deviceHandle, s.handle, formatCount, nil))
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

	res = loader.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(s.surfaceFormatsFunc, deviceHandle, s.handle, formatCount, (*C.VkSurfaceFormatKHR)(formatsPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	formatSlice := ([]C.VkSurfaceFormatKHR)(unsafe.Slice((*C.VkSurfaceFormatKHR)(formatsPtr), count))
	var result []Format
	for i := 0; i < count; i++ {
		result = append(result, Format{
			Format:     core.DataFormat(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, res, nil
}

func (s *Surface) PresentModes(allocator cgoalloc.Allocator, device *resource.PhysicalDevice) ([]PresentMode, loader.VkResult, error) {
	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(modeCountPtr)

	modeCount := (*C.uint32_t)(modeCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := loader.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(s.presentModesFunc, deviceHandle, s.handle, modeCount, nil))
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

	res = loader.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(s.presentModesFunc, deviceHandle, s.handle, modeCount, presentModes))
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

func buildSurface(allocator *cgoalloc.ArenaAllocator, instance *resource.Instance, surfaceHandle C.VkSurfaceKHR) *Surface {
	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))
	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkDestroySurfaceKHR"))))

	return &Surface{
		handle:   surfaceHandle,
		instance: instanceHandle,

		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}
}
