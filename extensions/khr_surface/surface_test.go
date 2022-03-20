package khr_surface_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	core_mocks "github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/CannibalVox/VKng/extensions/khr_surface/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanSurface_PresentModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 2

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pPresentModeCount)

			presentModeSlice := ([]khr_surface.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 2))
			presentModeSlice[0] = khr_surface.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR
			presentModeSlice[1] = khr_surface.VkPresentModeKHR(3) // VK_PRESENT_MODE_FIFO_RELAXED_KHR

			return core1_0.VKSuccess, nil
		})

	presentModes, res, err := surface.PresentModes(device)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NoError(t, err)
	require.Len(t, presentModes, 2)
	require.Equal(t, khr_surface.PresentImmediate, presentModes[0])
	require.Equal(t, khr_surface.PresentFIFORelaxed, presentModes[1])
}

func TestVulkanSurface_PresentModes_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 1

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(1), *pPresentModeCount)

			presentModeSlice := ([]khr_surface.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 1))
			presentModeSlice[0] = khr_surface.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR

			return core1_0.VKIncomplete, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 2

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pPresentModeCount)

			presentModeSlice := ([]khr_surface.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 2))
			presentModeSlice[0] = khr_surface.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR
			presentModeSlice[1] = khr_surface.VkPresentModeKHR(3) // VK_PRESENT_MODE_FIFO_RELAXED_KHR

			return core1_0.VKSuccess, nil
		})

	presentModes, res, err := surface.PresentModes(device)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NoError(t, err)
	require.Len(t, presentModes, 2)
	require.Equal(t, khr_surface.PresentImmediate, presentModes[0])
	require.Equal(t, khr_surface.PresentFIFORelaxed, presentModes[1])
}

func TestVulkanSurface_SupportsDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceSupportKHR(
		device.Handle(),
		driver.Uint32(3),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, queueFamilyIndex driver.Uint32, surface khr_surface.VkSurfaceKHR, pSupport *driver.VkBool32) (common.VkResult, error) {
			*pSupport = driver.VkBool32(1)

			return core1_0.VKSuccess, nil
		})

	supports, _, err := surface.SupportsDevice(device, 3)
	require.NoError(t, err)
	require.True(t, supports)
}

func TestVulkanSurface_Capabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceCapabilitiesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pCapabilities *khr_surface.VkSurfaceCapabilitiesKHR) (common.VkResult, error) {
			val := reflect.ValueOf(pCapabilities).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("currentTransform").UnsafeAddr())) = uint32(0x00000002) // VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxImageCount").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(val.FieldByName("minImageCount").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxImageArrayLayers").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedTransforms").UnsafeAddr())) = uint32(0x00000010)     // VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedCompositeAlpha").UnsafeAddr())) = uint32(0x00000002) // VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedUsageFlags").UnsafeAddr())) = uint32(0x00000002)     // VK_IMAGE_USAGE_TRANSFER_DST_BIT

			extent := val.FieldByName("currentExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(3)

			extent = val.FieldByName("maxImageExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(17)

			extent = val.FieldByName("minImageExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(19)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(23)

			return core1_0.VKSuccess, nil
		})

	capabilities, _, err := surface.Capabilities(device)
	require.NoError(t, err)
	require.Equal(t, 1, capabilities.CurrentExtent.Width)
	require.Equal(t, 3, capabilities.CurrentExtent.Height)
	require.Equal(t, khr_surface.TransformRotate90, capabilities.CurrentTransform)
	require.Equal(t, 5, capabilities.MaxImageArrayLayers)
	require.Equal(t, 7, capabilities.MaxImageCount)
	require.Equal(t, 11, capabilities.MinImageCount)
	require.Equal(t, 13, capabilities.MaxImageExtent.Width)
	require.Equal(t, 17, capabilities.MaxImageExtent.Height)
	require.Equal(t, 19, capabilities.MinImageExtent.Width)
	require.Equal(t, 23, capabilities.MinImageExtent.Height)
	require.Equal(t, khr_surface.TransformHorizontalMirror, capabilities.SupportedTransforms)
	require.Equal(t, khr_surface.CompositeAlphaModePreMultiplied, capabilities.SupportedCompositeAlpha)
	require.Equal(t, core1_0.ImageUsageTransferDst, capabilities.SupportedImageUsage)
}

func TestVulkanSurface_Formats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pFormatCount)

			formatSlice := ([]khr_surface.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 2))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			format = val.Index(1)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(162)   // VK_FORMAT_ASTC_5x5_SRGB_BLOCK
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKSuccess, nil
		})

	formats, _, err := surface.Formats(device)
	require.NoError(t, err)
	require.Len(t, formats, 2)

	require.Equal(t, core1_0.DataFormatA2B10G10R10UnsignedNormalized, formats[0].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[0].ColorSpace)

	require.Equal(t, core1_0.DataFormatASTC5x5_sRGB, formats[1].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[1].ColorSpace)
}

func TestVulkanSurface_Formats_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := core_mocks.EasyMockInstance(ctrl, coreDriver)
	surfaceDriver := mock_surface.NewMockDriver(ctrl)
	device := core_mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	surface, _, err := khr_surface.CreateSurface(nil, instance, surfaceDriver)
	require.NoError(t, err)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = driver.Uint32(1)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(1), *pFormatCount)

			formatSlice := ([]khr_surface.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 1))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKIncomplete, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkPhysicalDevice, surface khr_surface.VkSurfaceKHR, pFormatCount *driver.Uint32, pFormats *khr_surface.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pFormatCount)

			formatSlice := ([]khr_surface.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 2))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			format = val.Index(1)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(162)   // VK_FORMAT_ASTC_5x5_SRGB_BLOCK
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKSuccess, nil
		})

	formats, _, err := surface.Formats(device)
	require.NoError(t, err)
	require.Len(t, formats, 2)

	require.Equal(t, core1_0.DataFormatA2B10G10R10UnsignedNormalized, formats[0].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[0].ColorSpace)

	require.Equal(t, core1_0.DataFormatASTC5x5_sRGB, formats[1].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[1].ColorSpace)
}
