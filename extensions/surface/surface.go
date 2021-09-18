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
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
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
	SupportsDevice(physicalDevice core.PhysicalDevice, queueFamilyIndex int) (bool, core.VkResult, error)
	Capabilities(device core.PhysicalDevice) (*Capabilities, core.VkResult, error)
	Formats(device core.PhysicalDevice) ([]Format, core.VkResult, error)
	PresentModes(device core.PhysicalDevice) ([]PresentMode, core.VkResult, error)
}

func (s *vulkanSurface) Handle() Handle {
	return Handle(s.handle)
}

func (s *vulkanSurface) Destroy() {
	C.cgoDestroySurfaceKHR(s.destroyFunc, s.instance, s.handle, nil)
}

func (s *vulkanSurface) SupportsDevice(physicalDevice core.PhysicalDevice, queueFamilyIndex int) (bool, core.VkResult, error) {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(s.physicalSurfaceSupportFunc, deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent))

	return canPresent != C.VK_FALSE, res, res.ToError()
}

func (s *vulkanSurface) Capabilities(device core.PhysicalDevice) (*Capabilities, core.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(s.physicalSurfaceCapabilitiesFunc, (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle())), s.handle, cCapabilities))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Capabilities{
		MinImageCount: uint32(cCapabilities.minImageCount),
		MaxImageCount: uint32(cCapabilities.maxImageCount),
		CurrentExtent: common.Extent2D{
			Width:  uint32(cCapabilities.currentExtent.width),
			Height: uint32(cCapabilities.currentExtent.height),
		},
		MinImageExtent: common.Extent2D{
			Width:  uint32(cCapabilities.minImageExtent.width),
			Height: uint32(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: common.Extent2D{
			Width:  uint32(cCapabilities.maxImageExtent.width),
			Height: uint32(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: uint32(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransforms(cCapabilities.supportedTransforms),
		CurrentTransform:    SurfaceTransforms(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaModes(cCapabilities.supportedCompositeAlpha),
		SupportedImageUsage:     common.ImageUsages(cCapabilities.supportedUsageFlags),
	}, res, nil
}

func (s *vulkanSurface) Formats(device core.PhysicalDevice) ([]Format, core.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	formatCount := (*C.uint32_t)(formatCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(s.surfaceFormatsFunc, deviceHandle, s.handle, formatCount, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, res, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))

	res = core.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(s.surfaceFormatsFunc, deviceHandle, s.handle, formatCount, (*C.VkSurfaceFormatKHR)(formatsPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	formatSlice := ([]C.VkSurfaceFormatKHR)(unsafe.Slice((*C.VkSurfaceFormatKHR)(formatsPtr), count))
	var result []Format
	for i := 0; i < count; i++ {
		result = append(result, Format{
			Format:     common.DataFormat(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, res, nil
}

func (s *vulkanSurface) PresentModes(device core.PhysicalDevice) ([]PresentMode, core.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	modeCount := (*C.uint32_t)(modeCountPtr)
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(device.Handle()))

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(s.presentModesFunc, deviceHandle, s.handle, modeCount, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, res, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	presentModes := (*C.VkPresentModeKHR)(modesPtr)

	res = core.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(s.presentModesFunc, deviceHandle, s.handle, modeCount, presentModes))
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

func CreateSurface(surfacePtr unsafe.Pointer, instance core.Instance) (Surface, core.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	instanceHandle := (C.VkInstance)(unsafe.Pointer(instance.Handle()))
	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkDestroySurfaceKHR"))))

	return &vulkanSurface{
		handle:   (C.VkSurfaceKHR)(surfacePtr),
		instance: instanceHandle,

		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}, core.VKSuccess, nil
}
