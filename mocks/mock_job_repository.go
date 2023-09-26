// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repositories/job_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "coverCraft/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJobRepository is a mock of JobRepository interface.
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository.
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance.
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockJobRepository) Create(job entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockJobRepositoryMockRecorder) Create(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockJobRepository)(nil).Create), job)
}

// ListAllJobs mocks base method.
func (m *MockJobRepository) ListAllJobs() ([]entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllJobs")
	ret0, _ := ret[0].([]entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllJobs indicates an expected call of ListAllJobs.
func (mr *MockJobRepositoryMockRecorder) ListAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllJobs", reflect.TypeOf((*MockJobRepository)(nil).ListAllJobs))
}