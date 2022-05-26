package khr_draw_indirect_count

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_draw_indirect_count

type Extension interface {
	CmdDrawIndexedIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	CmdDrawIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
}
