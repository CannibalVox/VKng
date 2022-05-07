package khr_variable_pointers_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_variable_pointers"
	khr_variable_pointers_driver "github.com/CannibalVox/VKng/extensions/khr_variable_pointers/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVariablePointersFeatureOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
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

			featuresPtr := (*khr_variable_pointers_driver.VkPhysicalDeviceVariablePointersFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("variablePointers").Uint())
			require.Equal(t, uint64(0), val.FieldByName("variablePointersStorageBuffer").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueOptions{
			{
				CreatedQueuePriorities: []float32{0},
			},
		},

		HaveNext: common.HaveNext{Next: khr_variable_pointers.VariablePointersFeatureOptions{
			VariablePointers:              true,
			VariablePointersStorageBuffer: false,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestVariablePointersFeatureOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	var pointersOutData khr_variable_pointers.VariablePointersFeatureOutData

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

			outData := (*khr_variable_pointers_driver.VkPhysicalDeviceVariablePointersFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointers").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointersStorageBuffer").UnsafeAddr())) = driver.VkBool32(1)
		})

	err := extension.PhysicalDeviceFeatures(physicalDevice, &khr_get_physical_device_properties2.DeviceFeaturesOutData{
		HaveNext: common.HaveNext{Next: &pointersOutData},
	})
	require.NoError(t, err)
	require.True(t, pointersOutData.VariablePointersStorageBuffer)
	require.False(t, pointersOutData.VariablePointers)
}
