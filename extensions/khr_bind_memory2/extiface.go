package khr_bind_memory2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_bind_memory2

type Extension interface {
	BindBufferMemory(device core1_0.Device, options []BindBufferMemoryOptions) (common.VkResult, error)
	BindImageMemory(device core1_0.Device, options []BindImageMemoryOptions) (common.VkResult, error)
}
