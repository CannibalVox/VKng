package pipeline

import "github.com/CannibalVox/VKng/core"

type ViewportOptions struct {
	Viewports []core.Viewport
	Scissors  []core.Rect2D

	Next core.Options
}
