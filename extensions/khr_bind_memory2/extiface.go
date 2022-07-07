package khr_bind_memory2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_bind_memory2

type Extension interface {
	BindBufferMemory2(device core1_0.Device, options []BindBufferMemoryInfo) (common.VkResult, error)
	BindImageMemory2(device core1_0.Device, options []BindImageMemoryInfo) (common.VkResult, error)
}
