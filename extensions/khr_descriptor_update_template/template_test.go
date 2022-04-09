package khr_descriptor_update_template_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template"
	khr_descriptor_update_template_driver "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/driver"
	mock_descriptor_update_template "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	buffer := mocks.EasyMockBuffer(ctrl)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		infoPtr := (*driver.VkDescriptorBufferInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, buffer.Handle(), (driver.VkBuffer)(info.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("offset").Uint())
		require.Equal(t, uint64(3), info.FieldByName("_range").Uint())
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorTemplateOptions{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromBuffer(descriptorSet, core1_0.DescriptorBufferInfo{
		Buffer: buffer,
		Offset: 1,
		Range:  3,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	sampler := mocks.EasyMockSampler(ctrl)
	imageView := mocks.EasyMockImageView(ctrl)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		infoPtr := (*driver.VkDescriptorImageInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, sampler.Handle(), (driver.VkSampler)(info.FieldByName("sampler").UnsafePointer()))
		require.Equal(t, imageView.Handle(), (driver.VkImageView)(info.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(7), info.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorTemplateOptions{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromImage(descriptorSet, core1_0.DescriptorImageInfo{
		Sampler:     sampler,
		ImageView:   imageView,
		ImageLayout: core1_0.ImageLayoutTransferDstOptimal,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromObjectHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	descriptorSet := mocks.EasyMockDescriptorSet(ctrl)
	bufferView := mocks.EasyMockBufferView(ctrl)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		descriptorSet driver.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		info := (driver.VkBufferView)(pData)
		require.Equal(t, bufferView.Handle(), info)
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorTemplateOptions{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromObjectHandle(descriptorSet, driver.VulkanHandle(bufferView.Handle()))
}
