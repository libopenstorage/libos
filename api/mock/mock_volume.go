// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageVolumeServer,OpenStorageVolumeClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockOpenStorageVolumeServer is a mock of OpenStorageVolumeServer interface
type MockOpenStorageVolumeServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageVolumeServerMockRecorder
}

// MockOpenStorageVolumeServerMockRecorder is the mock recorder for MockOpenStorageVolumeServer
type MockOpenStorageVolumeServerMockRecorder struct {
	mock *MockOpenStorageVolumeServer
}

// NewMockOpenStorageVolumeServer creates a new mock instance
func NewMockOpenStorageVolumeServer(ctrl *gomock.Controller) *MockOpenStorageVolumeServer {
	mock := &MockOpenStorageVolumeServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageVolumeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageVolumeServer) EXPECT() *MockOpenStorageVolumeServerMockRecorder {
	return m.recorder
}

// CapacityUsage mocks base method
func (m *MockOpenStorageVolumeServer) CapacityUsage(arg0 context.Context, arg1 *api.SdkVolumeCapacityUsageRequest) (*api.SdkVolumeCapacityUsageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityUsage", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeCapacityUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CapacityUsage indicates an expected call of CapacityUsage
func (mr *MockOpenStorageVolumeServerMockRecorder) CapacityUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityUsage", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).CapacityUsage), arg0, arg1)
}

// Clone mocks base method
func (m *MockOpenStorageVolumeServer) Clone(arg0 context.Context, arg1 *api.SdkVolumeCloneRequest) (*api.SdkVolumeCloneResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeCloneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clone indicates an expected call of Clone
func (mr *MockOpenStorageVolumeServerMockRecorder) Clone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Clone), arg0, arg1)
}

// Create mocks base method
func (m *MockOpenStorageVolumeServer) Create(arg0 context.Context, arg1 *api.SdkVolumeCreateRequest) (*api.SdkVolumeCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOpenStorageVolumeServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockOpenStorageVolumeServer) Delete(arg0 context.Context, arg1 *api.SdkVolumeDeleteRequest) (*api.SdkVolumeDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockOpenStorageVolumeServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Delete), arg0, arg1)
}

// Enumerate mocks base method
func (m *MockOpenStorageVolumeServer) Enumerate(arg0 context.Context, arg1 *api.SdkVolumeEnumerateRequest) (*api.SdkVolumeEnumerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enumerate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockOpenStorageVolumeServerMockRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Enumerate), arg0, arg1)
}

// EnumerateWithFilters mocks base method
func (m *MockOpenStorageVolumeServer) EnumerateWithFilters(arg0 context.Context, arg1 *api.SdkVolumeEnumerateWithFiltersRequest) (*api.SdkVolumeEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateWithFilters", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateWithFilters indicates an expected call of EnumerateWithFilters
func (mr *MockOpenStorageVolumeServerMockRecorder) EnumerateWithFilters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateWithFilters", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).EnumerateWithFilters), arg0, arg1)
}

// Inspect mocks base method
func (m *MockOpenStorageVolumeServer) Inspect(arg0 context.Context, arg1 *api.SdkVolumeInspectRequest) (*api.SdkVolumeInspectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeInspectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockOpenStorageVolumeServerMockRecorder) Inspect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Inspect), arg0, arg1)
}

// InspectWithFilters mocks base method
func (m *MockOpenStorageVolumeServer) InspectWithFilters(arg0 context.Context, arg1 *api.SdkVolumeInspectWithFiltersRequest) (*api.SdkVolumeInspectWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectWithFilters", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeInspectWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectWithFilters indicates an expected call of InspectWithFilters
func (mr *MockOpenStorageVolumeServerMockRecorder) InspectWithFilters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectWithFilters", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).InspectWithFilters), arg0, arg1)
}

// MigrationCancel mocks base method
func (m *MockOpenStorageVolumeServer) MigrationCancel(arg0 context.Context, arg1 *api.SdkVolumeMigrationCancelRequest) (*api.SdkVolumeMigrationCancelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationCancel", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationCancelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationCancel indicates an expected call of MigrationCancel
func (mr *MockOpenStorageVolumeServerMockRecorder) MigrationCancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationCancel", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).MigrationCancel), arg0, arg1)
}

