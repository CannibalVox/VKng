package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type MessageType uint32

const (
	TypeGeneral     MessageType = C.VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT
	TypeValidation  MessageType = C.VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
	TypePerformance MessageType = C.VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT
	TypeAll         MessageType = TypeGeneral | TypeValidation | TypePerformance
)

func (f MessageType) String() string {
	hasOne := false
	sb := strings.Builder{}

	if f&TypeGeneral != 0 {
		sb.WriteString("General")
		hasOne = true
	}

	if f&TypeValidation != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("Validation")
		hasOne = true
	}

	if f&TypePerformance != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("Performance")
	}

	return sb.String()
}

type MessageSeverity uint32

const (
	SeverityVerbose MessageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT
	SeverityInfo    MessageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT
	SeverityWarning MessageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT
	SeverityError   MessageSeverity = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
	SeverityAll     MessageSeverity = SeverityVerbose | SeverityInfo | SeverityWarning | SeverityError
)

func (s MessageSeverity) String() string {
	hasOne := false
	sb := strings.Builder{}

	if s&SeverityVerbose != 0 {
		sb.WriteString("Verbose")
		hasOne = true
	}

	if s&SeverityInfo != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("Info")
		hasOne = true
	}

	if s&SeverityWarning != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("Warning")
		hasOne = true
	}

	if s&SeverityError != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("Error")
	}

	return sb.String()
}
