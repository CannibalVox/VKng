package khr_imageless_framebuffer_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_imageless_framebuffer"
	khr_imageless_framebuffer_driver "github.com/CannibalVox/VKng/extensions/khr_imageless_framebuffer/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestFramebufferAttachmentsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockFramebuffer := mocks.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCreateFramebuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkFramebufferCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pFramebuffer *driver.VkFramebuffer) (common.VkResult, error) {

		*pFramebuffer = mockFramebuffer.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO

		next := (*khr_imageless_framebuffer_driver.VkFramebufferAttachmentsCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentImageInfoCount").Uint())

		imageInfos := (*khr_imageless_framebuffer_driver.VkFramebufferAttachmentImageInfoKHR)(val.FieldByName("pAttachmentImageInfos").UnsafePointer())
		imageInfoSlice := unsafe.Slice(imageInfos, 2)
		val = reflect.ValueOf(imageInfoSlice)

		info := val.Index(0)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), info.FieldByName("flags").Uint()) // VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT
		require.Equal(t, uint64(4), info.FieldByName("usage").Uint())    // VK_IMAGE_USAGE_SAMPLED_BIT
		require.Equal(t, uint64(1), info.FieldByName("width").Uint())
		require.Equal(t, uint64(3), info.FieldByName("height").Uint())
		require.Equal(t, uint64(5), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(2), info.FieldByName("viewFormatCount").Uint())

		viewFormats := (*driver.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice := unsafe.Slice(viewFormats, 2)

		require.Equal(t, []driver.VkFormat{68, 53}, viewFormatSlice)

		info = val.Index(1)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("flags").Uint())    // VK_IMAGE_CREATE_SPARSE_BINDING_BIT
		require.Equal(t, uint64(0x10), info.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
		require.Equal(t, uint64(7), info.FieldByName("width").Uint())
		require.Equal(t, uint64(11), info.FieldByName("height").Uint())
		require.Equal(t, uint64(13), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(3), info.FieldByName("viewFormatCount").Uint())

		viewFormats = (*driver.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice = unsafe.Slice(viewFormats, 3)

		require.Equal(t, []driver.VkFormat{161, 164, 163}, viewFormatSlice)

		return core1_0.VKSuccess, nil
	})

	framebuffer, _, err := device.CreateFramebuffer(
		nil,
		core1_0.FramebufferCreateOptions{
			NextOptions: common.NextOptions{
				khr_imageless_framebuffer.FramebufferAttachmentsCreateOptions{
					AttachmentImageInfos: []khr_imageless_framebuffer.FramebufferAttachmentImageOptions{
						{
							Flags:      core1_0.ImageCreateCubeCompatible,
							Usage:      core1_0.ImageUsageSampled,
							Width:      1,
							Height:     3,
							LayerCount: 5,
							ViewFormats: []core1_0.DataFormat{
								core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
								core1_0.DataFormatA8B8G8R8UnsignedScaledPacked,
							},
						},
						{
							Flags:      core1_0.ImageCreateSparseBinding,
							Usage:      core1_0.ImageUsageColorAttachment,
							Width:      7,
							Height:     11,
							LayerCount: 13,
							ViewFormats: []core1_0.DataFormat{
								core1_0.DataFormatASTC5x5_UnsignedNormalized,
								core1_0.DataFormatASTC6x5_sRGB,
								core1_0.DataFormatASTC6x5_UnsignedNormalized,
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFramebuffer.Handle(), framebuffer.Handle())
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := extensions.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*khr_imageless_framebuffer_driver.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("imagelessFramebuffer").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice.CreateDevice(
		nil,
		core1_0.DeviceCreateOptions{
			QueueFamilies: []core1_0.DeviceQueueCreateOptions{
				{
					CreatedQueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures{
					ImagelessFramebuffer: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_imageless_framebuffer_driver.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imagelessFramebuffer").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeatures{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures{
		ImagelessFramebuffer: true,
	}, outData)
}

func TestRenderPassAttachmentBeginInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := extensions.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_0)

	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		driver.VkSubpassContents(0), // VK_SUBPASS_CONTENTS_INLINE
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		contents driver.VkSubpassContents) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO

		next := (*khr_imageless_framebuffer_driver.VkRenderPassAttachmentBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

		firstImageView := val.FieldByName("pAttachments").UnsafePointer()
		require.Equal(t, imageView1.Handle(), *(*driver.VkImageView)(firstImageView))

		secondImageView := unsafe.Add(firstImageView, unsafe.Sizeof(uintptr(0)))
		require.Equal(t, imageView2.Handle(), *(*driver.VkImageView)(secondImageView))
	})

	err := commandBuffer.CmdBeginRenderPass(core1_0.SubpassContentsInline, core1_0.RenderPassBeginOptions{
		NextOptions: common.NextOptions{khr_imageless_framebuffer.RenderPassAttachmentBeginOptions{
			Attachments: []core1_0.ImageView{imageView1, imageView2},
		}},
	})
	require.NoError(t, err)
}
