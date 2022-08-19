// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/storage/persistence (interfaces: Storage)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	purpose "github.com/artefactual-sdps/enduro/internal/storage/purpose"
	source "github.com/artefactual-sdps/enduro/internal/storage/source"
	status "github.com/artefactual-sdps/enduro/internal/storage/status"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method.
func (m *MockStorage) CreateLocation(arg0 context.Context, arg1 string, arg2 *string, arg3 source.LocationSource, arg4 purpose.LocationPurpose, arg5 uuid.UUID) (*storage.StoredLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*storage.StoredLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockStorageMockRecorder) CreateLocation(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockStorage)(nil).CreateLocation), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreatePackage mocks base method.
func (m *MockStorage) CreatePackage(arg0 context.Context, arg1 string, arg2, arg3 uuid.UUID) (*storage.StoredStoragePackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*storage.StoredStoragePackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackage indicates an expected call of CreatePackage.
func (mr *MockStorageMockRecorder) CreatePackage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackage", reflect.TypeOf((*MockStorage)(nil).CreatePackage), arg0, arg1, arg2, arg3)
}

// ListLocations mocks base method.
func (m *MockStorage) ListLocations(arg0 context.Context) (storage.StoredLocationCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", arg0)
	ret0, _ := ret[0].(storage.StoredLocationCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockStorageMockRecorder) ListLocations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockStorage)(nil).ListLocations), arg0)
}

// ListPackages mocks base method.
func (m *MockStorage) ListPackages(arg0 context.Context) ([]*storage.StoredStoragePackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackages", arg0)
	ret0, _ := ret[0].([]*storage.StoredStoragePackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackages indicates an expected call of ListPackages.
func (mr *MockStorageMockRecorder) ListPackages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackages", reflect.TypeOf((*MockStorage)(nil).ListPackages), arg0)
}

// ReadLocation mocks base method.
func (m *MockStorage) ReadLocation(arg0 context.Context, arg1 uuid.UUID) (*storage.StoredLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLocation", arg0, arg1)
	ret0, _ := ret[0].(*storage.StoredLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLocation indicates an expected call of ReadLocation.
func (mr *MockStorageMockRecorder) ReadLocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLocation", reflect.TypeOf((*MockStorage)(nil).ReadLocation), arg0, arg1)
}

// ReadPackage mocks base method.
func (m *MockStorage) ReadPackage(arg0 context.Context, arg1 uuid.UUID) (*storage.StoredStoragePackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.StoredStoragePackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPackage indicates an expected call of ReadPackage.
func (mr *MockStorageMockRecorder) ReadPackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPackage", reflect.TypeOf((*MockStorage)(nil).ReadPackage), arg0, arg1)
}

// UpdatePackageLocation mocks base method.
func (m *MockStorage) UpdatePackageLocation(arg0 context.Context, arg1 string, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageLocation indicates an expected call of UpdatePackageLocation.
func (mr *MockStorageMockRecorder) UpdatePackageLocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageLocation", reflect.TypeOf((*MockStorage)(nil).UpdatePackageLocation), arg0, arg1, arg2)
}

// UpdatePackageStatus mocks base method.
func (m *MockStorage) UpdatePackageStatus(arg0 context.Context, arg1 status.PackageStatus, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageStatus indicates an expected call of UpdatePackageStatus.
func (mr *MockStorageMockRecorder) UpdatePackageStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageStatus", reflect.TypeOf((*MockStorage)(nil).UpdatePackageStatus), arg0, arg1, arg2)
}