// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/selflow/selflow/pkg/workflow (interfaces: Step)

// Package mock_workflow is a generated GoMock package.
package mock_workflow

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	workflow "github.com/selflow/selflow/pkg/workflow"
)

// MockStep is a mock of Step interface.
type MockStep struct {
	ctrl     *gomock.Controller
	recorder *MockStepMockRecorder
}

// MockStepMockRecorder is the mock recorder for MockStep.
type MockStepMockRecorder struct {
	mock *MockStep
}

// NewMockStep creates a new mock instance.
func NewMockStep(ctrl *gomock.Controller) *MockStep {
	mock := &MockStep{ctrl: ctrl}
	mock.recorder = &MockStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStep) EXPECT() *MockStepMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockStep) Cancel() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockStepMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockStep)(nil).Cancel))
}

// Execute mocks base method.
func (m *MockStep) Execute(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockStepMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStep)(nil).Execute), arg0)
}

// GetId mocks base method.
func (m *MockStep) GetId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockStepMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockStep)(nil).GetId))
}

// GetOutput mocks base method.
func (m *MockStep) GetOutput() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutput")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetOutput indicates an expected call of GetOutput.
func (mr *MockStepMockRecorder) GetOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutput", reflect.TypeOf((*MockStep)(nil).GetOutput))
}

// GetStatus mocks base method.
func (m *MockStep) GetStatus() workflow.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(workflow.Status)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockStepMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockStep)(nil).GetStatus))
}
