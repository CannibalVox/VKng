package render_pass

import "github.com/CannibalVox/VKng/core"

type AttachmentDescription struct {
	Flags   AttachmentDescriptionFlags
	Format  core.ColorFormat
	Samples core.SampleCounts

	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp

	InitialLayout core.ImageLayout
	FinalLayout   core.ImageLayout
}
