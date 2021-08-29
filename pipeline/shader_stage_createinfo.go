package pipeline

import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
)

type ShaderStageOptions struct {
	Name               string
	Stage              core.ShaderStages
	Shader             *VKng.ShaderModule
	SpecializationInfo map[uint32]interface{}

	Next core.Options
}
