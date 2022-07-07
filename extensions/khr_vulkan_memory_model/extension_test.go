package khr_vulkan_memory_model_test

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
	"github.com/CannibalVox/VKng/extensions/khr_vulkan_memory_model"
	khr_vulkan_memory_model_driver "github.com/CannibalVox/VKng/extensions/khr_vulkan_memory_model/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceVulkanMemoryModelFeaturesOptions(t *testing.T) {
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
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pCreateInfo *driver.VkDeviceCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pDevice *driver.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*khr_vulkan_memory_model_driver.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModel").Uint())
			require.Equal(t, uint64(0), val.FieldByName("vulkanMemoryModelDeviceScope").Uint())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures{
			VulkanMemoryModel:                             true,
			VulkanMemoryModelDeviceScope:                  false,
			VulkanMemoryModelAvailabilityVisibilityChains: true,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceVulkanMemoryModelFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice driver.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			outData := (*khr_vulkan_memory_model_driver.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModel").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelDeviceScope").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").UnsafeAddr())) = driver.VkBool32(1)
		})

	var outData khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures
	err := extension.PhysicalDeviceFeatures2(physicalDevice, &khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &outData},
	})
	require.NoError(t, err)
	require.Equal(t, khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures{
		VulkanMemoryModel:                             true,
		VulkanMemoryModelDeviceScope:                  false,
		VulkanMemoryModelAvailabilityVisibilityChains: true,
	}, outData)
}
