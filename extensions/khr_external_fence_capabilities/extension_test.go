package khr_external_fence_capabilities_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities"
	khr_external_fence_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities/driver"
	mock_external_fence_capabilities "github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/golang/mock/gomock"
	uuid2 "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_PhysicalDeviceExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_external_fence_capabilities.NewMockDriver(ctrl)
	extension := khr_external_fence_capabilities.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceExternalFencePropertiesKHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pExternalFenceInfo *khr_external_fence_capabilities_driver.VkPhysicalDeviceExternalFenceInfoKHR,
		pExternalFenceProperties *khr_external_fence_capabilities_driver.VkExternalFencePropertiesKHR,
	) {
		val := reflect.ValueOf(pExternalFenceInfo).Elem()
		require.Equal(t, uint64(1000112000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR

		val = reflect.ValueOf(pExternalFenceProperties).Elem()
		*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalFenceFeatures").UnsafeAddr())) = uint32(1)         // VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR
	})

	var outData khr_external_fence_capabilities.ExternalFenceOutData
	err := extension.ExternalFenceProperties(
		physicalDevice,
		khr_external_fence_capabilities.ExternalFenceOptions{
			HandleType: khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32KMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_fence_capabilities.ExternalFenceOutData{
		ExportFromImportedHandleTypes: khr_external_fence_capabilities.ExternalFenceHandleTypeSyncFD,
		CompatibleHandleTypes:         khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32KMT,
		ExternalFenceFeatures:         khr_external_fence_capabilities.ExternalFenceFeatureExportable,
	}, outData)
}

func TestPhysicalDeviceIDOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	deviceUUID, err := uuid2.NewRandom()
	require.NoError(t, err)

	driverUUID, err := uuid2.NewRandom()
	require.NoError(t, err)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice driver.VkPhysicalDevice,
			pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
		) {
			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_external_fence_capabilities_driver.VkPhysicalDeviceIDPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000071004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			for i := 0; i < len(deviceUUID); i++ {
				*(*byte)(unsafe.Pointer(val.FieldByName("deviceUUID").Index(i).UnsafeAddr())) = deviceUUID[i]
				*(*byte)(unsafe.Pointer(val.FieldByName("driverUUID").Index(i).UnsafeAddr())) = driverUUID[i]
			}

			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(0).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(1).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(2).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(3).UnsafeAddr())) = byte(0xde)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(4).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(5).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(6).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(7).UnsafeAddr())) = byte(0xde)

			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceNodeMask").UnsafeAddr())) = uint32(7)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("deviceLUIDValid").UnsafeAddr())) = driver.VkBool32(1)
		})

	var properties khr_get_physical_device_properties2.DevicePropertiesOutData
	var outData khr_external_fence_capabilities.PhysicalDeviceIDOutData
	properties.NextOutData = common.NextOutData{&outData}

	err = extension.PhysicalDeviceProperties2(
		physicalDevice,
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_fence_capabilities.PhysicalDeviceIDOutData{
		DeviceUUID:      deviceUUID,
		DriverUUID:      driverUUID,
		DeviceLUID:      0xdeadbeefdeadbeef,
		DeviceNodeMask:  7,
		DeviceLUIDValid: true,
	}, outData)
}
