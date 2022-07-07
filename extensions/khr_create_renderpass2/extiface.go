package khr_create_renderpass2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_create_renderpass2

type Extension interface {
	CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error
	CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd SubpassEndInfo) error
	CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error

	CreateRenderPass2(device core1_0.Device, allocator *driver.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error)
}
