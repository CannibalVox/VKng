package ext_sampler_filter_minmax

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/core1_0"

type SamplerReductionMode int32

var samplerReductionModeMapping = make(map[SamplerReductionMode]string)

func (e SamplerReductionMode) Register(str string) {
	samplerReductionModeMapping[e] = str
}

func (e SamplerReductionMode) String() string {
	return samplerReductionModeMapping[e]
}

////

const (
	ExtensionName string = C.VK_EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME

	FormatFeatureSampledImageFilterMinmax core1_0.FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT

	SamplerReductionModeMax             SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MAX_EXT
	SamplerReductionModeMin             SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MIN_EXT
	SamplerReductionModeWeightedAverage SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT
)

func init() {
	FormatFeatureSampledImageFilterMinmax.Register("Sampled Image Filter Min-Max")

	SamplerReductionModeMin.Register("Min")
	SamplerReductionModeMax.Register("Max")
	SamplerReductionModeWeightedAverage.Register("Weighted Average")
}
