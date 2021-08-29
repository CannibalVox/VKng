package pipeline

import "github.com/CannibalVox/VKng/core"

type InputAssemblyOptions struct {
	Topology               core.PrimitiveTopology
	EnablePrimitiveRestart bool

	Next core.Options
}
