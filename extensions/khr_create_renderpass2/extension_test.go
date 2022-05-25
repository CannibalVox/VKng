package khr_create_renderpass2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_create_renderpass2"
	khr_create_renderpass2_driver "github.com/CannibalVox/VKng/extensions/khr_create_renderpass2/driver"
	mock_create_renderpass2 "github.com/CannibalVox/VKng/extensions/khr_create_renderpass2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_CmdBeginRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	framebuffer := mocks.EasyMockFramebuffer(ctrl)

	extDriver.EXPECT().VkCmdBeginRenderPass2KHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		pSubpassBeginInfo *khr_create_renderpass2_driver.VkSubpassBeginInfoKHR) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, renderPass.Handle(), driver.VkRenderPass(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), driver.VkFramebuffer(val.FieldByName("framebuffer").UnsafePointer()))
		require.Equal(t, int64(1), val.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), val.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), val.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), val.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
		require.Equal(t, uint64(1), val.FieldByName("clearValueCount").Uint())

		values := (*driver.Float)(unsafe.Pointer(val.FieldByName("pClearValues").Elem().UnsafeAddr()))
		valueSlice := ([]driver.Float)(unsafe.Slice(values, 4))
		val = reflect.ValueOf(valueSlice)
		require.InDelta(t, 1.0, val.Index(0).Float(), 0.0001)
		require.InDelta(t, 3.0, val.Index(1).Float(), 0.0001)
		require.InDelta(t, 5.0, val.Index(2).Float(), 0.0001)
		require.InDelta(t, 7.0, val.Index(3).Float(), 0.0001)

		val = reflect.ValueOf(pSubpassBeginInfo).Elem()
		require.Equal(t, uint64(1000109005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("contents").Uint()) // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	})

	err := extension.CmdBeginRenderPass2(
		commandBuffer,
		core1_0.RenderPassBeginOptions{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			RenderArea:  common.Rect2D{Offset: common.Offset2D{1, 3}, Extent: common.Extent2D{5, 7}},
			ClearValues: []common.ClearValue{common.ClearValueFloat{1, 3, 5, 7}},
		},
		khr_create_renderpass2.SubpassBeginOptions{
			Contents: core1_0.SubpassContentsSecondaryCommandBuffers,
		})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdEndRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	extDriver.EXPECT().VKCmdEndRenderPass2KHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pSubpassEndInfo *khr_create_renderpass2_driver.VkSubpassEndInfoKHR) {

		val := reflect.ValueOf(pSubpassEndInfo).Elem()

		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := extension.CmdEndRenderPass2(
		commandBuffer,
		khr_create_renderpass2.SubpassEndOptions{})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdNextSubpass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	extDriver.EXPECT().VkCmdNextSubpass2KHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pSubpassBeginInfo *khr_create_renderpass2_driver.VkSubpassBeginInfoKHR,
		pSubpassEndInfo *khr_create_renderpass2_driver.VkSubpassEndInfoKHR) {

		val := reflect.ValueOf(pSubpassBeginInfo).Elem()
		require.Equal(t, uint64(1000109005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("contents").Uint()) // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS

		val = reflect.ValueOf(pSubpassEndInfo).Elem()
		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := extension.CmdNextSubpass2(
		commandBuffer,
		khr_create_renderpass2.SubpassBeginOptions{
			Contents: core1_0.SubpassContentsSecondaryCommandBuffers,
		},
		khr_create_renderpass2.SubpassEndOptions{})
	require.NoError(t, err)
}

func TestVulkanExtension_CreateRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)

	extDriver.EXPECT().VkCreateRenderPass2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *khr_create_renderpass2_driver.VkRenderPassCreateInfo2KHR,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		require.Equal(t, uint64(1), val.FieldByName("attachmentCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("dependencyCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("correlatedViewMaskCount").Uint())

		attachmentsPtr := (*khr_create_renderpass2_driver.VkAttachmentDescription2KHR)(val.FieldByName("pAttachments").UnsafePointer())
		attachmentsSlice := unsafe.Slice(attachmentsPtr, 1)
		attachment := reflect.ValueOf(attachmentsSlice).Index(0)

		require.Equal(t, uint64(1000109000), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
		require.True(t, attachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())          // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
		require.Equal(t, uint64(68), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32
		require.Equal(t, uint64(8), attachment.FieldByName("samples").Uint())        // VK_SAMPLE_COUNT_8_BIT
		require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_CLEAR
		require.Equal(t, uint64(1), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_DONT_CARE
		require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_DONT_CARE
		require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_STORE
		require.Equal(t, uint64(4), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
		require.Equal(t, uint64(8), attachment.FieldByName("finalLayout").Uint())    // VK_IMAGE_LAYOUT_PREINITIALIZED

		viewMasks := (*uint32)(val.FieldByName("pCorrelatedViewMasks").UnsafePointer())
		viewMaskSlice := ([]uint32)(unsafe.Slice(viewMasks, 3))
		require.Equal(t, []uint32{29, 31, 37}, viewMaskSlice)

		subpassPtr := (*khr_create_renderpass2_driver.VkSubpassDescription2KHR)(val.FieldByName("pSubpasses").UnsafePointer())
		subpassSlice := unsafe.Slice(subpassPtr, 1)
		subpass := reflect.ValueOf(subpassSlice).Index(0)

		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
		require.True(t, subpass.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), subpass.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
		require.Equal(t, uint64(1), subpass.FieldByName("viewMask").Uint())
		require.Equal(t, uint64(2), subpass.FieldByName("inputAttachmentCount").Uint())
		require.Equal(t, uint64(1), subpass.FieldByName("colorAttachmentCount").Uint())
		require.Equal(t, uint64(2), subpass.FieldByName("preserveAttachmentCount").Uint())

		preserveAttachments := (*uint32)(subpass.FieldByName("pPreserveAttachments").UnsafePointer())
		preserveAttachmentSlice := ([]uint32)(unsafe.Slice(preserveAttachments, 2))
		require.Equal(t, []uint32{59, 61}, preserveAttachmentSlice)

		inputAttachmentPtr := (*khr_create_renderpass2_driver.VkAttachmentReference2KHR)(subpass.FieldByName("pInputAttachments").UnsafePointer())
		inputAttachmentSlice := ([]khr_create_renderpass2_driver.VkAttachmentReference2KHR)(unsafe.Slice(inputAttachmentPtr, 2))
		inputAttachment := reflect.ValueOf(inputAttachmentSlice)
		require.Equal(t, uint64(1000109001), inputAttachment.Index(0).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, inputAttachment.Index(0).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), inputAttachment.Index(0).FieldByName("attachment").Uint())
		require.Equal(t, uint64(6), inputAttachment.Index(0).FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
		require.Equal(t, uint64(4), inputAttachment.Index(0).FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT

		require.Equal(t, uint64(1000109001), inputAttachment.Index(1).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, inputAttachment.Index(1).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(5), inputAttachment.Index(1).FieldByName("attachment").Uint())
		require.Equal(t, uint64(6), inputAttachment.Index(1).FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
		require.Equal(t, uint64(8), inputAttachment.Index(1).FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT

		colorAttachment := subpass.FieldByName("pColorAttachments").Elem()
		require.Equal(t, uint64(1000109001), colorAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, colorAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(41), colorAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(8), colorAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_PREINITIALIZED
		require.Equal(t, uint64(1), colorAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

		resolveAttachment := subpass.FieldByName("pResolveAttachments").Elem()
		require.Equal(t, uint64(1000109001), resolveAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, resolveAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(43), resolveAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(1), resolveAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_GENERAL
		require.Equal(t, uint64(2), resolveAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT

		depthAttachment := subpass.FieldByName("pDepthStencilAttachment").Elem()
		require.Equal(t, uint64(1000109001), depthAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, depthAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(47), depthAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(7), depthAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		require.Equal(t, uint64(1), depthAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

		dependenciesPtr := (*khr_create_renderpass2_driver.VkSubpassDependency2KHR)(val.FieldByName("pDependencies").UnsafePointer())
		dependenciesSlice := ([]khr_create_renderpass2_driver.VkSubpassDependency2KHR)(unsafe.Slice(dependenciesPtr, 2))
		val = reflect.ValueOf(dependenciesSlice)
		require.Equal(t, uint64(1000109003), val.Index(0).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
		require.True(t, val.Index(0).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.Index(0).FieldByName("srcSubpass").Uint())
		require.Equal(t, uint64(11), val.Index(0).FieldByName("dstSubpass").Uint())
		require.Equal(t, uint64(0x800), val.Index(0).FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
		require.Equal(t, uint64(2), val.Index(0).FieldByName("dstStageMask").Uint())      // VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
		require.Equal(t, uint64(2), val.Index(0).FieldByName("srcAccessMask").Uint())     // VK_ACCESS_INDEX_READ_BIT
		require.Equal(t, uint64(0x100), val.Index(0).FieldByName("dstAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
		require.Equal(t, uint64(1), val.Index(0).FieldByName("dependencyFlags").Uint())   // VK_DEPENDENCY_BY_REGION_BIT
		require.Equal(t, int64(13), val.Index(0).FieldByName("viewOffset").Int())

		require.Equal(t, uint64(1000109003), val.Index(1).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
		require.True(t, val.Index(1).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(17), val.Index(1).FieldByName("srcSubpass").Uint())
		require.Equal(t, uint64(19), val.Index(1).FieldByName("dstSubpass").Uint())
		require.Equal(t, uint64(0x40), val.Index(1).FieldByName("srcStageMask").Uint())   // VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT
		require.Equal(t, uint64(0x4000), val.Index(1).FieldByName("dstStageMask").Uint()) // VK_PIPELINE_STAGE_HOST_BIT
		require.Equal(t, uint64(0x80), val.Index(1).FieldByName("srcAccessMask").Uint())  // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
		require.Equal(t, uint64(0x200), val.Index(1).FieldByName("dstAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
		require.Equal(t, uint64(0), val.Index(1).FieldByName("dependencyFlags").Uint())
		require.Equal(t, int64(23), val.Index(1).FieldByName("viewOffset").Int())

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := extension.CreateRenderPass2(
		device,
		khr_create_renderpass2.RenderPassCreateOptions{
			Flags: 0,
			Attachments: []khr_create_renderpass2.AttachmentDescriptionOptions{
				{
					Flags:          core1_0.AttachmentDescriptionMayAlias,
					Format:         core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
					Samples:        core1_0.Samples8,
					LoadOp:         core1_0.LoadOpClear,
					StoreOp:        core1_0.StoreOpDontCare,
					StencilLoadOp:  core1_0.LoadOpDontCare,
					StencilStoreOp: core1_0.StoreOpStore,
					InitialLayout:  core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
					FinalLayout:    core1_0.ImageLayoutPreInitialized,
				},
			},
			Subpasses: []khr_create_renderpass2.SubpassDescriptionOptions{
				{
					Flags:             0,
					PipelineBindPoint: core1_0.BindCompute,
					ViewMask:          1,
					InputAttachments: []khr_create_renderpass2.AttachmentReferenceOptions{
						{
							Attachment: 3,
							Layout:     core1_0.ImageLayoutTransferSrcOptimal,
							AspectMask: core1_0.AspectStencil,
						},
						{
							Attachment: 5,
							Layout:     core1_0.ImageLayoutTransferSrcOptimal,
							AspectMask: core1_0.AspectMetadata,
						},
					},
					ColorAttachments: []khr_create_renderpass2.AttachmentReferenceOptions{
						{
							Attachment: 41,
							Layout:     core1_0.ImageLayoutPreInitialized,
							AspectMask: core1_0.AspectColor,
						},
					},
					ResolveAttachments: []khr_create_renderpass2.AttachmentReferenceOptions{
						{
							Attachment: 43,
							Layout:     core1_0.ImageLayoutGeneral,
							AspectMask: core1_0.AspectDepth,
						},
					},
					DepthStencilAttachment: &khr_create_renderpass2.AttachmentReferenceOptions{
						Attachment: 47,
						Layout:     core1_0.ImageLayoutTransferDstOptimal,
						AspectMask: core1_0.AspectColor,
					},
					PreserveAttachments: []int{59, 61},
				},
			},
			Dependencies: []khr_create_renderpass2.SubpassDependencyOptions{
				{
					SrcSubpassIndex: 7,
					DstSubpassIndex: 11,
					SrcStageMask:    core1_0.PipelineStageComputeShader,
					DstStageMask:    core1_0.PipelineStageDrawIndirect,
					SrcAccessMask:   core1_0.AccessIndexRead,
					DstAccessMask:   core1_0.AccessColorAttachmentWrite,
					DependencyFlags: core1_0.DependencyByRegion,
					ViewOffset:      13,
				},
				{
					SrcSubpassIndex: 17,
					DstSubpassIndex: 19,
					SrcStageMask:    core1_0.PipelineStageGeometryShader,
					DstStageMask:    core1_0.PipelineStageHost,
					SrcAccessMask:   core1_0.AccessColorAttachmentRead,
					DstAccessMask:   core1_0.AccessDepthStencilAttachmentRead,
					DependencyFlags: 0,
					ViewOffset:      23,
				},
			},
			CorrelatedViewMasks: []uint32{29, 31, 37},
		}, nil)
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())
}
