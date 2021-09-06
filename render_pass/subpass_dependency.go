package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
)

const SubpassExternal = int(C.VK_SUBPASS_EXTERNAL)

type SubPassDependency struct {
	Flags VKng.DependencyFlags

	SrcSubPassIndex int
	DstSubPassIndex int

	SrcStageMask VKng.PipelineStages
	DstStageMask VKng.PipelineStages

	SrcAccess VKng.AccessFlags
	DstAccess VKng.AccessFlags
}