// MigrationComplete mocks base method
func (m *MockOpenStorageVolumeServer) MigrationComplete(arg0 context.Context, arg1 *api.SdkVolumeMigrationCompleteRequest) (*api.SdkVolumeMigrationCompleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationComplete", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationCompleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationComplete indicates an expected call of MigrationComplete
func (mr *MockOpenStorageVolumeServerMockRecorder) MigrationComplete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationComplete", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).MigrationComplete), arg0, arg1)
}

// MigrationFailover mocks base method
func (m *MockOpenStorageVolumeServer) MigrationFailover(arg0 context.Context, arg1 *api.SdkVolumeMigrationFailoverRequest) (*api.SdkVolumeMigrationFailoverResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationFailover", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationFailoverResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationFailover indicates an expected call of MigrationFailover
func (mr *MockOpenStorageVolumeServerMockRecorder) MigrationFailover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationFailover", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).MigrationFailover), arg0, arg1)
}

// MigrationStart mocks base method
func (m *MockOpenStorageVolumeServer) MigrationStart(arg0 context.Context, arg1 *api.SdkVolumeMigrationStartRequest) (*api.SdkVolumeMigrationStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationStart", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationStart indicates an expected call of MigrationStart
func (mr *MockOpenStorageVolumeServerMockRecorder) MigrationStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationStart", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).MigrationStart), arg0, arg1)
}

// SnapshotCreate mocks base method
func (m *MockOpenStorageVolumeServer) SnapshotCreate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotCreateRequest) (*api.SdkVolumeSnapshotCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotCreate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotCreate indicates an expected call of SnapshotCreate
func (mr *MockOpenStorageVolumeServerMockRecorder) SnapshotCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotCreate", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).SnapshotCreate), arg0, arg1)
}

// SnapshotEnumerate mocks base method
func (m *MockOpenStorageVolumeServer) SnapshotEnumerate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotEnumerateRequest) (*api.SdkVolumeSnapshotEnumerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotEnumerate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotEnumerate indicates an expected call of SnapshotEnumerate
func (mr *MockOpenStorageVolumeServerMockRecorder) SnapshotEnumerate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotEnumerate", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).SnapshotEnumerate), arg0, arg1)
}

// SnapshotEnumerateWithFilters mocks base method
func (m *MockOpenStorageVolumeServer) SnapshotEnumerateWithFilters(arg0 context.Context, arg1 *api.SdkVolumeSnapshotEnumerateWithFiltersRequest) (*api.SdkVolumeSnapshotEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotEnumerateWithFilters", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotEnumerateWithFilters indicates an expected call of SnapshotEnumerateWithFilters
func (mr *MockOpenStorageVolumeServerMockRecorder) SnapshotEnumerateWithFilters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotEnumerateWithFilters", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).SnapshotEnumerateWithFilters), arg0, arg1)
}

// SnapshotRestore mocks base method
func (m *MockOpenStorageVolumeServer) SnapshotRestore(arg0 context.Context, arg1 *api.SdkVolumeSnapshotRestoreRequest) (*api.SdkVolumeSnapshotRestoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotRestore", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotRestoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotRestore indicates an expected call of SnapshotRestore
func (mr *MockOpenStorageVolumeServerMockRecorder) SnapshotRestore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRestore", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).SnapshotRestore), arg0, arg1)
}

// SnapshotScheduleUpdate mocks base method
func (m *MockOpenStorageVolumeServer) SnapshotScheduleUpdate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotScheduleUpdateRequest) (*api.SdkVolumeSnapshotScheduleUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotScheduleUpdate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotScheduleUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotScheduleUpdate indicates an expected call of SnapshotScheduleUpdate
func (mr *MockOpenStorageVolumeServerMockRecorder) SnapshotScheduleUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotScheduleUpdate", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).SnapshotScheduleUpdate), arg0, arg1)
}

// Stats mocks base method
func (m *MockOpenStorageVolumeServer) Stats(arg0 context.Context, arg1 *api.SdkVolumeStatsRequest) (*api.SdkVolumeStatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockOpenStorageVolumeServerMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Stats), arg0, arg1)
}

// Update mocks base method
func (m *MockOpenStorageVolumeServer) Update(arg0 context.Context, arg1 *api.SdkVolumeUpdateRequest) (*api.SdkVolumeUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockOpenStorageVolumeServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).Update), arg0, arg1)
}

// VolumeCatalog mocks base method
func (m *MockOpenStorageVolumeServer) VolumeCatalog(arg0 context.Context, arg1 *api.SdkVolumeCatalogRequest) (*api.SdkVolumeCatalogResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeCatalog", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkVolumeCatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeCatalog indicates an expected call of VolumeCatalog
func (mr *MockOpenStorageVolumeServerMockRecorder) VolumeCatalog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeCatalog", reflect.TypeOf((*MockOpenStorageVolumeServer)(nil).VolumeCatalog), arg0, arg1)
}

