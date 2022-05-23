package internal1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestSubmitToQueue_SignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	queue := internal_mocks.EasyDummyQueue(mockDriver, mockDevice)

	fence := mocks.EasyMockFence(ctrl)
	buffer := mocks.EasyMockCommandBuffer(ctrl)

	waitSemaphore1 := mocks.EasyMockSemaphore(ctrl)
	waitSemaphore2 := mocks.EasyMockSemaphore(ctrl)

	signalSemaphore1 := mocks.EasyMockSemaphore(ctrl)
	signalSemaphore2 := mocks.EasyMockSemaphore(ctrl)
	signalSemaphore3 := mocks.EasyMockSemaphore(ctrl)

	mockDriver.EXPECT().VkQueueSubmit(queue.Handle(), driver.Uint32(1), gomock.Not(nil), fence.Handle()).DoAndReturn(
		func(queue driver.VkQueue, submitCount driver.Uint32, pSubmits *driver.VkSubmitInfo, fence driver.VkFence) (common.VkResult, error) {
			submitSlices := ([]driver.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

			for _, submit := range submitSlices {
				v := reflect.ValueOf(submit)
				require.Equal(t, uint64(4), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
				require.True(t, v.FieldByName("pNext").IsNil())
				require.Equal(t, uint64(2), v.FieldByName("waitSemaphoreCount").Uint())
				require.Equal(t, uint64(1), v.FieldByName("commandBufferCount").Uint())
				require.Equal(t, uint64(3), v.FieldByName("signalSemaphoreCount").Uint())

				waitSemaphorePtr := unsafe.Pointer(v.FieldByName("pWaitSemaphores").Elem().UnsafeAddr())
				waitSemaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice((*driver.VkSemaphore)(waitSemaphorePtr), 2))

				require.Equal(t, waitSemaphore1.Handle(), waitSemaphoreSlice[0])
				require.Equal(t, waitSemaphore2.Handle(), waitSemaphoreSlice[1])

				waitDstStageMaskPtr := unsafe.Pointer(v.FieldByName("pWaitDstStageMask").Elem().UnsafeAddr())
				waitDstStageMaskSlice := ([]driver.VkPipelineStageFlags)(unsafe.Slice((*driver.VkPipelineStageFlags)(waitDstStageMaskPtr), 2))

				require.ElementsMatch(t, []driver.VkPipelineStageFlags{8, 128}, waitDstStageMaskSlice)

				commandBufferPtr := unsafe.Pointer(v.FieldByName("pCommandBuffers").Elem().UnsafeAddr())
				commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice((*driver.VkCommandBuffer)(commandBufferPtr), 1))

				require.Equal(t, buffer.Handle(), commandBufferSlice[0])

				signalSemaphorePtr := unsafe.Pointer(v.FieldByName("pSignalSemaphores").Elem().UnsafeAddr())
				signalSemaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice((*driver.VkSemaphore)(signalSemaphorePtr), 3))

				require.Equal(t, signalSemaphore1.Handle(), signalSemaphoreSlice[0])
				require.Equal(t, signalSemaphore2.Handle(), signalSemaphoreSlice[1])
				require.Equal(t, signalSemaphore3.Handle(), signalSemaphoreSlice[2])
			}

			return core1_0.VKSuccess, nil
		})

	_, err := queue.SubmitToQueue(fence, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStages:    []common.PipelineStages{core1_0.PipelineStageVertexShader, core1_0.PipelineStageFragmentShader},
			SignalSemaphores: []core1_0.Semaphore{signalSemaphore1, signalSemaphore2, signalSemaphore3},
		},
	})
	require.NoError(t, err)
}

func TestSubmitToQueue_NoSignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	queue := internal_mocks.EasyDummyQueue(mockDriver, mockDevice)

	buffer := mocks.EasyMockCommandBuffer(ctrl)

	mockDriver.EXPECT().VkQueueSubmit(queue.Handle(), driver.Uint32(1), gomock.Not(nil), driver.VkFence(driver.NullHandle)).DoAndReturn(
		func(queue driver.VkQueue, submitCount driver.Uint32, pSubmits *driver.VkSubmitInfo, fence driver.VkFence) (common.VkResult, error) {
			submitSlices := ([]driver.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

			for _, submit := range submitSlices {
				v := reflect.ValueOf(submit)
				require.Equal(t, uint64(4), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
				require.True(t, v.FieldByName("pNext").IsNil())
				require.Equal(t, uint64(0), v.FieldByName("waitSemaphoreCount").Uint())
				require.Equal(t, uint64(1), v.FieldByName("commandBufferCount").Uint())
				require.Equal(t, uint64(0), v.FieldByName("signalSemaphoreCount").Uint())

				require.True(t, v.FieldByName("pWaitSemaphores").IsNil())
				require.True(t, v.FieldByName("pWaitDstStageMask").IsNil())
				require.True(t, v.FieldByName("pSignalSemaphores").IsNil())

				commandBufferPtr := unsafe.Pointer(v.FieldByName("pCommandBuffers").Elem().UnsafeAddr())
				commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice((*driver.VkCommandBuffer)(commandBufferPtr), 1))

				require.Equal(t, buffer.Handle(), commandBufferSlice[0])
			}

			return core1_0.VKSuccess, nil
		})

	_, err := queue.SubmitToQueue(nil, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{},
			WaitDstStages:    []common.PipelineStages{},
			SignalSemaphores: []core1_0.Semaphore{},
		},
	})
	require.NoError(t, err)
}

func TestSubmitToQueue_MismatchWaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	queue := internal_mocks.EasyDummyQueue(mockDriver, mockDevice)

	buffer := mocks.EasyMockCommandBuffer(ctrl)

	waitSemaphore1 := mocks.EasyMockSemaphore(ctrl)
	waitSemaphore2 := mocks.EasyMockSemaphore(ctrl)

	_, err := queue.SubmitToQueue(nil, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStages:    []common.PipelineStages{core1_0.PipelineStageFragmentShader},
			SignalSemaphores: []core1_0.Semaphore{},
		},
	})
	require.EqualError(t, err, "attempted to submit with 2 wait semaphores but 1 dst stages- these should match")
}
