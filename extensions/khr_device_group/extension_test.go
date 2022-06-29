package khr_device_group_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/driver"
	mock_bind_memory2 "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_device_group"
	khr_device_group_driver "github.com/CannibalVox/VKng/extensions/khr_device_group/driver"
	mock_device_group "github.com/CannibalVox/VKng/extensions/khr_device_group/mocks"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
	mock_surface "github.com/CannibalVox/VKng/extensions/khr_surface/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	mock_swapchain "github.com/CannibalVox/VKng/extensions/khr_swapchain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestVulkanExtension_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, false, false)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	extDriver.EXPECT().VkCmdDispatchBaseKHR(
		commandBuffer.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Uint32(5),
		driver.Uint32(7),
		driver.Uint32(11),
		driver.Uint32(13),
	)

	extension.CmdDispatchBase(commandBuffer, 1, 3, 5, 7, 11, 13)
}

func TestVulkanExtension_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, false, false)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	extDriver.EXPECT().VkCmdSetDeviceMaskKHR(commandBuffer.Handle(), driver.Uint32(3))

	extension.CmdSetDeviceMask(commandBuffer, 3)
}

func TestVulkanExtension_GetDeviceGroupPeerMemoryFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, false, false)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceGroupPeerMemoryFeaturesKHR(
		device.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Uint32(5),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		heapIndex, localDeviceIndex, remoteDeviceIndex driver.Uint32,
		pPeerMemoryFeatures *khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR,
	) {
		*pPeerMemoryFeatures = khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR(1) // VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR
	})

	features := extension.DeviceGroupPeerMemoryFeatures(
		device,
		1, 3, 5,
	)
	require.Equal(t, khr_device_group.PeerMemoryFeatureCopySrc, features)
}

func TestVulkanExtension_WithKHRSurface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, true, true)
	require.NotNil(t, extension.WithKHRSurface())
	require.NotNil(t, extension.WithKHRSwapchain())
}

func TestVulkanExtension_WithKHRSurface_None(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, false, false)
	require.Nil(t, extension.WithKHRSurface())
	require.Nil(t, extension.WithKHRSwapchain())
}

func TestVulkanExtensionWithKHRSurface_GetDeviceGroupPresentCapabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, true, false)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		info *khr_device_group_driver.VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {

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

	var outData khr_device_group.DeviceGroupPresentCapabilitiesOutData
	_, err := extension.WithKHRSurface().DeviceGroupPresentCapabilities(
		device,
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_device_group.DeviceGroupPresentCapabilitiesOutData{
		PresentMask: [32]uint32{1, 2, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, outData)
}

func TestVulkanExtensionWithKHRSurface_GetDeviceGroupSurfacePresentModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, true, false)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	extDriver.EXPECT().VkGetDeviceGroupSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pModes *khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {

		*pModes = khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR(4) // VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	modes, _, err := extension.WithKHRSurface().DeviceGroupSurfacePresentModes(device, surface)
	require.NoError(t, err)
	require.Equal(t, khr_device_group.DeviceGroupPresentModeSum, modes)
}

func TestVulkanExtensionWithKHRSurface_GetPhysicalDevicePresentRectangles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, true, false)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
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

	rects, _, err := extension.WithKHRSurface().PhysicalDevicePresentRectangles(physicalDevice, surface)
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

func TestVulkanExtensionWithKHRSurface_GetPhysicalDevicePresentRectangles_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, true, false)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
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

	rects, _, err := extension.WithKHRSurface().PhysicalDevicePresentRectangles(physicalDevice, surface)
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

	extDriver := mock_device_group.NewMockDriver(ctrl)
	extension := khr_device_group.CreateExtensionFromDriver(extDriver, false, true)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)
	semaphore := mocks.EasyMockSemaphore(ctrl)

	extDriver.EXPECT().VkAcquireNextImage2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pAcquireInfo *khr_device_group_driver.VkAcquireNextImageInfoKHR,
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

	index, _, err := extension.WithKHRSwapchain().AcquireNextImage(
		device,
		khr_device_group.AcquireNextImageOptions{
			Swapchain:  swapchain,
			Timeout:    time.Second,
			Semaphore:  semaphore,
			DeviceMask: 3,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 2, index)
}

func TestMemoryAllocateFlagsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	mockMemory := mocks.EasyMockDeviceMemory(ctrl)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)

	coreDriver.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pAllocateInfo *driver.VkMemoryAllocateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pMemory *driver.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*khr_device_group_driver.VkMemoryAllocateFlagsInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000060000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) //VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR
			require.Equal(t, uint64(5), val.FieldByName("deviceMask").Uint())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := device.AllocateMemory(nil,
		core1_0.MemoryAllocateOptions{
			AllocationSize:  1,
			MemoryTypeIndex: 3,
			HaveNext: common.HaveNext{Next: khr_device_group.MemoryAllocateFlagsOptions{
				Flags:      khr_device_group.MemoryAllocateDeviceMask,
				DeviceMask: 5,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestDeviceGroupCommandBufferBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	mockCommandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	commandBuffer := extensions.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mockCommandBuffer.Handle(), common.Vulkan1_0)

	coreDriver.EXPECT().VkBeginCommandBuffer(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer, pBeginInfo *driver.VkCommandBufferBeginInfo) (common.VkResult, error) {
		val := reflect.ValueOf(pBeginInfo).Elem()

		require.Equal(t, uint64(42), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint())  // VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
		require.True(t, val.FieldByName("pInheritanceInfo").IsNil())

		next := (*khr_device_group_driver.VkDeviceGroupCommandBufferBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := commandBuffer.Begin(core1_0.BeginOptions{
		Flags: core1_0.BeginInfoOneTimeSubmit,
		HaveNext: common.HaveNext{Next: khr_device_group.DeviceGroupCommandBufferBeginOptions{
			DeviceMask: 3,
		}},
	})
	require.NoError(t, err)
}

func TestBindBufferMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	buffer := mocks.EasyMockBuffer(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkBindBufferMemory2KHR(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		infoCount driver.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
		require.Equal(t, buffer.Handle(), (driver.VkBuffer)(val.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindBufferMemoryDeviceGroupInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060013), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*driver.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 3))
		val = reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), val.Index(0).Uint())
		require.Equal(t, uint64(2), val.Index(1).Uint())
		require.Equal(t, uint64(7), val.Index(2).Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindBufferMemory(device, []khr_bind_memory2.BindBufferMemoryOptions{
		{
			Buffer:       buffer,
			Memory:       memory,
			MemoryOffset: 1,

			HaveNext: common.HaveNext{
				khr_device_group.BindBufferMemoryDeviceGroupOptions{
					DeviceIndices: []int{1, 2, 7},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestBindImageMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		infoCount driver.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), (driver.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindImageMemoryDeviceGroupInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060014), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*driver.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 3))
		indexVal := reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), indexVal.Index(0).Uint())
		require.Equal(t, uint64(2), indexVal.Index(1).Uint())
		require.Equal(t, uint64(7), indexVal.Index(2).Uint())

		require.Equal(t, uint64(2), val.FieldByName("splitInstanceBindRegionCount").Uint())

		regions := (*driver.VkRect2D)(val.FieldByName("pSplitInstanceBindRegions").UnsafePointer())
		regionSlice := ([]driver.VkRect2D)(unsafe.Slice(regions, 2))
		regionVal := reflect.ValueOf(regionSlice)

		oneRegion := regionVal.Index(0)
		require.Equal(t, int64(3), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(5), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(7), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(11), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		oneRegion = regionVal.Index(1)
		require.Equal(t, int64(13), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(17), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(19), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(23), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory(device, []khr_bind_memory2.BindImageMemoryOptions{
		{
			Image:        image,
			Memory:       memory,
			MemoryOffset: 1,

			HaveNext: common.HaveNext{
				khr_device_group.BindImageMemoryDeviceGroupOptions{
					DeviceIndices: []int{1, 2, 7},
					SplitInstanceBindRegions: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 3, Y: 5},
							Extent: core1_0.Extent2D{Width: 7, Height: 11},
						},
						{
							Offset: core1_0.Offset2D{X: 13, Y: 17},
							Extent: core1_0.Extent2D{Width: 19, Height: 23},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestBindImageMemorySwapchainOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)
	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		infoCount driver.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), (driver.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindImageMemorySwapchainInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060009), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("swapchain").UnsafePointer()))
		require.Equal(t, uint64(3), val.FieldByName("imageIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory(device, []khr_bind_memory2.BindImageMemoryOptions{
		{
			Image:        image,
			Memory:       memory,
			MemoryOffset: 1,

			HaveNext: common.HaveNext{
				khr_device_group.BindImageMemorySwapchainOptions{
					Swapchain:  swapchain,
					ImageIndex: 3,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestDeviceGroupBindSparseOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockQueue := mocks.EasyMockQueue(ctrl)
	fence := mocks.EasyMockFence(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	queue := extensions.CreateQueueObject(coreDriver, device.Handle(), mockQueue.Handle(), common.Vulkan1_0)

	coreDriver.EXPECT().VkQueueBindSparse(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(
		queue driver.VkQueue,
		optionCount driver.Uint32,
		pSparseOptions *driver.VkBindSparseInfo,
		fence driver.VkFence,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pSparseOptions).Elem()

		require.Equal(t, uint64(7), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, semaphore1.Handle(), driver.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		semaphores := (*driver.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*khr_device_group_driver.VkDeviceGroupBindSparseInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("resourceDeviceIndex").Uint())
		require.Equal(t, uint64(3), val.FieldByName("memoryDeviceIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := queue.BindSparse(fence, []core1_0.BindSparseOptions{
		{
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			HaveNext: common.HaveNext{
				khr_device_group.DeviceGroupBindSparseOptions{
					ResourceDeviceIndex: 1,
					MemoryDeviceIndex:   3,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestImageSwapchainCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockImage := mocks.EasyMockImage(ctrl)
	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *driver.VkImageCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pImage *driver.VkImage,
	) (common.VkResult, error) {
		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
		require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

		next := (*khr_device_group_driver.VkImageSwapchainCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("swapchain").UnsafePointer()))

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(nil, core1_0.ImageCreateOptions{
		MipLevels:   1,
		ArrayLayers: 3,
		HaveNext: common.HaveNext{
			khr_device_group.ImageSwapchainCreateOptions{
				Swapchain: swapchain,
			},
		},
	})
	require.Equal(t, mockImage.Handle(), image.Handle())
	require.NoError(t, err)
}

func TestDeviceGroupPresentOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(extDriver)

	queue := mocks.EasyMockQueue(ctrl)
	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)

	extDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		queue driver.VkQueue,
		pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pPresentInfo).Elem()
		require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
		require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("pSwapchains").Elem().UnsafePointer()))
		require.Equal(t, uint64(3), val.FieldByName("pImageIndices").Elem().Uint())

		next := (*khr_device_group_driver.VkDeviceGroupPresentInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060011), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())
		require.Equal(t, uint64(7), val.FieldByName("pDeviceMasks").Elem().Uint())
		require.Equal(t, uint64(4), val.FieldByName("mode").Uint()) // VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	_, err := extension.PresentToQueue(queue, khr_swapchain.PresentOptions{
		Swapchains:   []khr_swapchain.Swapchain{swapchain},
		ImageIndices: []int{3},
		HaveNext: common.HaveNext{
			khr_device_group.DeviceGroupPresentOptions{
				DeviceMasks: []uint32{7},
				Mode:        khr_device_group.DeviceGroupPresentModeSum,
			},
		},
	})
	require.NoError(t, err)
}

func TestDeviceGroupRenderPassBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	mockCommandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	commandBuffer := extensions.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mockCommandBuffer.Handle(), common.Vulkan1_0)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	framebuffer := mocks.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		driver.VkSubpassContents(1), // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	).DoAndReturn(func(
		commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		contents driver.VkSubpassContents,
	) {
		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.Equal(t, renderPass.Handle(), (driver.VkRenderPass)(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), (driver.VkFramebuffer)(val.FieldByName("framebuffer").UnsafePointer()))

		next := (*khr_device_group_driver.VkDeviceGroupRenderPassBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.FieldByName("deviceMask").Uint())
		require.Equal(t, uint64(2), val.FieldByName("deviceRenderAreaCount").Uint())

		areas := (*driver.VkRect2D)(val.FieldByName("pDeviceRenderAreas").UnsafePointer())
		areaSlice := ([]driver.VkRect2D)(unsafe.Slice(areas, 2))
		val = reflect.ValueOf(areaSlice)

		oneArea := val.Index(0)
		require.Equal(t, int64(1), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), oneArea.FieldByName("extent").FieldByName("height").Uint())

		oneArea = val.Index(1)
		require.Equal(t, int64(11), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(13), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(17), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(19), oneArea.FieldByName("extent").FieldByName("height").Uint())
	})

	err := commandBuffer.CmdBeginRenderPass(
		core1_0.SubpassContentsSecondaryCommandBuffers,
		core1_0.RenderPassBeginOptions{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			HaveNext: common.HaveNext{
				khr_device_group.DeviceGroupRenderPassBeginOptions{
					DeviceMask: 7,
					DeviceRenderAreas: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 1, Y: 3},
							Extent: core1_0.Extent2D{Width: 5, Height: 7},
						},
						{
							Offset: core1_0.Offset2D{X: 11, Y: 13},
							Extent: core1_0.Extent2D{Width: 17, Height: 19},
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestDeviceGroupSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockQueue := mocks.EasyMockQueue(ctrl)
	fence := mocks.EasyMockFence(ctrl)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	queue := extensions.CreateQueueObject(coreDriver, device.Handle(), mockQueue.Handle(), common.Vulkan1_0)

	coreDriver.EXPECT().VkQueueSubmit(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue driver.VkQueue, submitCount driver.Uint32, pSubmits *driver.VkSubmitInfo, fence driver.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), driver.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), driver.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*driver.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*khr_device_group_driver.VkDeviceGroupSubmitInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, uint64(1), val.FieldByName("pWaitSemaphoreDeviceIndices").Elem().Uint())
		require.Equal(t, uint64(2), val.FieldByName("pCommandBufferDeviceMasks").Elem().Uint())

		indices := (*driver.Uint32)(val.FieldByName("pSignalSemaphoreDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 2))
		require.Equal(t, []driver.Uint32{3, 5}, indexSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := queue.SubmitToQueue(fence, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			WaitDstStages:    []core1_0.PipelineStages{core1_0.PipelineStageBottomOfPipe},

			HaveNext: common.HaveNext{
				khr_device_group.DeviceGroupSubmitOptions{
					WaitSemaphoreDeviceIndices:   []int{1},
					CommandBufferDeviceMasks:     []uint32{2},
					SignalSemaphoreDeviceIndices: []int{3, 5},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestDeviceGroupSwapchainCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)
	mockSwapchain := mock_swapchain.EasyMockSwapchain(ctrl)

	extDriver.EXPECT().VkCreateSwapchainKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR,
			pAllocator *driver.VkAllocationCallbacks,
			pSwapchain *khr_swapchain_driver.VkSwapchainKHR,
		) (common.VkResult, error) {
			*pSwapchain = mockSwapchain.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000001000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
			require.Equal(t, surface.Handle(), khr_surface_driver.VkSurfaceKHR(val.FieldByName("surface").UnsafePointer()))

			next := (*khr_device_group_driver.VkDeviceGroupSwapchainCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000060012), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("modes").Uint()) // VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(
		device,
		nil,
		khr_swapchain.CreateOptions{
			Surface: surface,
			HaveNext: common.HaveNext{
				khr_device_group.DeviceGroupSwapchainCreateOptions{
					Modes: khr_device_group.DeviceGroupPresentModeLocal,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockSwapchain.Handle(), swapchain.Handle())
}
