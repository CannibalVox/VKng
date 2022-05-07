package khr_external_memory_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_memory"
	khr_external_memory_driver "github.com/CannibalVox/VKng/extensions/khr_external_memory/driver"
	"github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestExternalMemoryAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pAllocateInfo *driver.VkMemoryAllocateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pMemory *driver.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*khr_external_memory_driver.VkExportMemoryAllocateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.MemoryAllocateOptions{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		HaveNext: common.HaveNext{
			khr_external_memory.ExternalMemoryAllocateOptions{
				HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11TextureKMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestExternalMemoryImageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockImage := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkImageCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pImage *driver.VkImage,
		) (common.VkResult, error) {
			*pImage = mockImage.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
			require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

			next := (*khr_external_memory_driver.VkExternalMemoryImageCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x20), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(
		device,
		nil,
		core1_0.ImageCreateOptions{
			MipLevels:   1,
			ArrayLayers: 3,

			HaveNext: common.HaveNext{
				khr_external_memory.ExternalMemoryImageOptions{
					HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D12Heap,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}

func TestExternalMemoryBufferOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockBuffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkBufferCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pImage *driver.VkBuffer,
		) (common.VkResult, error) {
			*pImage = mockBuffer.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("size").Uint())
			require.Equal(t, uint64(8), val.FieldByName("usage").Uint())

			next := (*khr_external_memory_driver.VkExternalMemoryImageCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	buffer, _, err := loader.CreateBuffer(
		device,
		nil,
		core1_0.BufferCreateOptions{
			BufferSize: 1,
			Usage:      core1_0.BufferUsageStorageTexelBuffer,

			HaveNext: common.HaveNext{
				khr_external_memory.ExternalMemoryBufferOptions{
					HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11Texture,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockBuffer.Handle(), buffer.Handle())
}
