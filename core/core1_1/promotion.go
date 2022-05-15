package core1_1

import "github.com/CannibalVox/VKng/core/core1_0"
import _ "unsafe"

//go:linkname PromoteBuffer github.com/CannibalVox/VKng/core/internal/core1_1.PromoteBuffer
func PromoteBuffer(buffer core1_0.Buffer) Buffer

//go:linkname PromoteBufferView github.com/CannibalVox/VKng/core/internal/core1_1.PromoteBufferView
func PromoteBufferView(bufferView core1_0.BufferView) BufferView

//go:linkname PromoteCommandBuffer github.com/CannibalVox/VKng/core/internal/core1_1.PromoteCommandBuffer
func PromoteCommandBuffer(commandBuffer core1_0.CommandBuffer) CommandBuffer

//go:linkname PromoteCommandBufferSlice github.com/CannibalVox/VKng/core/internal/core1_1.PromoteCommandBufferSlice
func PromoteCommandBufferSlice(commandBuffers []core1_0.CommandBuffer) []CommandBuffer

//go:linkname PromoteCommandPool github.com/CannibalVox/VKng/core/internal/core1_1.PromoteCommandPool
func PromoteCommandPool(commandPool core1_0.CommandPool) CommandPool

//go:linkname PromoteDescriptorPool github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDescriptorPool
func PromoteDescriptorPool(descriptorPool core1_0.DescriptorPool) DescriptorPool

//go:linkname PromoteDescriptorSet github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDescriptorSet
func PromoteDescriptorSet(descriptorSet core1_0.DescriptorSet) DescriptorSet

//go:linkname PromoteDescriptorSetSlice github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDescriptorSetSlice
func PromoteDescriptorSetSlice(descriptorSets []core1_0.DescriptorSet) []DescriptorSet

//go:linkname PromoteDescriptorSetLayout github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDescriptorSetLayout
func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) DescriptorSetLayout

//go:linkname PromoteDevice github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDevice
func PromoteDevice(device core1_0.Device) Device

//go:linkname PromoteDeviceMemory github.com/CannibalVox/VKng/core/internal/core1_1.PromoteDeviceMemory
func PromoteDeviceMemory(deviceMemory core1_0.DeviceMemory) DeviceMemory

//go:linkname PromoteEvent github.com/CannibalVox/VKng/core/internal/core1_1.PromoteEvent
func PromoteEvent(event core1_0.Event) Event

//go:linkname PromoteFence github.com/CannibalVox/VKng/core/internal/core1_1.PromoteFence
func PromoteFence(fence core1_0.Fence) Fence

//go:linkname PromoteFramebuffer github.com/CannibalVox/VKng/core/internal/core1_1.PromoteFramebuffer
func PromoteFramebuffer(framebuffer core1_0.Framebuffer) Framebuffer

//go:linkname PromoteImage github.com/CannibalVox/VKng/core/internal/core1_1.PromoteImage
func PromoteImage(image core1_0.Image) Image

//go:linkname PromoteImageView github.com/CannibalVox/VKng/core/internal/core1_1.PromoteImageView
func PromoteImageView(imageView core1_0.ImageView) ImageView

//go:linkname PromoteInstance github.com/CannibalVox/VKng/core/internal/core1_1.PromoteInstance
func PromoteInstance(instance core1_0.Instance) Instance

//go:linkname PromotePipeline github.com/CannibalVox/VKng/core/internal/core1_1.PromotePipeline
func PromotePipeline(pipeline core1_0.Pipeline) Pipeline

//go:linkname PromotePipelineSlice github.com/CannibalVox/VKng/core/internal/core1_1.PromotePipelineSlice
func PromotePipelineSlice(pipelines []core1_0.Pipeline) []Pipeline

//go:linkname PromotePipelineCache github.com/CannibalVox/VKng/core/internal/core1_1.PromotePipelineCache
func PromotePipelineCache(pipelineCache core1_0.PipelineCache) PipelineCache

//go:linkname PromotePipelineLayout github.com/CannibalVox/VKng/core/internal/core1_1.PromotePipelineLayout
func PromotePipelineLayout(pipelineLayout core1_0.PipelineLayout) PipelineLayout

//go:linkname PromoteQueryPool github.com/CannibalVox/VKng/core/internal/core1_1.PromoteQueryPool
func PromoteQueryPool(queryPool core1_0.QueryPool) QueryPool

//go:linkname PromoteQueue github.com/CannibalVox/VKng/core/internal/core1_1.PromoteQueue
func PromoteQueue(queue core1_0.Queue) Queue

//go:linkname PromoteRenderPass github.com/CannibalVox/VKng/core/internal/core1_1.PromoteRenderPass
func PromoteRenderPass(renderPass core1_0.RenderPass) RenderPass

//go:linkname PromoteSampler github.com/CannibalVox/VKng/core/internal/core1_1.PromoteSampler
func PromoteSampler(sampler core1_0.Sampler) Sampler

//go:linkname PromoteSemaphore github.com/CannibalVox/VKng/core/internal/core1_1.PromoteSemaphore
func PromoteSemaphore(semaphore core1_0.Semaphore) Semaphore

//go:linkname PromoteShaderModule github.com/CannibalVox/VKng/core/internal/core1_1.PromoteShaderModule
func PromoteShaderModule(shaderModule core1_0.ShaderModule) ShaderModule

//go:linkname PromotePhysicalDevice github.com/CannibalVox/VKng/core/internal/core1_1.PromotePhysicalDevice
func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) PhysicalDevice

//go:linkname PromoteInstanceScopedPhysicalDevice github.com/CannibalVox/VKng/core/internal/core1_1.PromoteInstanceScopedPhysicalDevice
func PromoteInstanceScopedPhysicalDevice(physicalDevice core1_0.PhysicalDevice) InstanceScopedPhysicalDevice
