package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestCommandPoolCreateBasic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	expectedPoolHandle := mocks.NewFakeCommandPoolHandle()

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	mockDriver.EXPECT().VkCreateCommandPool(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandPoolCreateInfo, allocator *driver.VkAllocationCallbacks, commandPool *driver.VkCommandPool) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(39), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
			require.Equal(t, uint64(1), val.FieldByName("queueFamilyIndex").Uint())

			*commandPool = expectedPoolHandle

			return core1_0.VKSuccess, nil
		})

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	graphicsFamily := 1
	pool, res, err := loader.CreateCommandPool(device, nil, &core1_0.CommandPoolOptions{
		Flags:               core1_0.CommandPoolCreateResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, pool)
	require.Same(t, expectedPoolHandle, pool.Handle())
}

func TestCommandPoolNullQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	pool, res, err := loader.CreateCommandPool(mockDevice, nil, &core1_0.CommandPoolOptions{
		Flags: core1_0.CommandPoolCreateResetBuffer,
	})
	require.Error(t, err)
	require.Equal(t, core1_0.VKErrorUnknown, res)
	require.Nil(t, pool)
}

func TestCommandBufferSingleAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	commandPool := internal_mocks.EasyDummyCommandPool(t, loader, mockDevice)

	bufferHandle := mocks.NewFakeCommandBufferHandle()

	mockDriver.EXPECT().VkAllocateCommandBuffers(mocks.Exactly(mockDevice.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandBufferAllocateInfo, commandBuffers *driver.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_PRIMARY
			require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())

			require.Same(t, commandPool.Handle(), driver.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBuffers, 1))
			bufferSlice[0] = bufferHandle

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := loader.AllocateCommandBuffers(&core1_0.CommandBufferOptions{
		CommandPool: commandPool,
		Level:       core1_0.LevelPrimary,
		BufferCount: 1,
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Len(t, buffers, 1)
	require.Same(t, buffers[0].Handle(), bufferHandle)

	mockDriver.EXPECT().VkFreeCommandBuffers(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(commandPool.Handle()), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, commandPool driver.VkCommandPool, bufferCount driver.Uint32, buffers *driver.VkCommandBuffer) error {
			slice := ([]driver.VkCommandBuffer)(unsafe.Slice(buffers, 1))
			require.Same(t, bufferHandle, slice[0])

			return nil
		})

	loader.FreeCommandBuffers(buffers)
}

func TestCommandBufferMultiAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	commandPool := internal_mocks.EasyDummyCommandPool(t, loader, mockDevice)

	bufferHandles := []driver.VkCommandBuffer{
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
	}

	mockDriver.EXPECT().VkAllocateCommandBuffers(mocks.Exactly(mockDevice.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandBufferAllocateInfo, commandBuffers *driver.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_SECONDARY
			require.Equal(t, uint64(3), val.FieldByName("commandBufferCount").Uint())

			require.Same(t, commandPool.Handle(), driver.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBuffers, 3))
			bufferSlice[0] = bufferHandles[0]
			bufferSlice[1] = bufferHandles[1]
			bufferSlice[2] = bufferHandles[2]

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := loader.AllocateCommandBuffers(&core1_0.CommandBufferOptions{
		CommandPool: commandPool,
		Level:       core1_0.LevelSecondary,
		BufferCount: 3,
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Len(t, buffers, 3)

	require.Same(t, bufferHandles[0], buffers[0].Handle())
	require.Same(t, bufferHandles[1], buffers[1].Handle())
	require.Same(t, bufferHandles[2], buffers[2].Handle())

	mockDriver.EXPECT().VkFreeCommandBuffers(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(commandPool.Handle()), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, commandPool driver.VkCommandPool, bufferCount driver.Uint32, buffers *driver.VkCommandBuffer) {
			slice := ([]driver.VkCommandBuffer)(unsafe.Slice(buffers, 3))
			require.Same(t, bufferHandles[0], slice[0])
			require.Same(t, bufferHandles[1], slice[1])
			require.Same(t, bufferHandles[2], slice[2])
		})

	loader.FreeCommandBuffers(buffers)
}

func TestVulkanCommandPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	commandPool := internal_mocks.EasyDummyCommandPool(t, loader, mockDevice)

	mockDriver.EXPECT().VkResetCommandPool(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(commandPool.Handle()),
		driver.VkCommandPoolResetFlags(1), // VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT
	).Return(core1_0.VKSuccess, nil)

	_, err = commandPool.Reset(core1_0.CommandPoolResetReleaseResources)
	require.NoError(t, err)
}
