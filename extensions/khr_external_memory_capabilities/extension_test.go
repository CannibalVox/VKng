package khr_external_memory_capabilities_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities"
	khr_external_memory_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities/driver"
	mock_external_memory_capabilities "github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities/mocks"
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

func TestVulkanExtension_ExternalBufferProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_external_memory_capabilities.NewMockDriver(ctrl)
	extension := khr_external_memory_capabilities.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceExternalBufferPropertiesKHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pExternalBufferInfo *khr_external_memory_capabilities_driver.VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *khr_external_memory_capabilities_driver.VkExternalBufferPropertiesKHR) {
		val := reflect.ValueOf(pExternalBufferInfo).Elem()

		require.Equal(t, uint64(1000071002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("flags").Uint())               // VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT
		require.Equal(t, uint64(8), val.FieldByName("usage").Uint())               // VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
		require.Equal(t, uint64(0x00000010), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR

		val = reflect.ValueOf(pExternalBufferProperties).Elem()
		require.Equal(t, uint64(1000071003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("externalMemoryFeatures").UnsafeAddr())) = uint32(2)
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(0x40)
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(2)
	})

	var outData khr_external_memory_capabilities.ExternalBufferOutData
	err := extension.ExternalBufferProperties(
		physicalDevice,
		khr_external_memory_capabilities.ExternalBufferOptions{
			Flags:      core1_0.BufferCreateSparseResidency,
			Usage:      core1_0.BufferUsageStorageTexelBuffer,
			HandleType: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11TextureKMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_memory_capabilities.ExternalBufferOutData{
		ExternalMemoryProperties: khr_external_memory_capabilities.ExternalMemoryProperties{
			ExternalMemoryFeatures:        khr_external_memory_capabilities.ExternalMemoryFeatureExportable,
			ExportFromImportedHandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D12Resource,
			CompatibleHandleTypes:         khr_external_memory_capabilities.ExternalMemoryHandleTypeOpaqueWin32,
		},
	}, outData)
}

func TestExternalImageFormatOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		physicalDevice driver.VkPhysicalDevice,
		pImageFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceImageFormatInfo2KHR,
		pImageFormatProperties *khr_get_physical_device_properties2_driver.VkImageFormatProperties2KHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pImageFormatInfo).Elem()

		require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
		require.Equal(t, uint64(68), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32

		next := (*khr_external_memory_capabilities_driver.VkPhysicalDeviceExternalImageFormatInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000071000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR

		val = reflect.ValueOf(pImageFormatProperties).Elem()

		require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR

		outDataNext := (*khr_external_memory_capabilities_driver.VkExternalImageFormatPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(outDataNext).Elem()

		require.Equal(t, uint64(1000071001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("externalMemoryFeatures").UnsafeAddr())) = uint32(4)        // VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(0x20)      // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	var outData khr_external_memory_capabilities.ExternalImageFormatOutData
	format := khr_get_physical_device_properties2.ImageFormatOutData{
		HaveNext: common.HaveNext{&outData},
	}
	_, err := extension.PhysicalDeviceImageFormatProperties(
		physicalDevice,
		khr_get_physical_device_properties2.ImageFormatOptions{
			Format: core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
			HaveNext: common.HaveNext{
				khr_external_memory_capabilities.ExternalImageFormatOptions{
					HandleType: khr_external_memory_capabilities.ExternalMemoryHandleTypeOpaqueFD,
				},
			},
		},
		&format,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_memory_capabilities.ExternalImageFormatOutData{
		ExternalMemoryProperties: khr_external_memory_capabilities.ExternalMemoryProperties{
			ExternalMemoryFeatures:        khr_external_memory_capabilities.ExternalMemoryFeatureImportable,
			ExportFromImportedHandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11Texture,
			CompatibleHandleTypes:         khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D12Heap,
		},
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

			next := (*khr_external_memory_capabilities_driver.VkPhysicalDeviceIDPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
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
	var outData khr_external_memory_capabilities.PhysicalDeviceIDOutData
	properties.HaveNext = common.HaveNext{&outData}

	err = extension.PhysicalDeviceProperties(
		physicalDevice,
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_memory_capabilities.PhysicalDeviceIDOutData{
		DeviceUUID:      deviceUUID,
		DriverUUID:      driverUUID,
		DeviceLUID:      0xdeadbeefdeadbeef,
		DeviceNodeMask:  7,
		DeviceLUIDValid: true,
	}, outData)
}
