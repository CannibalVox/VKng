package khr_16bit_storage_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_16bit_storage"
	khr_16bit_storage_driver "github.com/CannibalVox/VKng/extensions/khr_16bit_storage/driver"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestDevice16BitStorageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)

	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	expectedDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = expectedDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			storageFeatures := (*khr_16bit_storage_driver.VkPhysicalDevice16BitStorageFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			storageVal := reflect.ValueOf(storageFeatures).Elem()

			require.Equal(t, uint64(1000083000), storageVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
			require.True(t, storageVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), storageVal.FieldByName("storageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("uniformAndStorageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(0), storageVal.FieldByName("storagePushConstant16").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("storageInputOutput16").Uint())

			return core1_0.VKSuccess, nil
		})

	storage := khr_16bit_storage.Device16BitStorageOptions{
		StorageInputOutput16:               true,
		UniformAndStorageBuffer16BitAccess: true,
		StoragePushConstant16:              false,
		StorageBuffer16BitAccess:           false,
	}
	device, _, err := loader.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueOptions{
			{
				CreatedQueuePriorities: []float32{0},
			},
		},

		HaveNext: common.HaveNext{Next: storage},
	})

	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Handle())
}

func TestDevice16BitStorageOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			featureVal := val.FieldByName("features")
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = driver.VkBool32(1)

			outDataPtr := (*khr_16bit_storage_driver.VkPhysicalDevice16BitStorageFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			outDataVal := reflect.ValueOf(outDataPtr).Elem()
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageBuffer16BitAccess").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("uniformAndStorageBuffer16BitAccess").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storagePushConstant16").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageInputOutput16").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := &khr_16bit_storage.Device16BitStorageOutData{}
	features := &khr_get_physical_device_properties2.DeviceFeaturesOutData{
		HaveNext: common.HaveNext{Next: outData},
	}

	err := extension.PhysicalDeviceFeatures(physicalDevice, features)
	require.NoError(t, err)

	require.True(t, outData.StoragePushConstant16)
	require.False(t, outData.UniformAndStorageBuffer16BitAccess)
	require.True(t, outData.StorageInputOutput16)
	require.False(t, outData.StorageBuffer16BitAccess)

	require.True(t, features.Features.FillModeNonSolid)
}
