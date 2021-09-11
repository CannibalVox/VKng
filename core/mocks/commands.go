// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/CannibalVox/VKng/core"
	commands "github.com/CannibalVox/VKng/core/commands"
	loader "github.com/CannibalVox/VKng/core/loader"
	pipeline "github.com/CannibalVox/VKng/core/pipeline"
	resource "github.com/CannibalVox/VKng/core/resources"
	cgoalloc "github.com/CannibalVox/cgoalloc"
	gomock "github.com/golang/mock/gomock"
)

// MockCommandBuffer is a mock of CommandBuffer interface.
type MockCommandBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockCommandBufferMockRecorder
}

// MockCommandBufferMockRecorder is the mock recorder for MockCommandBuffer.
type MockCommandBufferMockRecorder struct {
	mock *MockCommandBuffer
}

// NewMockCommandBuffer creates a new mock instance.
func NewMockCommandBuffer(ctrl *gomock.Controller) *MockCommandBuffer {
	mock := &MockCommandBuffer{ctrl: ctrl}
	mock.recorder = &MockCommandBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandBuffer) EXPECT() *MockCommandBufferMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockCommandBuffer) Begin(allocator cgoalloc.Allocator, o *commands.BeginOptions) (loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", allocator, o)
	ret0, _ := ret[0].(loader.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockCommandBufferMockRecorder) Begin(allocator, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockCommandBuffer)(nil).Begin), allocator, o)
}

// CmdBeginRenderPass mocks base method.
func (m *MockCommandBuffer) CmdBeginRenderPass(allocator cgoalloc.Allocator, contents commands.SubpassContents, o *commands.RenderPassBeginOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdBeginRenderPass", allocator, contents, o)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdBeginRenderPass indicates an expected call of CmdBeginRenderPass.
func (mr *MockCommandBufferMockRecorder) CmdBeginRenderPass(allocator, contents, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdBeginRenderPass", reflect.TypeOf((*MockCommandBuffer)(nil).CmdBeginRenderPass), allocator, contents, o)
}

// CmdBindIndexBuffer mocks base method.
func (m *MockCommandBuffer) CmdBindIndexBuffer(buffer resource.Buffer, offset int, indexType core.IndexType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdBindIndexBuffer", buffer, offset, indexType)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdBindIndexBuffer indicates an expected call of CmdBindIndexBuffer.
func (mr *MockCommandBufferMockRecorder) CmdBindIndexBuffer(buffer, offset, indexType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdBindIndexBuffer", reflect.TypeOf((*MockCommandBuffer)(nil).CmdBindIndexBuffer), buffer, offset, indexType)
}

// CmdBindPipeline mocks base method.
func (m *MockCommandBuffer) CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline pipeline.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdBindPipeline", bindPoint, pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdBindPipeline indicates an expected call of CmdBindPipeline.
func (mr *MockCommandBufferMockRecorder) CmdBindPipeline(bindPoint, pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdBindPipeline", reflect.TypeOf((*MockCommandBuffer)(nil).CmdBindPipeline), bindPoint, pipeline)
}

// CmdBindVertexBuffers mocks base method.
func (m *MockCommandBuffer) CmdBindVertexBuffers(allocator cgoalloc.Allocator, firstBinding uint32, buffers []resource.Buffer, bufferOffsets []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdBindVertexBuffers", allocator, firstBinding, buffers, bufferOffsets)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdBindVertexBuffers indicates an expected call of CmdBindVertexBuffers.
func (mr *MockCommandBufferMockRecorder) CmdBindVertexBuffers(allocator, firstBinding, buffers, bufferOffsets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdBindVertexBuffers", reflect.TypeOf((*MockCommandBuffer)(nil).CmdBindVertexBuffers), allocator, firstBinding, buffers, bufferOffsets)
}

// CmdCopyBuffer mocks base method.
func (m *MockCommandBuffer) CmdCopyBuffer(allocator cgoalloc.Allocator, srcBuffer, dstBuffer resource.Buffer, copyRegions []commands.BufferCopy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdCopyBuffer", allocator, srcBuffer, dstBuffer, copyRegions)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdCopyBuffer indicates an expected call of CmdCopyBuffer.
func (mr *MockCommandBufferMockRecorder) CmdCopyBuffer(allocator, srcBuffer, dstBuffer, copyRegions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdCopyBuffer", reflect.TypeOf((*MockCommandBuffer)(nil).CmdCopyBuffer), allocator, srcBuffer, dstBuffer, copyRegions)
}

// CmdDraw mocks base method.
func (m *MockCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdDraw", vertexCount, instanceCount, firstVertex, firstInstance)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdDraw indicates an expected call of CmdDraw.
func (mr *MockCommandBufferMockRecorder) CmdDraw(vertexCount, instanceCount, firstVertex, firstInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdDraw", reflect.TypeOf((*MockCommandBuffer)(nil).CmdDraw), vertexCount, instanceCount, firstVertex, firstInstance)
}

// CmdDrawIndexed mocks base method.
func (m *MockCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdDrawIndexed", indexCount, instanceCount, firstIndex, vertexOffset, firstInstance)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdDrawIndexed indicates an expected call of CmdDrawIndexed.
func (mr *MockCommandBufferMockRecorder) CmdDrawIndexed(indexCount, instanceCount, firstIndex, vertexOffset, firstInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdDrawIndexed", reflect.TypeOf((*MockCommandBuffer)(nil).CmdDrawIndexed), indexCount, instanceCount, firstIndex, vertexOffset, firstInstance)
}

// CmdEndRenderPass mocks base method.
func (m *MockCommandBuffer) CmdEndRenderPass() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdEndRenderPass")
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdEndRenderPass indicates an expected call of CmdEndRenderPass.
func (mr *MockCommandBufferMockRecorder) CmdEndRenderPass() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdEndRenderPass", reflect.TypeOf((*MockCommandBuffer)(nil).CmdEndRenderPass))
}

// Destroy mocks base method.
func (m *MockCommandBuffer) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockCommandBufferMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockCommandBuffer)(nil).Destroy))
}

// End mocks base method.
func (m *MockCommandBuffer) End() (loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "End")
	ret0, _ := ret[0].(loader.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// End indicates an expected call of End.
func (mr *MockCommandBufferMockRecorder) End() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "End", reflect.TypeOf((*MockCommandBuffer)(nil).End))
}

