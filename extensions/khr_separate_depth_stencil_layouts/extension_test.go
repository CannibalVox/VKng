package khr_separate_depth_stencil_layouts_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_create_renderpass2"
	khr_create_renderpass2_driver "github.com/CannibalVox/VKng/extensions/khr_create_renderpass2/driver"
	mock_create_renderpass2 "github.com/CannibalVox/VKng/extensions/khr_create_renderpass2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_separate_depth_stencil_layouts"
	khr_separate_depth_stencil_layouts_driver "github.com/CannibalVox/VKng/extensions/khr_separate_depth_stencil_layouts/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestAttachmentDescriptionStencilLayoutOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)

	extDriver.EXPECT().VkCreateRenderPass2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *khr_create_renderpass2_driver.VkRenderPassCreateInfo2KHR,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("attachmentCount").Uint())

		attachment := val.FieldByName("pAttachments").Elem()
		require.Equal(t, uint64(1000109000), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
		attachmentNext := (*khr_separate_depth_stencil_layouts_driver.VkAttachmentDescriptionStencilLayoutKHR)(attachment.FieldByName("pNext").UnsafePointer())

		attachmentLayout := reflect.ValueOf(attachmentNext).Elem()
		require.Equal(t, uint64(1000241002), attachmentLayout.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR
		require.True(t, attachmentLayout.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1000241000), attachmentLayout.FieldByName("stencilInitialLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR
		require.Equal(t, uint64(1000241003), attachmentLayout.FieldByName("stencilFinalLayout").Uint())   // VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR

		subpass := val.FieldByName("pSubpasses").Elem()
		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
		require.True(t, subpass.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())

		inputAttachment := subpass.FieldByName("pInputAttachments").Elem()
		require.Equal(t, uint64(1000109001), inputAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR

		inputAttachmentNext := (*khr_separate_depth_stencil_layouts_driver.VkAttachmentReferenceStencilLayoutKHR)(inputAttachment.FieldByName("pNext").UnsafePointer())
		stencilRef := reflect.ValueOf(inputAttachmentNext).Elem()
		require.Equal(t, uint64(1000241001), stencilRef.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR
		require.True(t, stencilRef.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1000241000), stencilRef.FieldByName("stencilLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := extension.CreateRenderPass2(
		device,
		nil,
		khr_create_renderpass2.RenderPassCreateOptions{
			Attachments: []khr_create_renderpass2.AttachmentDescriptionOptions{
				{
					HaveNext: common.HaveNext{khr_separate_depth_stencil_layouts.AttachmentDescriptionStencilLayoutOptions{
						StencilInitialLayout: khr_separate_depth_stencil_layouts.ImageLayoutDepthAttachmentOptimal,
						StencilFinalLayout:   khr_separate_depth_stencil_layouts.ImageLayoutStencilReadOnlyOptimal,
					}},
				},
			},
			Subpasses: []khr_create_renderpass2.SubpassDescriptionOptions{
				{
					InputAttachments: []khr_create_renderpass2.AttachmentReferenceOptions{
						{
							HaveNext: common.HaveNext{
								khr_separate_depth_stencil_layouts.AttachmentReferenceStencilLayoutOptions{
									StencilLayout: khr_separate_depth_stencil_layouts.ImageLayoutDepthAttachmentOptimal,
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())
}

func TestPhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions(t *testing.T) {
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

		next := (*khr_separate_depth_stencil_layouts_driver.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000241000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("separateDepthStencilLayouts").Uint())

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
			HaveNext: common.HaveNext{
				khr_separate_depth_stencil_layouts.PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions{
					SeparateDepthStencilLayouts: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData(t *testing.T) {
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

		next := (*khr_separate_depth_stencil_layouts_driver.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000241000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("separateDepthStencilLayouts").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData khr_separate_depth_stencil_layouts.PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeaturesOutData{
			HaveNext: common.HaveNext{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_separate_depth_stencil_layouts.PhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData{
		SeparateDepthStencilLayouts: true,
	}, outData)
}
