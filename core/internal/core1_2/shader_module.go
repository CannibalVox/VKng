package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	core1_1.ShaderModule
}

func PromoteShaderModule(shaderModule core1_0.ShaderModule) core1_2.ShaderModule {
	if !shaderModule.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return shaderModule.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(shaderModule.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanShaderModule{core1_1.PromoteShaderModule(shaderModule)}
		}).(core1_2.ShaderModule)
}
