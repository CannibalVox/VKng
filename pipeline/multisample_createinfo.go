package pipeline

import "github.com/CannibalVox/VKng/core"

type MultisampleOptions struct {
	RasterizationSamples core.SampleCounts

	SampleShading    bool
	MinSampleShading float32
	SampleMask       []uint32

	AlphaToCoverage bool
	AlphaToOne      bool

	Next core.Options
}
