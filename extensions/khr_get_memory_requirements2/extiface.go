package khr_get_memory_requirements2

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_memory_requirements2

type Extension interface {
	BufferMemoryRequirements(device core1_0.Device, o BufferMemoryRequirementsOptions, out *MemoryRequirementsOutData) error
	ImageMemoryRequirements(device core1_0.Device, o ImageMemoryRequirementsOptions, out *MemoryRequirementsOutData) error
	SparseImageMemoryRequirements(device core1_0.Device, o ImageSparseMemoryRequirementsOptions, outDataFactory func() *SparseImageMemoryRequirementsOutData) ([]*SparseImageMemoryRequirementsOutData, error)
}
