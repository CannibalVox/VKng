package ext_surface

//go:generate mockgen -source surface.go -destination ./mocks/mocks.go -package mock_surface

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
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SURFACE_EXTENSION_NAME

type Handle C.VkSurfaceKHR
type vulkanSurface struct {
	instance C.VkInstance
	handle   C.VkSurfaceKHR

	physicalSurfaceCapabilitiesFunc C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	physicalSurfaceSupportFunc      C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR
	surfaceFormatsFunc              C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR
	presentModesFunc                C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR
	destroyFunc                     C.PFN_vkDestroySurfaceKHR
}

type Surface interface {
	Handle() Handle
	Destroy()
	SupportsDevice(physicalDevice resources.PhysicalDevice, queueFamilyIndex int) (bool, loader.VkResult, error)
	Capabilities(allocator cgoalloc.Allocator, device resources.PhysicalDevice) (*Capabilities, loader.VkResult, error)
	Formats(allocator cgoalloc.Allocator, device resources.PhysicalDevice) ([]Format, loader.VkResult, error)
	PresentModes(allocator cgoalloc.Allocator, device resources.PhysicalDevice) ([]PresentMode, loader.VkResult, error)
}

func (s *vulkanSurface) Handle() Handle {
	return Handle(s.handle)
}

func (s *vulkanSurface) Destroy() {
	C.cgoDestroySurfaceKHR(s.destroyFunc, s.instance, s.handle, nil)
}

func (s *vulkanSurface) SupportsDevice(physicalDevice resources.PhysicalDevice, queueFamilyIndex int) (bool, loader.VkResult, error) {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	res := loader.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(s.physicalSurfaceSupportFunc, deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent))

	return canPresent != C.VK_FALSE, res, res.ToError()
}

func (s *vulkanSurface) Capabilities(allocator cgoalloc.Allocator, device resources.PhysicalDevice) (*Capabilities, loader.VkResult, error) {
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

func (s *vulkanSurface) Formats(allocator cgoalloc.Allocator, device resources.PhysicalDevice) ([]Format, loader.VkResult, error) {
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

func (s *vulkanSurface) PresentModes(allocator cgoalloc.Allocator, device resources.PhysicalDevice) ([]PresentMode, loader.VkResult, error) {
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

func CreateSurface(allocator cgoalloc.Allocator, surfacePtr unsafe.Pointer, instance resources.Instance) (Surface, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))
	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(allocator, "vkDestroySurfaceKHR"))))

	return &vulkanSurface{
		handle:   (C.VkSurfaceKHR)(surfacePtr),
		instance: instanceHandle,

		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}, loader.VKSuccess, nil
}
