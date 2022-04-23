package khr_external_fence

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type FenceImportFlags int32

var fenceImportFlagsMapping = common.NewFlagStringMapping[FenceImportFlags]()

func (f FenceImportFlags) Register(str string) {
	fenceImportFlagsMapping.Register(f, str)
}

func (f FenceImportFlags) String() string {
	return fenceImportFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_FENCE_EXTENSION_NAME

	FenceImportTemporary FenceImportFlags = C.VK_FENCE_IMPORT_TEMPORARY_BIT
)

func init() {
	FenceImportTemporary.Register("Temporary")
}
