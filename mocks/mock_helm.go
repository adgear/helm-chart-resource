// Code generated by MockGen. DO NOT EDIT.
// Source: helm/helm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	utils "github.com/adgear/helm-chart-resource/utils"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHelm is a mock of Helm interface
type MockHelm struct {
	ctrl     *gomock.Controller
	recorder *MockHelmMockRecorder
}

// MockHelmMockRecorder is the mock recorder for MockHelm
type MockHelmMockRecorder struct {
	mock *MockHelm
}

// NewMockHelm creates a new mock instance
func NewMockHelm(ctrl *gomock.Controller) *MockHelm {
	mock := &MockHelm{ctrl: ctrl}
	mock.recorder = &MockHelmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelm) EXPECT() *MockHelmMockRecorder {
	return m.recorder
}

// DepUpdate mocks base method
func (m *MockHelm) DepUpdate(path string) (string, error) {
	ret := m.ctrl.Call(m, "DepUpdate", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DepUpdate indicates an expected call of DepUpdate
func (mr *MockHelmMockRecorder) DepUpdate(path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepUpdate", reflect.TypeOf((*MockHelm)(nil).DepUpdate), path)
}

// RepoUpdate mocks base method
func (m *MockHelm) RepoUpdate() {
	m.ctrl.Call(m, "RepoUpdate")
}

// RepoUpdate indicates an expected call of RepoUpdate
func (mr *MockHelmMockRecorder) RepoUpdate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoUpdate", reflect.TypeOf((*MockHelm)(nil).RepoUpdate))
}

// Search mocks base method
func (m *MockHelm) Search(repo string) (string, error) {
	ret := m.ctrl.Call(m, "Search", repo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockHelmMockRecorder) Search(repo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockHelm)(nil).Search), repo)
}

// InstallHelmRepo mocks base method
func (m *MockHelm) InstallHelmRepo(repos []utils.Repo) error {
	ret := m.ctrl.Call(m, "InstallHelmRepo", repos)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallHelmRepo indicates an expected call of InstallHelmRepo
func (mr *MockHelmMockRecorder) InstallHelmRepo(repos interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallHelmRepo", reflect.TypeOf((*MockHelm)(nil).InstallHelmRepo), repos)
}

// BuildHelmChart mocks base method
func (m *MockHelm) BuildHelmChart(source, path string) error {
	ret := m.ctrl.Call(m, "BuildHelmChart", source, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildHelmChart indicates an expected call of BuildHelmChart
func (mr *MockHelmMockRecorder) BuildHelmChart(source, path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildHelmChart", reflect.TypeOf((*MockHelm)(nil).BuildHelmChart), source, path)
}

// PackageHelmChart mocks base method
func (m *MockHelm) PackageHelmChart(source, path, tmpdir string) error {
	ret := m.ctrl.Call(m, "PackageHelmChart", source, path, tmpdir)
	ret0, _ := ret[0].(error)
	return ret0
}

// PackageHelmChart indicates an expected call of PackageHelmChart
func (mr *MockHelmMockRecorder) PackageHelmChart(source, path, tmpdir interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackageHelmChart", reflect.TypeOf((*MockHelm)(nil).PackageHelmChart), source, path, tmpdir)
}

// ExtractChartVersion mocks base method
func (m *MockHelm) ExtractChartVersion(source, path string) (string, error) {
	ret := m.ctrl.Call(m, "ExtractChartVersion", source, path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractChartVersion indicates an expected call of ExtractChartVersion
func (mr *MockHelmMockRecorder) ExtractChartVersion(source, path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractChartVersion", reflect.TypeOf((*MockHelm)(nil).ExtractChartVersion), source, path)
}
