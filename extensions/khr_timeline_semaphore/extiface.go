package khr_timeline_semaphore

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"time"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_timeline_semaphore

type Extension interface {
	SemaphoreCounterValue(semaphore core1_0.Semaphore) (uint64, common.VkResult, error)
	SignalSemaphore(device core1_0.Device, o SemaphoreSignalInfo) (common.VkResult, error)
	WaitSemaphores(device core1_0.Device, timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error)
}
