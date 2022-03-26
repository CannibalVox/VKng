package khr_bind_memory2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/driver"
	mock_bind_memory2 "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_BindBufferMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)

	memory1 := mocks.EasyMockDeviceMemory(ctrl)
	memory2 := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkBindBufferMemory2KHR(device.Handle(), driver.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer1.Handle(), (driver.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer2.Handle(), (driver.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := extension.BindBufferMemory(device, []khr_bind_memory2.BindBufferMemoryOptions{
		{
			Buffer:       buffer1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Buffer:       buffer2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	image1 := mocks.EasyMockImage(ctrl)
	image2 := mocks.EasyMockImage(ctrl)

	memory1 := mocks.EasyMockDeviceMemory(ctrl)
	memory2 := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkBindImageMemory2KHR(device.Handle(), driver.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image1.Handle(), (driver.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image2.Handle(), (driver.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := extension.BindImageMemory(device, []khr_bind_memory2.BindImageMemoryOptions{
		{
			Image:        image1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Image:        image2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}
