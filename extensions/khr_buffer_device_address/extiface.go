package khr_buffer_device_address

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_buffer_device_address

type Extension interface {
	GetBufferDeviceAddress(device core1_0.Device, o BufferDeviceAddressOptions) (uint64, error)
	GetBufferOpaqueCaptureAddress(device core1_0.Device, o BufferDeviceAddressOptions) (uint64, error)
	GetDeviceMemoryOpaqueCaptureAddress(device core1_0.Device, o DeviceMemoryOpaqueAddressOptions) (uint64, error)
}
