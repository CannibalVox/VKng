package pipeline

import (
	"github.com/CannibalVox/VKng"
)

type StencilOpState struct {
	FailOp      VKng.StencilOp
	PassOp      VKng.StencilOp
	DepthFailOp VKng.StencilOp

	CompareOp   VKng.CompareOp
	CompareMask uint32
	WriteMask   uint32

	Reference uint32
}
