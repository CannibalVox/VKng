package khr_get_memory_requirements2

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_memory_requirements2

type Extension interface {
	BufferMemoryRequirements2(device core1_0.Device, o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageMemoryRequirements2(device core1_0.Device, o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageSparseMemoryRequirements2(device core1_0.Device, o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)
}
