package pipeline

import "C"
import "github.com/CannibalVox/VKng/core"

type ColorBlendAttachment struct {
	BlendEnabled bool

	SrcColor     core.BlendFactor
	DstColor     core.BlendFactor
	ColorBlendOp core.BlendOp

	SrcAlpha     core.BlendFactor
	DstAlpha     core.BlendFactor
	AlphaBlendOp core.BlendOp

	WriteMask core.ColorComponents
}

type ColorBlendOptions struct {
	LogicOpEnabled bool
	LogicOp        core.LogicOp

	BlendConstants [4]float32
	Attachments    []ColorBlendAttachment

	Next core.Options
}
