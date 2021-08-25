package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/google/uuid"
	"strings"
	"unsafe"
)

type SampleCount int32 

const (
	Samples1 = C.VK_SAMPLE_COUNT_1_BIT
	Samples2 = C.VK_SAMPLE_COUNT_2_BIT
	Samples4 = C.VK_SAMPLE_COUNT_4_BIT
	Samples8 = C.VK_SAMPLE_COUNT_8_BIT
	Samples16 = C.VK_SAMPLE_COUNT_16_BIT
	Samples32 = C.VK_SAMPLE_COUNT_32_BIT
	Samples64 = C.VK_SAMPLE_COUNT_64_BIT
)

func (c SampleCount) String() string {
	hasOne := false
	var sb strings.Builder

	if c & Samples1 != 0 {
		sb.WriteString("1 Sample")
		hasOne = true
	}

	if c & Samples2 != 0 {
		if hasOne {sb.WriteString("|")}

		sb.WriteString("2 Samples")
		hasOne = true
	}

	if c & Samples4 != 0 {
		if hasOne {sb.WriteString("|")}

		sb.WriteString("4 Samples")
		hasOne = true
	}

	if c & Samples8 != 0 {
		if hasOne { sb.WriteString("|")}

		sb.WriteString("8 Samples")
		hasOne = true
	}

	if c & Samples16 != 0 {
		if hasOne { sb.WriteString("|")}

		sb.WriteString("16 Samples")
		hasOne = true
	}

	if c & Samples32 != 0 {
		if hasOne { sb.WriteString("|")}

		sb.WriteString("32 Samples")
		hasOne = true
	}

	if c & Samples64 != 0 {
		if hasOne { sb.WriteString("|")}

		sb.WriteString("64 Samples")
	}

	return sb.String()
}

type PhysicalDeviceType int32

const (
	Other PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	IntegratedGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	DiscreteGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	VirtualGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	CPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

var physicalDeviceTypeToString = map[PhysicalDeviceType]string {
	Other: "Other",
	IntegratedGPU: "Integrated GPU",
	DiscreteGPU: "Discrete GPU",
	VirtualGPU: "Virtual GPU",
	CPU: "CPU",
}

func (t PhysicalDeviceType) String() string {
	return physicalDeviceTypeToString[t]
}

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape bool
	ResidencyAlignedMipSize bool
	ResidencyNonResidentStrict bool
}

func createSparseProperties(p *C.VkPhysicalDeviceSparseProperties) *PhysicalDeviceSparseProperties {
	return &PhysicalDeviceSparseProperties{
		ResidencyStandard2DBlockShape: p.residencyStandard2DBlockShape != C.VK_FALSE,
		ResidencyStandard2DMultisampleBlockShape: p.residencyStandard2DMultisampleBlockShape != C.VK_FALSE,
		ResidencyStandard3DBlockShape: p.residencyStandard3DBlockShape != C.VK_FALSE,
		ResidencyNonResidentStrict: p.residencyNonResidentStrict != C.VK_FALSE,
		ResidencyAlignedMipSize: p.residencyAlignedMipSize != C.VK_FALSE,
	}
}

type PhysicalDeviceProperties struct {
	Type PhysicalDeviceType
	Name string

	APIVersion VKng.Version
	DriverVersion VKng.Version
	VendorID uint32
	DeviceID uint32

	PipelineCacheUUID uuid.UUID
	Limits *PhysicalDeviceLimits
	SparseProperties *PhysicalDeviceSparseProperties
}

func createPhysicalDeviceProperties(p *C.VkPhysicalDeviceProperties) (*PhysicalDeviceProperties, error) {
	uuidBytes := C.GoBytes(unsafe.Pointer(&p.pipelineCacheUUID[0]), C.VK_UUID_SIZE)
	uuid, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return nil, err
	}

	return &PhysicalDeviceProperties{
		Type: PhysicalDeviceType(p.deviceType),
		Name: C.GoString((*C.char)(&p.deviceName[0])),

		APIVersion: VKng.Version(p.apiVersion),
		DriverVersion: VKng.Version(p.driverVersion),

		VendorID: uint32(p.vendorID),
		DeviceID: uint32(p.deviceID),

		PipelineCacheUUID: uuid,
		Limits: createPhysicalDeviceLimits(&p.limits),
		SparseProperties: createSparseProperties(&p.sparseProperties),
	}, nil
}
