package mocks

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func EasyMockDevice(ctrl *gomock.Controller, driver core.Driver) *MockDevice {
	device := NewMockDevice(ctrl)
	device.EXPECT().Handle().Return(NewFakeDeviceHandle()).AnyTimes()
	device.EXPECT().Driver().Return(driver).AnyTimes()

	return device
}

func EasyMockPhysicalDevice(ctrl *gomock.Controller, driver core.Driver) *MockPhysicalDevice {
	physicalDevice := NewMockPhysicalDevice(ctrl)
	physicalDevice.EXPECT().Handle().Return(NewFakePhysicalDeviceHandle()).AnyTimes()
	physicalDevice.EXPECT().Driver().Return(driver).AnyTimes()

	return physicalDevice
}

func EasyMockSampler(ctrl *gomock.Controller, driver core.Driver) *MockSampler {
	sampler := NewMockSampler(ctrl)
	sampler.EXPECT().Handle().Return(NewFakeSampler()).AnyTimes()

	return sampler
}

func EasyDummyCommandPool(t *testing.T, loader core.Loader1_0, device core.Device) core.CommandPool {
	handle := NewFakeCommandPoolHandle()
	driver := device.Driver().(*MockDriver)
	driver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkCommandPoolCreateInfo, allocator *core.VkAllocationCallbacks, commandPool *core.VkCommandPool) (core.VkResult, error) {
			*commandPool = handle

			return core.VKSuccess, nil
		})

	graphicsFamily := 0
	pool, res, err := loader.CreateCommandPool(device, &core.CommandPoolOptions{
		Flags:               core.CommandPoolResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, core.VKSuccess, res)

	return pool
}

func EasyDummyCommandBuffer(t *testing.T, device core.Device, commandPool core.CommandPool) core.CommandBuffer {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkAllocateCommandBuffers(gomock.Any(), gomock.Any(), gomock.Any()).Do(
		func(device core.VkDevice, pAllocateInfo *core.VkCommandBufferAllocateInfo, pCommandBuffers *core.VkCommandBuffer) {
			*pCommandBuffers = NewFakeCommandBufferHandle()
		})

	buffers, _, err := commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		BufferCount: 1,
	})
	require.NoError(t, err)

	return buffers[0]
}

func EasyDummyDescriptorPool(t *testing.T, loader core.Loader1_0, device core.Device) core.DescriptorPool {
	mockDriver := device.Driver().(*MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorPoolCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorPool *core.VkDescriptorPool) (core.VkResult, error) {
			*pDescriptorPool = NewFakeDescriptorPool()
			return core.VKSuccess, nil
		})

	pool, _, err := loader.CreateDescriptorPool(device, &core.DescriptorPoolOptions{})
	require.NoError(t, err)

	return pool
}

func EasyDummyDescriptorSetLayout(t *testing.T, loader core.Loader1_0, device core.Device) core.DescriptorSetLayout {
	mockDriver := device.Driver().(*MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorSetLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorSetLayout *core.VkDescriptorSetLayout) (core.VkResult, error) {
			*pDescriptorSetLayout = NewFakeDescriptorSetLayout()
			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(device, &core.DescriptorSetLayoutOptions{})
	require.NoError(t, err)
	return layout
}

func EasyDummyDevice(t *testing.T, ctrl *gomock.Controller, loader core.Loader1_0) core.Device {
	mockDriver := loader.Driver().(*MockDriver)

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pCreateInfo *core.VkDeviceCreateInfo, pAllocator *core.VkAllocationCallbacks, pDevice *core.VkDevice) (core.VkResult, error) {
			*pDevice = NewFakeDeviceHandle()
			return core.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(EasyMockPhysicalDevice(ctrl, mockDriver), &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueuePriorities: []float32{1},
			},
		},
	})
	require.NoError(t, err)

	return device
}

func EasyDummyFence(t *testing.T, loader core.Loader1_0, device core.Device) core.Fence {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateFence(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFenceCreateInfo, pAllocator *core.VkAllocationCallbacks, pFence *core.VkFence) (core.VkResult, error) {
			*pFence = NewFakeFence()
			return core.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, &core.FenceOptions{})
	require.NoError(t, err)

	return fence
}

func EasyDummyFramebuffer(t *testing.T, loader core.Loader1_0, device core.Device) core.Framebuffer {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateFramebuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFramebufferCreateInfo, pAllocator *core.VkAllocationCallbacks, pFramebuffer *core.VkFramebuffer) (core.VkResult, error) {
			*pFramebuffer = NewFakeFramebufferHandle()
			return core.VKSuccess, nil
		})

	framebuffer, _, err := loader.CreateFrameBuffer(device, &core.FramebufferOptions{})
	require.NoError(t, err)

	return framebuffer
}

func EasyDummyGraphicsPipeline(t *testing.T, loader core.Loader1_0, device core.Device) core.Pipeline {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateGraphicsPipelines(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, cache core.VkPipelineCache, createInfoCount core.Uint32, pCreateInfos *core.VkGraphicsPipelineCreateInfo, pAllocator *core.VkAllocationCallbacks, pPipelines *core.VkPipeline) (core.VkResult, error) {
			*pPipelines = NewFakePipeline()
			return core.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, []*core.GraphicsPipelineOptions{{}})
	require.NoError(t, err)

	return pipelines[0]
}

func EasyDummyQueue(t *testing.T, device core.Device) core.Queue {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkGetDeviceQueue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, queueFamilyIndex, queueIndex core.Uint32, pQueue *core.VkQueue) error {
			*pQueue = NewFakeQueue()
			return nil
		})

	queue, err := device.GetQueue(0, 0)
	require.NoError(t, err)

	return queue
}

func EasyDummyRenderPass(t *testing.T, loader core.Loader1_0, device core.Device) core.RenderPass {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateRenderPass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkRenderPassCreateInfo, pAllocator *core.VkAllocationCallbacks, pRenderPass *core.VkRenderPass) (core.VkResult, error) {
			*pRenderPass = NewFakeRenderPassHandle()
			return core.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, &core.RenderPassOptions{})
	require.NoError(t, err)

	return renderPass
}

func EasyDummySemaphore(t *testing.T, loader core.Loader1_0, device core.Device) core.Semaphore {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateSemaphore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkSemaphoreCreateInfo, pAllocator *core.VkAllocationCallbacks, pSemaphore *core.VkSemaphore) (core.VkResult, error) {
			*pSemaphore = NewFakeSemaphore()
			return core.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, &core.SemaphoreOptions{})
	require.NoError(t, err)

	return semaphore
}