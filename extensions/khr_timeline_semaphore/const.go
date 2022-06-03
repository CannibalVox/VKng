package khr_timeline_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type SemaphoreType int32

var semaphoreTypeMapping = make(map[SemaphoreType]string)

func (e SemaphoreType) Register(str string) {
	semaphoreTypeMapping[e] = str
}

func (e SemaphoreType) String() string {
	return semaphoreTypeMapping[e]
}

////

type SemaphoreWaitFlags int32

var semaphoreWaitFlagsMapping = common.NewFlagStringMapping[SemaphoreWaitFlags]()

func (f SemaphoreWaitFlags) Register(str string) {
	semaphoreWaitFlagsMapping.Register(f, str)
}
func (f SemaphoreWaitFlags) String() string {
	return semaphoreWaitFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_TIMELINE_SEMAPHORE_EXTENSION_NAME

	SemaphoreTypeBinary   SemaphoreType = C.VK_SEMAPHORE_TYPE_BINARY_KHR
	SemaphoreTypeTimeline SemaphoreType = C.VK_SEMAPHORE_TYPE_TIMELINE_KHR

	SemaphoreWaitAny SemaphoreWaitFlags = C.VK_SEMAPHORE_WAIT_ANY_BIT_KHR
)

func init() {
	SemaphoreTypeBinary.Register("Binary")
	SemaphoreTypeTimeline.Register("Timeline")

	SemaphoreWaitAny.Register("Any")
}