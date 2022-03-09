package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

type MessageTypes int32

var messageTypesMapping = common.NewFlagStringMapping[MessageTypes]()

func (f MessageTypes) Register(str string) {
	messageTypesMapping.Register(f, str)
}
func (f MessageTypes) String() string {
	return messageTypesMapping.FlagsToString(f)
}

////

type MessageSeverities int32

var messageSeveritiesMapping = common.NewFlagStringMapping[MessageSeverities]()

func (f MessageSeverities) Register(str string) {
	messageSeveritiesMapping.Register(f, str)
}
func (f MessageSeverities) String() string {
	return messageSeveritiesMapping.FlagsToString(f)
}

////

const (
	TypeGeneral     MessageTypes = C.VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT
	TypeValidation  MessageTypes = C.VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
	TypePerformance MessageTypes = C.VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT
	TypeAll         MessageTypes = TypeGeneral | TypeValidation | TypePerformance

	SeverityVerbose MessageSeverities = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT
	SeverityInfo    MessageSeverities = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT
	SeverityWarning MessageSeverities = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT
	SeverityError   MessageSeverities = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
	SeverityAll     MessageSeverities = SeverityVerbose | SeverityInfo | SeverityWarning | SeverityError

	ObjectTypeDebugUtilsMessenger common.ObjectType = C.VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT
)

func init() {
	TypeGeneral.Register("General")
	TypeValidation.Register("Validation")
	TypePerformance.Register("Performance")

	SeverityVerbose.Register("Verbose")
	SeverityInfo.Register("Info")
	SeverityWarning.Register("Warning")
	SeverityError.Register("Error")

	ObjectTypeDebugUtilsMessenger.Register("Debug Utils Messenger")
}
