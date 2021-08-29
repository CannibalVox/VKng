package pipeline

import "C"
import "github.com/CannibalVox/VKng/core"

type ColorBlendAttachment struct {
	BlendEnabled bool

	SrcColor     BlendFactor
	DstColor     BlendFactor
	ColorBlendOp BlendOp

	SrcAlpha     BlendFactor
	DstAlpha     BlendFactor
	AlphaBlendOp BlendOp

	WriteMask ColorComponents
}

type ColorBlendOptions struct {
	LogicOpEnabled bool
	LogicOp        core.LogicOp

	BlendConstants [4]float32
	Attachments    []ColorBlendAttachment

	Next core.Options
}
