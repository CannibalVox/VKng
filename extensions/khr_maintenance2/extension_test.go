package khr_maintenance2

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
	khr_maintenance2_driver "github.com/CannibalVox/VKng/extensions/khr_maintenance2/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestImageViewUsageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	image := mocks.EasyMockImage(ctrl)
	expectedImageView := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCreateImageView(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkImageViewCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImageView *driver.VkImageView) (common.VkResult, error) {
			*pImageView = expectedImageView.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.Equal(t, image.Handle(), (driver.VkImage)(val.FieldByName("image").UnsafePointer()))

			viewUsagePtr := (*khr_maintenance2_driver.VkImageViewUsageCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			viewUsage := reflect.ValueOf(viewUsagePtr).Elem()
			require.Equal(t, uint64(1000117002), viewUsage.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR
			require.True(t, viewUsage.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000080), viewUsage.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT

			return core1_0.VKSuccess, nil
		})

	imageView, _, err := device.CreateImageView(nil, core1_0.ImageViewCreateOptions{
		Image: image,
		NextOptions: common.NextOptions{Next: ImageViewUsageOptions{
			Usage: core1_0.ImageUsageInputAttachment,
		}},
	})

	require.NoError(t, err)
	require.Equal(t, expectedImageView.Handle(), imageView.Handle())
}

func TestTessellationDomainOriginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	expectedPipeline := mocks.EasyMockPipeline(ctrl)

	coreDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(0), driver.Uint32(1), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pPipelines, 1))
			pipelineSlice[0] = expectedPipeline.Handle()

			createInfoSlice := ([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1))
			val := reflect.ValueOf(createInfoSlice[0])

			require.Equal(t, uint64(28), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			tessellation := (*driver.VkPipelineTessellationStateCreateInfo)(val.FieldByName("pTessellationState").UnsafePointer())
			tessVal := reflect.ValueOf(tessellation).Elem()

			require.Equal(t, uint64(21), tessVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
			require.Equal(t, uint64(1), tessVal.FieldByName("patchControlPoints").Uint())

			domain := (*khr_maintenance2_driver.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(tessVal.FieldByName("pNext").UnsafePointer())
			domainVal := reflect.ValueOf(domain).Elem()

			require.Equal(t, uint64(1000117003), domainVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR
			require.True(t, domainVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), domainVal.FieldByName("domainOrigin").Uint())

			return core1_0.VKSuccess, nil
		})

	domainOriginState := PipelineTessellationDomainOriginStateOptions{
		DomainOrigin: TessellationDomainOriginLowerLeft,
	}
	pipelines, _, err := device.CreateGraphicsPipelines(nil, nil, []core1_0.GraphicsPipelineCreateOptions{
		{
			Tessellation: &core1_0.TessellationStateOptions{
				PatchControlPoints: 1,
				NextOptions:        common.NextOptions{Next: domainOriginState},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.Equal(t, expectedPipeline.Handle(), pipelines[0].Handle())
}

func TestInputAttachmentAspectOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	expectedRenderPass := mocks.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = expectedRenderPass.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

			aspectOptions := (*khr_maintenance2_driver.VkRenderPassInputAttachmentAspectCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			aspectVal := reflect.ValueOf(aspectOptions).Elem()
			require.Equal(t, uint64(1000117001), aspectVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR
			require.True(t, aspectVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), aspectVal.FieldByName("aspectReferenceCount").Uint())

			refsPtr := (*khr_maintenance2_driver.VkInputAttachmentAspectReferenceKHR)(aspectVal.FieldByName("pAspectReferences").UnsafePointer())
			refsSlice := ([]khr_maintenance2_driver.VkInputAttachmentAspectReferenceKHR)(unsafe.Slice(refsPtr, 2))
			refsVal := reflect.ValueOf(refsSlice)
			ref := refsVal.Index(0)
			require.Equal(t, uint64(1), ref.FieldByName("subpass").Uint())
			require.Equal(t, uint64(3), ref.FieldByName("inputAttachmentIndex").Uint())
			require.Equal(t, uint64(0x00000001), ref.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

			ref = refsVal.Index(1)
			require.Equal(t, uint64(5), ref.FieldByName("subpass").Uint())
			require.Equal(t, uint64(7), ref.FieldByName("inputAttachmentIndex").Uint())
			require.Equal(t, uint64(0x00000008), ref.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT

			return core1_0.VKSuccess, nil
		})

	aspectOptions := RenderPassInputAttachmentAspectOptions{
		AspectReferences: []InputAttachmentAspectReference{
			{
				Subpass:              1,
				InputAttachmentIndex: 3,
				AspectMask:           core1_0.AspectColor,
			},
			{
				Subpass:              5,
				InputAttachmentIndex: 7,
				AspectMask:           core1_0.AspectMetadata,
			},
		},
	}
	renderPass, _, err := device.CreateRenderPass(nil, core1_0.RenderPassCreateOptions{
		NextOptions: common.NextOptions{Next: aspectOptions},
	})
	require.NoError(t, err)
	require.Equal(t, expectedRenderPass.Handle(), renderPass.Handle())
}

func TestPointClippingOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR
			properties := val.FieldByName("properties")
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)

			limits := properties.FieldByName("limits")
			*(*float32)(unsafe.Pointer(limits.FieldByName("lineWidthGranularity").UnsafeAddr())) = float32(5)

			pointClippingPtr := (*khr_maintenance2_driver.VkPhysicalDevicePointClippingPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			pointClipping := reflect.ValueOf(pointClippingPtr).Elem()

			require.Equal(t, uint64(1000117000), pointClipping.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR
			require.True(t, pointClipping.FieldByName("pNext").IsNil())

			behavior := (*khr_maintenance2_driver.VkPointClippingBehaviorKHR)(unsafe.Pointer(pointClipping.FieldByName("pointClippingBehavior").UnsafeAddr()))
			*behavior = khr_maintenance2_driver.VkPointClippingBehaviorKHR(1) // VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR
		})

	pointClipping := &PhysicalDevicePointClippingOutData{}
	properties := &khr_get_physical_device_properties2.DevicePropertiesOutData{
		NextOutData: common.NextOutData{Next: pointClipping},
	}

	err := extension.PhysicalDeviceProperties2(physicalDevice, properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.Equal(t, PointClippingUserClipPlanesOnly, pointClipping.PointClippingBehavior)
}
