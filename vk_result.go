package VKng

/*
#include <vulkan/vulkan.h>
*/
import "C"
import "github.com/palantir/stacktrace"

type VKResult int

const (
	VKSuccess                                  VKResult = C.VK_SUCCESS
	VKNotReady                                 VKResult = C.VK_NOT_READY
	VKTimeout                                  VKResult = C.VK_TIMEOUT
	VKEventSet                                 VKResult = C.VK_EVENT_SET
	VKEventReset                               VKResult = C.VK_EVENT_RESET
	VKIncomplete                               VKResult = C.VK_INCOMPLETE
	VKErrorOutOfHostMemory                     VKResult = C.VK_ERROR_OUT_OF_HOST_MEMORY
	VKErrorOutOfDeviceMemory                   VKResult = C.VK_ERROR_OUT_OF_DEVICE_MEMORY
	VKErrorInitializationFailed                VKResult = C.VK_ERROR_INITIALIZATION_FAILED
	VKErrorDeviceLost                          VKResult = C.VK_ERROR_DEVICE_LOST
	VKErrorMemoryMapFailed                     VKResult = C.VK_ERROR_MEMORY_MAP_FAILED
	VKErrorLayerNotPresent                     VKResult = C.VK_ERROR_LAYER_NOT_PRESENT
	VKErrorExtensionNotPresent                 VKResult = C.VK_ERROR_EXTENSION_NOT_PRESENT
	VKErrorFeatureNotPresent                   VKResult = C.VK_ERROR_FEATURE_NOT_PRESENT
	VKErrorIncompatibleDriver                  VKResult = C.VK_ERROR_INCOMPATIBLE_DRIVER
	VKErrorTooManyObjects                      VKResult = C.VK_ERROR_TOO_MANY_OBJECTS
	VKErrorFormatNotSupported                  VKResult = C.VK_ERROR_FORMAT_NOT_SUPPORTED
	VKErrorFragmentedPool                      VKResult = C.VK_ERROR_FRAGMENTED_POOL
	VKErrorUnknown                             VKResult = C.VK_ERROR_UNKNOWN
	VKErrorOutOfPoolMemory                     VKResult = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
	VKErrorInvalidExternalHandle               VKResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
	VKErrorFragmentation                       VKResult = C.VK_ERROR_FRAGMENTATION_EXT
	VKErrorInvalidDeviceAddress                VKResult = C.VK_ERROR_INVALID_DEVICE_ADDRESS_EXT
	VKErrorSurfaceLost                         VKResult = C.VK_ERROR_SURFACE_LOST_KHR
	VKErrorNativeWindowInUse                   VKResult = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
	VKSuboptimal                               VKResult = C.VK_SUBOPTIMAL_KHR
	VKErrorOutOfDate                           VKResult = C.VK_ERROR_OUT_OF_DATE_KHR
	VKErrorIncompatibleDisplay                 VKResult = C.VK_ERROR_INCOMPATIBLE_DISPLAY_KHR
	VKErrorValidationFailed                    VKResult = C.VK_ERROR_VALIDATION_FAILED_EXT
	VKErrorInvalidShader                       VKResult = C.VK_ERROR_INVALID_SHADER_NV
	VKErrorInvalidDRMFormatModifierPlaneLayout VKResult = C.VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT
	VKErrorNotPermitted                        VKResult = C.VK_ERROR_NOT_PERMITTED_EXT
	VKErrorFullScreenExclusiveModeLost         VKResult = C.VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
	VKThreadIdle                               VKResult = C.VK_THREAD_IDLE_KHR
	VKThreadDone                               VKResult = C.VK_THREAD_DONE_KHR
	VKOperationDeferred                        VKResult = C.VK_OPERATION_DEFERRED_KHR
	VKOperationNotDeferred                     VKResult = C.VK_OPERATION_NOT_DEFERRED_KHR
	VKPipelineCompileRequired                  VKResult = C.VK_PIPELINE_COMPILE_REQUIRED_EXT
)

var vkResultToString = map[VKResult]string{
	VKSuccess:                                  "Success",
	VKNotReady:                                 "Not Ready",
	VKTimeout:                                  "Timeout",
	VKEventSet:                                 "Event Set",
	VKEventReset:                               "Event Reset",
	VKIncomplete:                               "Incomplete",
	VKErrorOutOfHostMemory:                     "out of host memory",
	VKErrorOutOfDeviceMemory:                   "out of device memory",
	VKErrorInitializationFailed:                "initialization failed",
	VKErrorDeviceLost:                          "device lost",
	VKErrorMemoryMapFailed:                     "memory map failed",
	VKErrorLayerNotPresent:                     "layer not present",
	VKErrorExtensionNotPresent:                 "extension not present",
	VKErrorFeatureNotPresent:                   "feature not present",
	VKErrorIncompatibleDriver:                  "incompatible driver",
	VKErrorTooManyObjects:                      "too many objects",
	VKErrorFormatNotSupported:                  "format not supported",
	VKErrorFragmentedPool:                      "fragmented pool",
	VKErrorUnknown:                             "unknown",
	VKErrorOutOfPoolMemory:                     "out of pool memory",
	VKErrorInvalidExternalHandle:               "invalid external handle",
	VKErrorFragmentation:                       "fragmentation",
	VKErrorInvalidDeviceAddress:                "invalid device address",
	VKErrorSurfaceLost:                         "surface lost",
	VKErrorNativeWindowInUse:                   "native window in use",
	VKSuboptimal:                               "Suboptimal",
	VKErrorOutOfDate:                           "out of date",
	VKErrorIncompatibleDisplay:                 "incompatible display",
	VKErrorValidationFailed:                    "validation failed",
	VKErrorInvalidShader:                       "invalid shader",
	VKErrorInvalidDRMFormatModifierPlaneLayout: "invalid drm format modifier plane layout",
	VKErrorNotPermitted:                        "not permitted",
	VKErrorFullScreenExclusiveModeLost:         "full-screen exclusive mode lost",
	VKThreadIdle:                               "Thread Idle",
	VKThreadDone:                               "Thread Done",
	VKOperationDeferred:                        "Operation Deferred",
	VKOperationNotDeferred:                     "Operation Not Deferred",
	VKPipelineCompileRequired:                  "Pipeline Compile Required",
}

func (r VKResult) String() string {
	return vkResultToString[r]
}

func (r VKResult) ToError() error {
	if r >= 0 {
		// All VKError* are <0
		return nil
	}

	return stacktrace.NewErrorWithCode(stacktrace.ErrorCode(r), "vulkan error: %s", r.String())
}
