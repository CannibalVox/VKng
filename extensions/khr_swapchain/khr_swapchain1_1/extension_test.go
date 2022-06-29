package khr_swapchain1_1_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
	mock_surface "github.com/CannibalVox/VKng/extensions/khr_surface/mocks"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain/khr_swapchain1_1"
	mock_swapchain "github.com/CannibalVox/VKng/extensions/khr_swapchain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestVulkanExtension_GetDeviceGroupPresentCapabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain1_1.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		info *khr_swapchain_driver.VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {

		val := reflect.ValueOf(info).Elem()
		require.Equal(t, uint64(1000060007), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		mask := val.FieldByName("presentMask")
		*(*uint32)(unsafe.Pointer(mask.Index(0).UnsafeAddr())) = uint32(1)
		*(*uint32)(unsafe.Pointer(mask.Index(1).UnsafeAddr())) = uint32(2)
		*(*uint32)(unsafe.Pointer(mask.Index(2).UnsafeAddr())) = uint32(7)
		for i := 3; i < 32; i++ {
			*(*uint32)(unsafe.Pointer(mask.Index(i).UnsafeAddr())) = uint32(0)
		}
		*(*uint32)(unsafe.Pointer(val.FieldByName("modes").UnsafeAddr())) = 0

		return core1_0.VKSuccess, nil
	})

	var outData khr_swapchain1_1.DeviceGroupPresentCapabilitiesOutData
	_, err := extension.DeviceGroupPresentCapabilities(
		device,
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_swapchain1_1.DeviceGroupPresentCapabilitiesOutData{
		PresentMask: [32]uint32{1, 2, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, outData)
}

func TestVulkanExtension_GetDeviceGroupSurfacePresentModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain1_1.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	extDriver.EXPECT().VkGetDeviceGroupSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pModes *khr_swapchain_driver.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {

		*pModes = khr_swapchain_driver.VkDeviceGroupPresentModeFlagsKHR(4) // VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	modes, _, err := extension.DeviceGroupSurfacePresentModes(device, surface)
	require.NoError(t, err)
	require.Equal(t, khr_swapchain1_1.DeviceGroupPresentModeSum, modes)
}

func TestVulkanExtension_GetPhysicalDevicePresentRectangles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain1_1.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pRectCount)

		rectSlice := ([]driver.VkRect2D)(unsafe.Slice(pRects, 3))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		r = val.Index(2)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(23)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(29)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(31)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(37)

		return core1_0.VKSuccess, nil
	})

	rects, _, err := extension.PhysicalDevicePresentRectangles(physicalDevice, surface)
	require.NoError(t, err)
	require.Equal(t, []core1_0.Rect2D{
		{
			Offset: core1_0.Offset2D{X: 1, Y: 3},
			Extent: core1_0.Extent2D{Width: 5, Height: 7},
		},
		{
			Offset: core1_0.Offset2D{X: 11, Y: 13},
			Extent: core1_0.Extent2D{Width: 17, Height: 19},
		},
		{
			Offset: core1_0.Offset2D{X: 23, Y: 29},
			Extent: core1_0.Extent2D{Width: 31, Height: 37},
		},
	}, rects)
}

func TestVulkanExtension_GetPhysicalDevicePresentRectangles_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain1_1.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = driver.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(2), *pRectCount)

		rectSlice := ([]driver.VkRect2D)(unsafe.Slice(pRects, 2))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		return core1_0.VKIncomplete, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *driver.Uint32,
		pRects *driver.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pRectCount)

		rectSlice := ([]driver.VkRect2D)(unsafe.Slice(pRects, 3))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		r = val.Index(2)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(23)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(29)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(31)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(37)

		return core1_0.VKSuccess, nil
	})

	rects, _, err := extension.PhysicalDevicePresentRectangles(physicalDevice, surface)
	require.NoError(t, err)
	require.Equal(t, []core1_0.Rect2D{
		{
			Offset: core1_0.Offset2D{X: 1, Y: 3},
			Extent: core1_0.Extent2D{Width: 5, Height: 7},
		},
		{
			Offset: core1_0.Offset2D{X: 11, Y: 13},
			Extent: core1_0.Extent2D{Width: 17, Height: 19},
		},
		{
			Offset: core1_0.Offset2D{X: 23, Y: 29},
			Extent: core1_0.Extent2D{Width: 31, Height: 37},
		},
	}, rects)
}

func TestVulkanExtensionWithKHRSwapchain_AcquireNextImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain1_1.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)
	semaphore := mocks.EasyMockSemaphore(ctrl)

	extDriver.EXPECT().VkAcquireNextImage2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pAcquireInfo *khr_swapchain_driver.VkAcquireNextImageInfoKHR,
		pImageIndex *driver.Uint32,
	) (common.VkResult, error) {
		*pImageIndex = driver.Uint32(2)

		val := reflect.ValueOf(pAcquireInfo).Elem()
		require.Equal(t, uint64(1000060010), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), khr_swapchain_driver.VkSwapchainKHR(val.FieldByName("swapchain").UnsafePointer()))
		require.Equal(t, uint64(1000000000), val.FieldByName("timeout").Uint())
		require.Equal(t, semaphore.Handle(), driver.VkSemaphore(val.FieldByName("semaphore").UnsafePointer()))
		require.True(t, val.FieldByName("fence").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	index, _, err := extension.AcquireNextImage(
		device,
		khr_swapchain1_1.AcquireNextImageOptions{
			Swapchain:  swapchain,
			Timeout:    time.Second,
			Semaphore:  semaphore,
			DeviceMask: 3,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 2, index)
}
