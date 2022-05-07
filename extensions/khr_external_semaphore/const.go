package khr_external_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type SemaphoreImportFlags int32

var semaphoreImportFlagsMapping = common.NewFlagStringMapping[SemaphoreImportFlags]()

func (f SemaphoreImportFlags) Register(str string) {
	semaphoreImportFlagsMapping.Register(f, str)
}

func (f SemaphoreImportFlags) String() string {
	return semaphoreImportFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME

	SemaphoreImportTemporary SemaphoreImportFlags = C.VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR
)

func init() {
	SemaphoreImportTemporary.Register("Temporary")
}
