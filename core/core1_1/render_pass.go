package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanRenderPass struct {
	core1_0.RenderPass
}

func PromoteRenderPass(renderPass core1_0.RenderPass) RenderPass {
	if !renderPass.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return renderPass.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(renderPass.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanRenderPass{renderPass}
		}).(RenderPass)
}
