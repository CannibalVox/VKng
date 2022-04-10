package ext_debug_utils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"image/color"
	"unsafe"
)

type LabelOptions struct {
	Name  string
	Color color.Color

	common.HaveNext
}

func (l LabelOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsLabelEXT)
	}

	label := (*C.VkDebugUtilsLabelEXT)(preallocatedPointer)
	label.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
	label.pNext = next
	label.pLabelName = (*C.char)(allocator.CString(l.Name))

	r, g, b, a := l.Color.RGBA()
	label.color[0] = C.float(float32(r) / 65535.0)
	label.color[1] = C.float(float32(g) / 65535.0)
	label.color[2] = C.float(float32(b) / 65535.0)
	label.color[3] = C.float(float32(a) / 65535.0)

	return preallocatedPointer, nil
}

func (l LabelOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (unsafe.Pointer, error) {
	label := (*C.VkDebugUtilsLabelEXT)(cDataPointer)

	return label.pNext, nil
}

func (l *LabelOptions) PopulateFromCPointer(cDataPointer unsafe.Pointer) {
	label := (*C.VkDebugUtilsLabelEXT)(cDataPointer)
	l.Name = ""

	if label.pLabelName != nil {
		l.Name = C.GoString(label.pLabelName)
	}

	r := uint8(float32(label.color[0])*65535.0 + 0.001)
	g := uint8(float32(label.color[1])*65535.0 + 0.001)
	b := uint8(float32(label.color[2])*65535.0 + 0.001)
	a := uint8(float32(label.color[3])*65535.0 + 0.001)

	l.Color = color.RGBA{R: r, G: g, B: b, A: a}
}
