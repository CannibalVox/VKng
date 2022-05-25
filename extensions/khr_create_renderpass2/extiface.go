package khr_create_renderpass2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_create_renderpass2

type Extension interface {
	CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginOptions, subpassBegin SubpassBeginOptions) error
	CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd SubpassEndOptions) error
	CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin SubpassBeginOptions, subpassEnd SubpassEndOptions) error

	CreateRenderPass2(device core1_0.Device, options RenderPassCreateOptions, allocator *driver.AllocationCallbacks) (core1_0.RenderPass, common.VkResult, error)
}
