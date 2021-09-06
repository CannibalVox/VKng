package render_pass

import (
	"github.com/CannibalVox/VKng"
)

type AttachmentDescription struct {
	Flags   VKng.AttachmentDescriptionFlags
	Format  VKng.DataFormat
	Samples VKng.SampleCounts

	LoadOp         VKng.AttachmentLoadOp
	StoreOp        VKng.AttachmentStoreOp
	StencilLoadOp  VKng.AttachmentLoadOp
	StencilStoreOp VKng.AttachmentStoreOp

	InitialLayout VKng.ImageLayout
	FinalLayout   VKng.ImageLayout
}
