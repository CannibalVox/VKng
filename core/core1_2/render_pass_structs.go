package core1_2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type SubpassBeginOptions struct {
	Contents core1_0.SubpassContents

	common.NextOptions
}

func (o SubpassBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassBeginInfo{})))
	}

	info := (*C.VkSubpassBeginInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
	info.pNext = next
	info.contents = C.VkSubpassContents(o.Contents)

	return preallocatedPointer, nil
}

////

type SubpassEndOptions struct {
	common.NextOptions
}

func (o SubpassEndOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassEndInfo{})))
	}

	info := (*C.VkSubpassEndInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_END_INFO
	info.pNext = next

	return preallocatedPointer, nil
}

////

type AttachmentDescriptionOptions struct {
	Flags          core1_0.AttachmentDescriptionFlags
	Format         core1_0.DataFormat
	Samples        core1_0.SampleCounts
	LoadOp         core1_0.AttachmentLoadOp
	StoreOp        core1_0.AttachmentStoreOp
	StencilLoadOp  core1_0.AttachmentLoadOp
	StencilStoreOp core1_0.AttachmentStoreOp
	InitialLayout  core1_0.ImageLayout
	FinalLayout    core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescriptionOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescription2{})))
	}

	info := (*C.VkAttachmentDescription2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2
	info.pNext = next
	info.flags = C.VkAttachmentDescriptionFlags(o.Flags)
	info.format = C.VkFormat(o.Format)
	info.samples = C.VkSampleCountFlagBits(o.Samples)
	info.loadOp = C.VkAttachmentLoadOp(o.LoadOp)
	info.storeOp = C.VkAttachmentStoreOp(o.StoreOp)
	info.stencilLoadOp = C.VkAttachmentLoadOp(o.StencilLoadOp)
	info.stencilStoreOp = C.VkAttachmentStoreOp(o.StencilStoreOp)
	info.initialLayout = C.VkImageLayout(o.InitialLayout)
	info.finalLayout = C.VkImageLayout(o.FinalLayout)

	return preallocatedPointer, nil
}

////

type AttachmentReferenceOptions struct {
	Attachment int
	Layout     core1_0.ImageLayout
	AspectMask core1_0.ImageAspectFlags

	common.NextOptions
}

func (o AttachmentReferenceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReference2{})))
	}

	info := (*C.VkAttachmentReference2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
	info.pNext = next
	info.attachment = C.uint32_t(o.Attachment)
	info.layout = C.VkImageLayout(o.Layout)
	info.aspectMask = C.VkImageAspectFlags(o.AspectMask)

	return preallocatedPointer, nil
}

////

type SubpassDescriptionOptions struct {
	Flags                  core1_0.SubPassDescriptionFlags
	PipelineBindPoint      core1_0.PipelineBindPoint
	ViewMask               uint32
	InputAttachments       []AttachmentReferenceOptions
	ColorAttachments       []AttachmentReferenceOptions
	ResolveAttachments     []AttachmentReferenceOptions
	DepthStencilAttachment *AttachmentReferenceOptions
	PreserveAttachments    []int

	common.NextOptions
}

func (o SubpassDescriptionOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescription2{})))
	}

	info := (*C.VkSubpassDescription2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2
	info.pNext = next
	info.flags = C.VkSubpassDescriptionFlags(o.Flags)
	info.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	info.viewMask = C.uint32_t(o.ViewMask)

	inputAttachmentCount := len(o.InputAttachments)
	colorAttachmentCount := len(o.ColorAttachments)
	resolveAttachmentCount := len(o.ResolveAttachments)
	preserveAttachmentCount := len(o.PreserveAttachments)

	if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
		return nil, errors.Newf("in this subpass, %d color attachments are defined, but %d resolve attachments are defined- they should be equal", colorAttachmentCount, resolveAttachmentCount)
	}

	info.inputAttachmentCount = C.uint32_t(inputAttachmentCount)
	info.pInputAttachments = nil
	info.colorAttachmentCount = C.uint32_t(colorAttachmentCount)
	info.pColorAttachments = nil
	info.pResolveAttachments = nil
	info.pDepthStencilAttachment = nil
	info.preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
	info.pPreserveAttachments = nil

	var err error
	if inputAttachmentCount > 0 {
		info.pInputAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReferenceOptions](allocator, o.InputAttachments)
		if err != nil {
			return nil, err
		}
	}

	if colorAttachmentCount > 0 {
		info.pColorAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReferenceOptions](allocator, o.ColorAttachments)
		if err != nil {
			return nil, err
		}

		info.pResolveAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReferenceOptions](allocator, o.ResolveAttachments)
		if err != nil {
			return nil, err
		}
	}

	if o.DepthStencilAttachment != nil {
		depthStencilPtr, err := common.AllocOptions(allocator, o.DepthStencilAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilAttachment = (*C.VkAttachmentReference2)(depthStencilPtr)
	}

	if preserveAttachmentCount > 0 {
		attachmentsPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		attachmentsSlice := unsafe.Slice(attachmentsPtr, preserveAttachmentCount)
		for i := 0; i < preserveAttachmentCount; i++ {
			attachmentsSlice[i] = C.uint32_t(o.PreserveAttachments[i])
		}
		info.pPreserveAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}

////

type SubpassDependencyOptions struct {
	SrcSubpassIndex int
	DstSubpassIndex int
	SrcStageMask    core1_0.PipelineStages
	DstStageMask    core1_0.PipelineStages
	SrcAccessMask   core1_0.AccessFlags
	DstAccessMask   core1_0.AccessFlags
	DependencyFlags core1_0.DependencyFlags
	ViewOffset      int

	common.NextOptions
}

func (o SubpassDependencyOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDependency2{})))
	}

	info := (*C.VkSubpassDependency2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2
	info.pNext = next
	info.srcSubpass = C.uint32_t(o.SrcSubpassIndex)
	info.dstSubpass = C.uint32_t(o.DstSubpassIndex)
	info.srcStageMask = C.VkPipelineStageFlags(o.SrcStageMask)
	info.dstStageMask = C.VkPipelineStageFlags(o.DstStageMask)
	info.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	info.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	info.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)
	info.viewOffset = C.int32_t(o.ViewOffset)

	return preallocatedPointer, nil
}

