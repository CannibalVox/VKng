package khr_external_semaphore_capabilities_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities"
	khr_external_semaphore_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities/driver"
	mock_external_semaphore_capabilities "github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities/mocks"
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

			next := (*khr_external_semaphore_capabilities_driver.VkPhysicalDeviceIDPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
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

	var properties khr_get_physical_device_properties2.PhysicalDeviceProperties2
	var outData khr_external_semaphore_capabilities.PhysicalDeviceIDProperties
	properties.NextOutData = common.NextOutData{&outData}

	err = extension.PhysicalDeviceProperties2(
		physicalDevice,
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_semaphore_capabilities.PhysicalDeviceIDProperties{
		DeviceUUID:      deviceUUID,
		DriverUUID:      driverUUID,
		DeviceLUID:      0xdeadbeefdeadbeef,
		DeviceNodeMask:  7,
		DeviceLUIDValid: true,
	}, outData)
}

func TestVulkanExtension_ExternalSemaphoreProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_external_semaphore_capabilities.NewMockDriver(ctrl)
	extension := khr_external_semaphore_capabilities.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pExternalSemaphoreInfo *khr_external_semaphore_capabilities_driver.VkPhysicalDeviceExternalSemaphoreInfoKHR,
			pExternalSemaphoreProperties *khr_external_semaphore_capabilities_driver.VkExternalSemaphorePropertiesKHR,
		) {
			val := reflect.ValueOf(pExternalSemaphoreInfo).Elem()

			require.Equal(t, uint64(1000076000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR

			val = reflect.ValueOf(pExternalSemaphoreProperties).Elem()
			require.Equal(t, uint64(1000076001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(1) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("externalSemaphoreFeatures").UnsafeAddr())) = uint32(2)     // VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR
		})

	var outData khr_external_semaphore_capabilities.ExternalSemaphoreProperties
	err := extension.PhysicalDeviceExternalSemaphoreProperties(
		physicalDevice,
		khr_external_semaphore_capabilities.PhysicalDeviceExternalSemaphoreInfo{
			HandleType: khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeSyncFD,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, khr_external_semaphore_capabilities.ExternalSemaphoreProperties{
		ExportFromImportedHandleTypes: khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeOpaqueFD,
		CompatibleHandleTypes:         khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeOpaqueWin32KMT,
		ExternalSemaphoreFeatures:     khr_external_semaphore_capabilities.ExternalSemaphoreFeatureImportable,
	}, outData)
}