// MockOpenStorageVolumeClient is a mock of OpenStorageVolumeClient interface
type MockOpenStorageVolumeClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageVolumeClientMockRecorder
}

// MockOpenStorageVolumeClientMockRecorder is the mock recorder for MockOpenStorageVolumeClient
type MockOpenStorageVolumeClientMockRecorder struct {
	mock *MockOpenStorageVolumeClient
}

// NewMockOpenStorageVolumeClient creates a new mock instance
func NewMockOpenStorageVolumeClient(ctrl *gomock.Controller) *MockOpenStorageVolumeClient {
	mock := &MockOpenStorageVolumeClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageVolumeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenStorageVolumeClient) EXPECT() *MockOpenStorageVolumeClientMockRecorder {
	return m.recorder
}

// CapacityUsage mocks base method
func (m *MockOpenStorageVolumeClient) CapacityUsage(arg0 context.Context, arg1 *api.SdkVolumeCapacityUsageRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeCapacityUsageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CapacityUsage", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeCapacityUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CapacityUsage indicates an expected call of CapacityUsage
func (mr *MockOpenStorageVolumeClientMockRecorder) CapacityUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityUsage", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).CapacityUsage), varargs...)
}

// Clone mocks base method
func (m *MockOpenStorageVolumeClient) Clone(arg0 context.Context, arg1 *api.SdkVolumeCloneRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeCloneResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Clone", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeCloneResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clone indicates an expected call of Clone
func (mr *MockOpenStorageVolumeClientMockRecorder) Clone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Clone), varargs...)
}

// Create mocks base method
func (m *MockOpenStorageVolumeClient) Create(arg0 context.Context, arg1 *api.SdkVolumeCreateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOpenStorageVolumeClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockOpenStorageVolumeClient) Delete(arg0 context.Context, arg1 *api.SdkVolumeDeleteRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockOpenStorageVolumeClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Delete), varargs...)
}

// Enumerate mocks base method
func (m *MockOpenStorageVolumeClient) Enumerate(arg0 context.Context, arg1 *api.SdkVolumeEnumerateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeEnumerateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Enumerate", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockOpenStorageVolumeClientMockRecorder) Enumerate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Enumerate), varargs...)
}

// EnumerateWithFilters mocks base method
func (m *MockOpenStorageVolumeClient) EnumerateWithFilters(arg0 context.Context, arg1 *api.SdkVolumeEnumerateWithFiltersRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnumerateWithFilters", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateWithFilters indicates an expected call of EnumerateWithFilters
func (mr *MockOpenStorageVolumeClientMockRecorder) EnumerateWithFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateWithFilters", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).EnumerateWithFilters), varargs...)
}

// Inspect mocks base method
func (m *MockOpenStorageVolumeClient) Inspect(arg0 context.Context, arg1 *api.SdkVolumeInspectRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeInspectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Inspect", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeInspectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockOpenStorageVolumeClientMockRecorder) Inspect(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Inspect), varargs...)
}

// InspectWithFilters mocks base method
func (m *MockOpenStorageVolumeClient) InspectWithFilters(arg0 context.Context, arg1 *api.SdkVolumeInspectWithFiltersRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeInspectWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InspectWithFilters", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeInspectWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectWithFilters indicates an expected call of InspectWithFilters
func (mr *MockOpenStorageVolumeClientMockRecorder) InspectWithFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectWithFilters", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).InspectWithFilters), varargs...)
}

// MigrationCancel mocks base method
func (m *MockOpenStorageVolumeClient) MigrationCancel(arg0 context.Context, arg1 *api.SdkVolumeMigrationCancelRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeMigrationCancelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MigrationCancel", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationCancelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationCancel indicates an expected call of MigrationCancel
func (mr *MockOpenStorageVolumeClientMockRecorder) MigrationCancel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationCancel", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).MigrationCancel), varargs...)
}

// MigrationComplete mocks base method
func (m *MockOpenStorageVolumeClient) MigrationComplete(arg0 context.Context, arg1 *api.SdkVolumeMigrationCompleteRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeMigrationCompleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MigrationComplete", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationCompleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationComplete indicates an expected call of MigrationComplete
func (mr *MockOpenStorageVolumeClientMockRecorder) MigrationComplete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationComplete", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).MigrationComplete), varargs...)
}