////

type RenderPassCreateOptions struct {
	Flags core1_0.RenderPassCreateFlags

	Attachments  []AttachmentDescriptionOptions
	Subpasses    []SubpassDescriptionOptions
	Dependencies []SubpassDependencyOptions

	CorrelatedViewMasks []uint32

	common.NextOptions
}

func (o RenderPassCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo2{})))
	}

	info := (*C.VkRenderPassCreateInfo2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
	info.pNext = next
	info.flags = C.VkRenderPassCreateFlags(o.Flags)

	attachmentCount := len(o.Attachments)
	subpassCount := len(o.Subpasses)
	dependencyCount := len(o.Dependencies)
	viewMaskCount := len(o.CorrelatedViewMasks)

	info.attachmentCount = C.uint32_t(attachmentCount)
	info.pAttachments = nil
	info.subpassCount = C.uint32_t(subpassCount)
	info.pSubpasses = nil
	info.dependencyCount = C.uint32_t(dependencyCount)
	info.pDependencies = nil
	info.correlatedViewMaskCount = C.uint32_t(viewMaskCount)
	info.pCorrelatedViewMasks = nil

	var err error
	if attachmentCount > 0 {
		info.pAttachments, err = common.AllocOptionSlice[C.VkAttachmentDescription2, AttachmentDescriptionOptions](allocator, o.Attachments)
		if err != nil {
			return nil, err
		}
	}

	if subpassCount > 0 {
		info.pSubpasses, err = common.AllocOptionSlice[C.VkSubpassDescription2, SubpassDescriptionOptions](allocator, o.Subpasses)
		if err != nil {
			return nil, err
		}
	}

	if dependencyCount > 0 {
		info.pDependencies, err = common.AllocOptionSlice[C.VkSubpassDependency2, SubpassDependencyOptions](allocator, o.Dependencies)
		if err != nil {
			return nil, err
		}
	}

	if viewMaskCount > 0 {
		viewMaskPtr := (*C.uint32_t)(allocator.Malloc(viewMaskCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := unsafe.Slice(viewMaskPtr, viewMaskCount)
		for i := 0; i < viewMaskCount; i++ {
			viewMaskSlice[i] = C.uint32_t(o.CorrelatedViewMasks[i])
		}
		info.pCorrelatedViewMasks = viewMaskPtr
	}

	return preallocatedPointer, nil
}

////

type AttachmentDescriptionStencilLayoutOptions struct {
	StencilInitialLayout core1_0.ImageLayout
	StencilFinalLayout   core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescriptionStencilLayoutOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescriptionStencilLayout{})))
	}

	info := (*C.VkAttachmentDescriptionStencilLayout)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT
	info.pNext = next
	info.stencilInitialLayout = C.VkImageLayout(o.StencilInitialLayout)
	info.stencilFinalLayout = C.VkImageLayout(o.StencilFinalLayout)

	return preallocatedPointer, nil
}

////

type AttachmentReferenceStencilLayoutOptions struct {
	StencilLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentReferenceStencilLayoutOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReferenceStencilLayout{})))
	}

	info := (*C.VkAttachmentReferenceStencilLayout)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT
	info.pNext = next
	info.stencilLayout = C.VkImageLayout(o.StencilLayout)

	return preallocatedPointer, nil
}

////

type RenderPassAttachmentBeginOptions struct {
	Attachments []core1_0.ImageView

	common.NextOptions
}

func (o RenderPassAttachmentBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassAttachmentBeginInfo{})))
	}

	info := (*C.VkRenderPassAttachmentBeginInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO
	info.pNext = next

	count := len(o.Attachments)
	info.attachmentCount = C.uint32_t(count)
	info.pAttachments = nil

	if count > 0 {
		info.pAttachments = (*C.VkImageView)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentSlice := unsafe.Slice(info.pAttachments, count)
		for i := 0; i < count; i++ {
			attachmentSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}
	}

	return preallocatedPointer, nil
}

////

type SubpassDescriptionDepthStencilResolveOptions struct {
	DepthResolveMode              ResolveModeFlags
	StencilResolveMode            ResolveModeFlags
	DepthStencilResolveAttachment *AttachmentReferenceOptions

	common.NextOptions
}

func (o SubpassDescriptionDepthStencilResolveOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescriptionDepthStencilResolve{})))
	}

	info := (*C.VkSubpassDescriptionDepthStencilResolve)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE
	info.pNext = next
	info.depthResolveMode = C.VkResolveModeFlagBits(o.DepthResolveMode)
	info.stencilResolveMode = C.VkResolveModeFlagBits(o.StencilResolveMode)
	info.pDepthStencilResolveAttachment = nil

	if o.DepthStencilResolveAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.DepthStencilResolveAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilResolveAttachment = (*C.VkAttachmentReference2)(attachment)
	}

	return preallocatedPointer, nil
}
