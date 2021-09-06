package render_pass

import (
	"github.com/CannibalVox/VKng"
)

type SubPass struct {
	BindPoint VKng.PipelineBindPoint

	InputAttachments           []VKng.AttachmentReference
	ColorAttachments           []VKng.AttachmentReference
	ResolveAttachments         []VKng.AttachmentReference
	DepthStencilAttachments    []VKng.AttachmentReference
	PreservedAttachmentIndices []int
}
