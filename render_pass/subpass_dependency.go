package render_pass

import "github.com/CannibalVox/VKng/core"

type SubPassDependency struct {
	Flags core.DependencyFlags

	SrcSubPassIndex int
	DstSubPassIndex int

	SrcStageMask core.PipelineStages
	DstStageMask core.PipelineStages

	SrcAccess core.AccessFlags
	DstAccess core.AccessFlags
}