// MigrationFailover mocks base method
func (m *MockOpenStorageVolumeClient) MigrationFailover(arg0 context.Context, arg1 *api.SdkVolumeMigrationFailoverRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeMigrationFailoverResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MigrationFailover", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationFailoverResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationFailover indicates an expected call of MigrationFailover
func (mr *MockOpenStorageVolumeClientMockRecorder) MigrationFailover(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationFailover", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).MigrationFailover), varargs...)
}

// MigrationStart mocks base method
func (m *MockOpenStorageVolumeClient) MigrationStart(arg0 context.Context, arg1 *api.SdkVolumeMigrationStartRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeMigrationStartResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MigrationStart", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeMigrationStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrationStart indicates an expected call of MigrationStart
func (mr *MockOpenStorageVolumeClientMockRecorder) MigrationStart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationStart", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).MigrationStart), varargs...)
}

// SnapshotCreate mocks base method
func (m *MockOpenStorageVolumeClient) SnapshotCreate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotCreateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeSnapshotCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnapshotCreate", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotCreate indicates an expected call of SnapshotCreate
func (mr *MockOpenStorageVolumeClientMockRecorder) SnapshotCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotCreate", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).SnapshotCreate), varargs...)
}

// SnapshotEnumerate mocks base method
func (m *MockOpenStorageVolumeClient) SnapshotEnumerate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotEnumerateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeSnapshotEnumerateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnapshotEnumerate", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotEnumerate indicates an expected call of SnapshotEnumerate
func (mr *MockOpenStorageVolumeClientMockRecorder) SnapshotEnumerate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotEnumerate", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).SnapshotEnumerate), varargs...)
}

// SnapshotEnumerateWithFilters mocks base method
func (m *MockOpenStorageVolumeClient) SnapshotEnumerateWithFilters(arg0 context.Context, arg1 *api.SdkVolumeSnapshotEnumerateWithFiltersRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeSnapshotEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnapshotEnumerateWithFilters", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotEnumerateWithFilters indicates an expected call of SnapshotEnumerateWithFilters
func (mr *MockOpenStorageVolumeClientMockRecorder) SnapshotEnumerateWithFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotEnumerateWithFilters", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).SnapshotEnumerateWithFilters), varargs...)
}

// SnapshotRestore mocks base method
func (m *MockOpenStorageVolumeClient) SnapshotRestore(arg0 context.Context, arg1 *api.SdkVolumeSnapshotRestoreRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeSnapshotRestoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnapshotRestore", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotRestoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotRestore indicates an expected call of SnapshotRestore
func (mr *MockOpenStorageVolumeClientMockRecorder) SnapshotRestore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRestore", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).SnapshotRestore), varargs...)
}

// SnapshotScheduleUpdate mocks base method
func (m *MockOpenStorageVolumeClient) SnapshotScheduleUpdate(arg0 context.Context, arg1 *api.SdkVolumeSnapshotScheduleUpdateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeSnapshotScheduleUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnapshotScheduleUpdate", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeSnapshotScheduleUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotScheduleUpdate indicates an expected call of SnapshotScheduleUpdate
func (mr *MockOpenStorageVolumeClientMockRecorder) SnapshotScheduleUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotScheduleUpdate", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).SnapshotScheduleUpdate), varargs...)
}

// Stats mocks base method
func (m *MockOpenStorageVolumeClient) Stats(arg0 context.Context, arg1 *api.SdkVolumeStatsRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeStatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stats", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockOpenStorageVolumeClientMockRecorder) Stats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Stats), varargs...)
}

// Update mocks base method
func (m *MockOpenStorageVolumeClient) Update(arg0 context.Context, arg1 *api.SdkVolumeUpdateRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockOpenStorageVolumeClientMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).Update), varargs...)
}

// VolumeCatalog mocks base method
func (m *MockOpenStorageVolumeClient) VolumeCatalog(arg0 context.Context, arg1 *api.SdkVolumeCatalogRequest, arg2 ...grpc.CallOption) (*api.SdkVolumeCatalogResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VolumeCatalog", varargs...)
	ret0, _ := ret[0].(*api.SdkVolumeCatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeCatalog indicates an expected call of VolumeCatalog
func (mr *MockOpenStorageVolumeClientMockRecorder) VolumeCatalog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeCatalog", reflect.TypeOf((*MockOpenStorageVolumeClient)(nil).VolumeCatalog), varargs...)
}
