package common

import (
	"github.com/google/uuid"
)

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

type PhysicalDeviceProperties struct {
	Type PhysicalDeviceType
	Name string

	APIVersion    APIVersion
	DriverVersion Version
	VendorID      uint32
	DeviceID      uint32

	PipelineCacheUUID uuid.UUID
	Limits            *PhysicalDeviceLimits
	SparseProperties  *PhysicalDeviceSparseProperties
}
