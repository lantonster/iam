// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/verification_code.go

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVerificationCodeRepo is a mock of VerificationCodeRepo interface.
type MockVerificationCodeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationCodeRepoMockRecorder
}

// MockVerificationCodeRepoMockRecorder is the mock recorder for MockVerificationCodeRepo.
type MockVerificationCodeRepoMockRecorder struct {
	mock *MockVerificationCodeRepo
}

// NewMockVerificationCodeRepo creates a new mock instance.
func NewMockVerificationCodeRepo(ctrl *gomock.Controller) *MockVerificationCodeRepo {
	mock := &MockVerificationCodeRepo{ctrl: ctrl}
	mock.recorder = &MockVerificationCodeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationCodeRepo) EXPECT() *MockVerificationCodeRepoMockRecorder {
	return m.recorder
}

// GenerateCode mocks base method.
func (m *MockVerificationCodeRepo) GenerateCode(c context.Context, email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCode", c, email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCode indicates an expected call of GenerateCode.
func (mr *MockVerificationCodeRepoMockRecorder) GenerateCode(c, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCode", reflect.TypeOf((*MockVerificationCodeRepo)(nil).GenerateCode), c, email)
}
