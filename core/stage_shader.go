package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type ShaderStage struct {
	Name               string
	Stage              common.ShaderStages
	Shader             ShaderModule
	SpecializationInfo map[uint32]interface{}

	common.HaveNext
}

func (s *ShaderStage) populate(allocator *cgoparam.Allocator, createInfo *C.VkPipelineShaderStageCreateInfo, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.stage = C.VkShaderStageFlagBits(s.Stage)
	createInfo.module = C.VkShaderModule(unsafe.Pointer(s.Shader.Handle()))
	createInfo.pName = (*C.char)(allocator.CString(s.Name))
	createInfo.pSpecializationInfo = nil

	if s.SpecializationInfo != nil && len(s.SpecializationInfo) > 0 {
		specInfo := (*C.VkSpecializationInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkSpecializationInfo{}))))
		mapLen := len(s.SpecializationInfo)
		specInfo.mapEntryCount = C.uint32_t(mapLen)

		mapEntryPtr := (*C.VkSpecializationMapEntry)(allocator.Malloc(mapLen * int(unsafe.Sizeof(C.VkSpecializationMapEntry{}))))
		mapEntrySlice := ([]C.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, mapLen))
		dataBytes := new(bytes.Buffer)
		mapIndex := 0

		for constantID, val := range s.SpecializationInfo {
			mapEntrySlice[mapIndex].constantID = C.uint32_t(constantID)
			mapEntrySlice[mapIndex].offset = C.uint32_t(dataBytes.Len())

			boolVal, isBool := val.(bool)
			if isBool {
				val = C.VK_FALSE
				if boolVal {
					val = C.VK_TRUE
				}
			}

			err := binary.Write(dataBytes, common.ByteOrder, val)
			if err != nil {
				return stacktrace.Propagate(err, "failed to populate shader stage with specialization values: %d -> %v", constantID, val)
			}
			mapEntrySlice[mapIndex].size = C.size_t(binary.Size(val))

			mapIndex++
		}
		specInfo.pMapEntries = mapEntryPtr
		specInfo.dataSize = C.size_t(dataBytes.Len())
		specInfo.pData = allocator.CBytes(dataBytes.Bytes())
		panic("AAA")
	}

	return nil
}

func (s *ShaderStage) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineShaderStageCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineShaderStageCreateInfo))
	err := s.populate(allocator, createInfo, next)
	return unsafe.Pointer(createInfo), err
}