// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go
//
// Generated by this command:
//
//	mockgen -source=manager.go -destination=mocks/mock.go -package=mocktransactor
//

// Package mocktransactor is a generated GoMock package.
package mocktransactor

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// WithinTransaction mocks base method.
func (m *MockManager) WithinTransaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithinTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithinTransaction indicates an expected call of WithinTransaction.
func (mr *MockManagerMockRecorder) WithinTransaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithinTransaction", reflect.TypeOf((*MockManager)(nil).WithinTransaction), ctx, fn)
}
