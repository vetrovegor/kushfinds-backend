// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vetrovegor/kushfinds-backend/internal/auth/service (interfaces: TokenManager)
//
// Generated by this command:
//
//	mockgen -destination=mocks/token/mock.go -package=mocktoken . TokenManager
//

// Package mocktoken is a generated GoMock package.
package mocktoken

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
	isgomock struct{}
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenManager) GenerateToken(userID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenManagerMockRecorder) GenerateToken(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenManager)(nil).GenerateToken), userID)
}

// GetRefreshTokenTTL mocks base method.
func (m *MockTokenManager) GetRefreshTokenTTL() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshTokenTTL")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetRefreshTokenTTL indicates an expected call of GetRefreshTokenTTL.
func (mr *MockTokenManagerMockRecorder) GetRefreshTokenTTL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshTokenTTL", reflect.TypeOf((*MockTokenManager)(nil).GetRefreshTokenTTL))
}