// Handle mocks base method.
func (m *MockCommandBuffer) Handle() loader.VkCommandBuffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(loader.VkCommandBuffer)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockCommandBufferMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockCommandBuffer)(nil).Handle))
}

// MockCommandPool is a mock of CommandPool interface.
type MockCommandPool struct {
	ctrl     *gomock.Controller
	recorder *MockCommandPoolMockRecorder
}

// MockCommandPoolMockRecorder is the mock recorder for MockCommandPool.
type MockCommandPoolMockRecorder struct {
	mock *MockCommandPool
}

// NewMockCommandPool creates a new mock instance.
func NewMockCommandPool(ctrl *gomock.Controller) *MockCommandPool {
	mock := &MockCommandPool{ctrl: ctrl}
	mock.recorder = &MockCommandPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandPool) EXPECT() *MockCommandPoolMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockCommandPool) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockCommandPoolMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockCommandPool)(nil).Destroy))
}

// DestroyBuffers mocks base method.
func (m *MockCommandPool) DestroyBuffers(allocator cgoalloc.Allocator, buffers []commands.CommandBuffer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyBuffers", allocator, buffers)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyBuffers indicates an expected call of DestroyBuffers.
func (mr *MockCommandPoolMockRecorder) DestroyBuffers(allocator, buffers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyBuffers", reflect.TypeOf((*MockCommandPool)(nil).DestroyBuffers), allocator, buffers)
}

// Handle mocks base method.
func (m *MockCommandPool) Handle() loader.VkCommandPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(loader.VkCommandPool)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockCommandPoolMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockCommandPool)(nil).Handle))
}
