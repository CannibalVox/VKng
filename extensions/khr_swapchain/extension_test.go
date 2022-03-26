package khr_swapchain_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
	mock_surface "github.com/CannibalVox/VKng/extensions/khr_surface/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	mock_swapchain "github.com/CannibalVox/VKng/extensions/khr_swapchain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_CreateSwapchain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)

	device := mocks.EasyMockDevice(ctrl, coreDriver)
	swapchainHandle := mock_swapchain.NewFakeSwapchain()
	surface := mock_surface.EasyMockSurface(ctrl)
	oldSwapchain := mock_swapchain.EasyMockSwapchain(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = swapchainHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(1000001000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			surfaceHandle := (khr_surface_driver.VkSurfaceKHR)(unsafe.Pointer(val.FieldByName("surface").Elem().UnsafeAddr()))
			require.Equal(t, surface.Handle(), surfaceHandle)

			require.Equal(t, uint64(1), val.FieldByName("minImageCount").Uint())
			require.Equal(t, uint64(67), val.FieldByName("imageFormat").Uint())    // VK_FORMAT_A2B10G10R10_SSCALED_PACK32
			require.Equal(t, uint64(0), val.FieldByName("imageColorSpace").Uint()) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			extent := val.FieldByName("imageExtent")
			require.Equal(t, uint64(3), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(5), extent.FieldByName("height").Uint())

			require.Equal(t, uint64(7), val.FieldByName("imageArrayLayers").Uint())
			require.Equal(t, uint64(0x00000010), val.FieldByName("imageUsage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
			require.Equal(t, uint64(1), val.FieldByName("imageSharingMode").Uint())    // VK_SHARING_MODE_CONCURRENT
			require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndexCount").Uint())

			queueFamiliesPtr := (*uint32)(unsafe.Pointer(val.FieldByName("pQueueFamilyIndices").Elem().UnsafeAddr()))
			queueFamilies := ([]uint32)(unsafe.Slice(queueFamiliesPtr, 3))
			require.Equal(t, []uint32{11, 13, 17}, queueFamilies)

			require.Equal(t, uint64(0x00000020), val.FieldByName("preTransform").Uint())   // VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
			require.Equal(t, uint64(0x00000001), val.FieldByName("compositeAlpha").Uint()) // VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
			require.Equal(t, uint64(1), val.FieldByName("presentMode").Uint())             // VK_PRESENT_MODE_MAILBOX_KHR
			require.Equal(t, uint64(1), val.FieldByName("clipped").Uint())

			oldSwapchainHandle := (khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("oldSwapchain").Elem().UnsafeAddr()))
			require.Equal(t, oldSwapchain.Handle(), oldSwapchainHandle)

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.CreateOptions{
		Surface:            surface,
		MinImageCount:      1,
		ImageFormat:        core1_0.DataFormatA2B10G10R10SignedScaled,
		ImageColorSpace:    khr_surface.ColorSpaceSRGBNonlinear,
		ImageExtent:        common.Extent2D{Width: 3, Height: 5},
		ImageArrayLayers:   7,
		ImageUsage:         core1_0.ImageUsageColorAttachment,
		SharingMode:        core1_0.SharingConcurrent,
		QueueFamilyIndices: []int{11, 13, 17},
		PreTransform:       khr_surface.TransformHorizontalMirrorRotate90,
		CompositeAlpha:     khr_surface.CompositeAlphaModeOpaque,
		PresentMode:        khr_surface.PresentMailbox,
		Clipped:            true,
		OldSwapchain:       oldSwapchain,
	})

	require.NoError(t, err)
	require.Equal(t, swapchainHandle, swapchain.Handle())
}

func TestVulkanExtension_PresentToQueue_NullOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)

	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)
	queue := mocks.EasyMockQueue(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)

	swapchainDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue driver.VkQueue, pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR) (common.VkResult, error) {
			val := reflect.ValueOf(*pPresentInfo)

			require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())

			semaphorePtr := (*driver.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			semaphores := ([]driver.VkSemaphore)(unsafe.Slice(semaphorePtr, 2))
			require.Equal(t, semaphore1.Handle(), semaphores[0])
			require.Equal(t, semaphore2.Handle(), semaphores[1])

			swapchainPtr := (*khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("pSwapchains").Elem().UnsafeAddr()))
			swapchains := ([]khr_swapchain_driver.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, 1))
			require.Equal(t, swapchain.Handle(), swapchains[0])

			imageIndicesPtr := (*driver.Uint32)(unsafe.Pointer(val.FieldByName("pImageIndices").Elem().UnsafeAddr()))
			imageIndices := ([]driver.Uint32)(unsafe.Slice(imageIndicesPtr, 1))
			require.Equal(t, driver.Uint32(2), imageIndices[0])

			resultsPtr := (*driver.VkResult)(unsafe.Pointer(val.FieldByName("pResults").Elem().UnsafeAddr()))
			results := ([]driver.VkResult)(unsafe.Slice(resultsPtr, 1))
			results[0] = driver.VkResult(core1_0.VKSuccess)

			return core1_0.VKSuccess, nil
		})

	options := khr_swapchain.PresentOptions{
		WaitSemaphores: []core1_0.Semaphore{semaphore1, semaphore2},
		Swapchains:     []khr_swapchain.Swapchain{swapchain},
		ImageIndices:   []int{2},
	}
	_, err := extension.PresentToQueue(queue, options)
	require.NoError(t, err)
	require.Nil(t, options.OutData)
}

func TestVulkanExtension_PresentToQueue_RealOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)

	swapchain := mock_swapchain.EasyMockSwapchain(ctrl)
	queue := mocks.EasyMockQueue(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)

	swapchainDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue driver.VkQueue, pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR) (common.VkResult, error) {
			val := reflect.ValueOf(*pPresentInfo)

			require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())

			semaphorePtr := (*driver.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			semaphores := ([]driver.VkSemaphore)(unsafe.Slice(semaphorePtr, 2))
			require.Equal(t, semaphore1.Handle(), semaphores[0])
			require.Equal(t, semaphore2.Handle(), semaphores[1])

			swapchainPtr := (*khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("pSwapchains").Elem().UnsafeAddr()))
			swapchains := ([]khr_swapchain_driver.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, 1))
			require.Equal(t, swapchain.Handle(), swapchains[0])

			imageIndicesPtr := (*driver.Uint32)(unsafe.Pointer(val.FieldByName("pImageIndices").Elem().UnsafeAddr()))
			imageIndices := ([]driver.Uint32)(unsafe.Slice(imageIndicesPtr, 1))
			require.Equal(t, driver.Uint32(2), imageIndices[0])

			resultsPtr := (*driver.VkResult)(unsafe.Pointer(val.FieldByName("pResults").Elem().UnsafeAddr()))
			results := ([]driver.VkResult)(unsafe.Slice(resultsPtr, 1))
			results[0] = driver.VkResult(core1_0.VKTimeout)

			return core1_0.VKSuccess, nil
		})

	outData := khr_swapchain.PresentOptionsOutData{}
	_, err := extension.PresentToQueue(queue, khr_swapchain.PresentOptions{
		WaitSemaphores: []core1_0.Semaphore{semaphore1, semaphore2},
		Swapchains:     []khr_swapchain.Swapchain{swapchain},
		ImageIndices:   []int{2},
		OutData:        &outData,
	})
	require.NoError(t, err)

	require.Len(t, outData.Results, 1)
	require.Equal(t, core1_0.VKTimeout, outData.Results[0])
}
