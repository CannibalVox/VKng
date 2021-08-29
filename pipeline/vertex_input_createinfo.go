package pipeline

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core"

type InputRate int32

const (
	Vertex   InputRate = C.VK_VERTEX_INPUT_RATE_VERTEX
	Instance InputRate = C.VK_VERTEX_INPUT_RATE_INSTANCE
)

var inputRateToString = map[InputRate]string{
	Vertex:   "Vertex",
	Instance: "Instance",
}

func (r InputRate) String() string {
	return inputRateToString[r]
}

type VertexBindingDescription struct {
	InputRate InputRate
	Binding   uint32
	Stride    uint32
}

type VertexAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   core.ColorFormat
	Offset   uint32
}

type VertexInputOptions struct {
	VertexBindingDescriptions   []VertexBindingDescription
	VertexAttributeDescriptions []VertexAttributeDescription

	Next core.Options
}
