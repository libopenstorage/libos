// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/libopenstorage/openstorage/volume (interfaces: VolumeDriver)

package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
)

// Mock of VolumeDriver interface
type MockVolumeDriver struct {
	ctrl     *gomock.Controller
	recorder *_MockVolumeDriverRecorder
}

// Recorder for MockVolumeDriver (not exported)
type _MockVolumeDriverRecorder struct {
	mock *MockVolumeDriver
}

func NewMockVolumeDriver(ctrl *gomock.Controller) *MockVolumeDriver {
	mock := &MockVolumeDriver{ctrl: ctrl}
	mock.recorder = &_MockVolumeDriverRecorder{mock}
	return mock
}

func (_m *MockVolumeDriver) EXPECT() *_MockVolumeDriverRecorder {
	return _m.recorder
}

func (_m *MockVolumeDriver) Attach(_param0 string, _param1 map[string]string) (string, error) {
	ret := _m.ctrl.Call(_m, "Attach", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Attach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Attach", arg0, arg1)
}

func (_m *MockVolumeDriver) CapacityUsage(_param0 string) (*api.CapacityUsageResponse, error) {
	ret := _m.ctrl.Call(_m, "CapacityUsage", _param0)
	ret0, _ := ret[0].(*api.CapacityUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CapacityUsage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CapacityUsage", arg0)
}

func (_m *MockVolumeDriver) Catalog(_param0 string, _param1 string, _param2 string) (api.CatalogResponse, error) {
	ret := _m.ctrl.Call(_m, "Catalog", _param0, _param1, _param2)
	ret0, _ := ret[0].(api.CatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Catalog(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Catalog", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) CloudBackupCatalog(_param0 *api.CloudBackupCatalogRequest) (*api.CloudBackupCatalogResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupCatalog", _param0)
	ret0, _ := ret[0].(*api.CloudBackupCatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupCatalog(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupCatalog", arg0)
}

func (_m *MockVolumeDriver) CloudBackupCreate(_param0 *api.CloudBackupCreateRequest) (*api.CloudBackupCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupCreate", _param0)
	ret0, _ := ret[0].(*api.CloudBackupCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupCreate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupCreate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupDelete(_param0 *api.CloudBackupDeleteRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupDelete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupDelete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupDelete", arg0)
}

func (_m *MockVolumeDriver) CloudBackupDeleteAll(_param0 *api.CloudBackupDeleteAllRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupDeleteAll", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupDeleteAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupDeleteAll", arg0)
}

func (_m *MockVolumeDriver) CloudBackupEnumerate(_param0 *api.CloudBackupEnumerateRequest) (*api.CloudBackupEnumerateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupEnumerate", _param0)
	ret0, _ := ret[0].(*api.CloudBackupEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupEnumerate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupEnumerate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupGroupCreate(_param0 *api.CloudBackupGroupCreateRequest) (*api.CloudBackupGroupCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupGroupCreate", _param0)
	ret0, _ := ret[0].(*api.CloudBackupGroupCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupGroupCreate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupGroupCreate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupGroupSchedCreate(_param0 *api.CloudBackupGroupSchedCreateRequest) (*api.CloudBackupSchedCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupGroupSchedCreate", _param0)
	ret0, _ := ret[0].(*api.CloudBackupSchedCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupGroupSchedCreate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupGroupSchedCreate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupGroupSchedUpdate(_param0 *api.CloudBackupGroupSchedUpdateRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupGroupSchedUpdate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupGroupSchedUpdate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupGroupSchedUpdate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupHistory(_param0 *api.CloudBackupHistoryRequest) (*api.CloudBackupHistoryResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupHistory", _param0)
	ret0, _ := ret[0].(*api.CloudBackupHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupHistory(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupHistory", arg0)
}

func (_m *MockVolumeDriver) CloudBackupRestore(_param0 *api.CloudBackupRestoreRequest) (*api.CloudBackupRestoreResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupRestore", _param0)
	ret0, _ := ret[0].(*api.CloudBackupRestoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupRestore(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupRestore", arg0)
}

func (_m *MockVolumeDriver) CloudBackupSchedCreate(_param0 *api.CloudBackupSchedCreateRequest) (*api.CloudBackupSchedCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupSchedCreate", _param0)
	ret0, _ := ret[0].(*api.CloudBackupSchedCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupSchedCreate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupSchedCreate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupSchedDelete(_param0 *api.CloudBackupSchedDeleteRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupSchedDelete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupSchedDelete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupSchedDelete", arg0)
}

func (_m *MockVolumeDriver) CloudBackupSchedEnumerate() (*api.CloudBackupSchedEnumerateResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupSchedEnumerate")
	ret0, _ := ret[0].(*api.CloudBackupSchedEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupSchedEnumerate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupSchedEnumerate")
}

func (_m *MockVolumeDriver) CloudBackupSchedUpdate(_param0 *api.CloudBackupSchedUpdateRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupSchedUpdate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupSchedUpdate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupSchedUpdate", arg0)
}

func (_m *MockVolumeDriver) CloudBackupSize(_param0 *api.SdkCloudBackupSizeRequest) (*api.SdkCloudBackupSizeResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupSize", _param0)
	ret0, _ := ret[0].(*api.SdkCloudBackupSizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupSize", arg0)
}

func (_m *MockVolumeDriver) CloudBackupStateChange(_param0 *api.CloudBackupStateChangeRequest) error {
	ret := _m.ctrl.Call(_m, "CloudBackupStateChange", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupStateChange(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupStateChange", arg0)
}

func (_m *MockVolumeDriver) CloudBackupStatus(_param0 *api.CloudBackupStatusRequest) (*api.CloudBackupStatusResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudBackupStatus", _param0)
	ret0, _ := ret[0].(*api.CloudBackupStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudBackupStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudBackupStatus", arg0)
}

func (_m *MockVolumeDriver) CloudMigrateCancel(_param0 *api.CloudMigrateCancelRequest) error {
	ret := _m.ctrl.Call(_m, "CloudMigrateCancel", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CloudMigrateCancel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudMigrateCancel", arg0)
}

func (_m *MockVolumeDriver) CloudMigrateStart(_param0 *api.CloudMigrateStartRequest) (*api.CloudMigrateStartResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudMigrateStart", _param0)
	ret0, _ := ret[0].(*api.CloudMigrateStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudMigrateStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudMigrateStart", arg0)
}

func (_m *MockVolumeDriver) CloudMigrateStatus(_param0 *api.CloudMigrateStatusRequest) (*api.CloudMigrateStatusResponse, error) {
	ret := _m.ctrl.Call(_m, "CloudMigrateStatus", _param0)
	ret0, _ := ret[0].(*api.CloudMigrateStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CloudMigrateStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloudMigrateStatus", arg0)
}

func (_m *MockVolumeDriver) Create(_param0 *api.VolumeLocator, _param1 *api.Source, _param2 *api.VolumeSpec) (string, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1, _param2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) CredsCreate(_param0 map[string]string) (string, error) {
	ret := _m.ctrl.Call(_m, "CredsCreate", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CredsCreate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CredsCreate", arg0)
}

func (_m *MockVolumeDriver) CredsDelete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CredsDelete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CredsDelete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CredsDelete", arg0)
}

func (_m *MockVolumeDriver) CredsEnumerate() (map[string]interface{}, error) {
	ret := _m.ctrl.Call(_m, "CredsEnumerate")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) CredsEnumerate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CredsEnumerate")
}

func (_m *MockVolumeDriver) CredsValidate(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CredsValidate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) CredsValidate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CredsValidate", arg0)
}

func (_m *MockVolumeDriver) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockVolumeDriver) Detach(_param0 string, _param1 map[string]string) error {
	ret := _m.ctrl.Call(_m, "Detach", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Detach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Detach", arg0, arg1)
}

func (_m *MockVolumeDriver) Enumerate(_param0 *api.VolumeLocator, _param1 map[string]string) ([]*api.Volume, error) {
	ret := _m.ctrl.Call(_m, "Enumerate", _param0, _param1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Enumerate", arg0, arg1)
}

func (_m *MockVolumeDriver) FilesystemCheckStart(_param0 *api.SdkFilesystemCheckStartRequest) (*api.SdkFilesystemCheckStartResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemCheckStart", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemCheckStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemCheckStart", arg0)
}

func (_m *MockVolumeDriver) FilesystemCheckStatus(_param0 *api.SdkFilesystemCheckStatusRequest) (*api.SdkFilesystemCheckStatusResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemCheckStatus", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemCheckStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemCheckStatus", arg0)
}

func (_m *MockVolumeDriver) FilesystemCheckStop(_param0 *api.SdkFilesystemCheckStopRequest) (*api.SdkFilesystemCheckStopResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemCheckStop", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemCheckStop(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemCheckStop", arg0)
}

func (_m *MockVolumeDriver) FilesystemTrimStart(_param0 *api.SdkFilesystemTrimStartRequest) (*api.SdkFilesystemTrimStartResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemTrimStart", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemTrimStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemTrimStart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemTrimStart", arg0)
}

func (_m *MockVolumeDriver) FilesystemTrimStatus(_param0 *api.SdkFilesystemTrimStatusRequest) (*api.SdkFilesystemTrimStatusResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemTrimStatus", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemTrimStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemTrimStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemTrimStatus", arg0)
}

func (_m *MockVolumeDriver) FilesystemTrimStop(_param0 *api.SdkFilesystemTrimStopRequest) (*api.SdkFilesystemTrimStopResponse, error) {
	ret := _m.ctrl.Call(_m, "FilesystemTrimStop", _param0)
	ret0, _ := ret[0].(*api.SdkFilesystemTrimStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) FilesystemTrimStop(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilesystemTrimStop", arg0)
}

func (_m *MockVolumeDriver) Flush(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Flush", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

func (_m *MockVolumeDriver) GetActiveRequests() (*api.ActiveRequests, error) {
	ret := _m.ctrl.Call(_m, "GetActiveRequests")
	ret0, _ := ret[0].(*api.ActiveRequests)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) GetActiveRequests() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetActiveRequests")
}

func (_m *MockVolumeDriver) Inspect(_param0 []string) ([]*api.Volume, error) {
	ret := _m.ctrl.Call(_m, "Inspect", _param0)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Inspect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Inspect", arg0)
}

func (_m *MockVolumeDriver) Mount(_param0 string, _param1 string, _param2 map[string]string) error {
	ret := _m.ctrl.Call(_m, "Mount", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Mount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Mount", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) MountedAt(_param0 string) string {
	ret := _m.ctrl.Call(_m, "MountedAt", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) MountedAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountedAt", arg0)
}

func (_m *MockVolumeDriver) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockVolumeDriver) Quiesce(_param0 string, _param1 uint64, _param2 string) error {
	ret := _m.ctrl.Call(_m, "Quiesce", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Quiesce(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Quiesce", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) Read(_param0 string, _param1 []byte, _param2 uint64, _param3 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Read(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1, arg2, arg3)
}

func (_m *MockVolumeDriver) Restore(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Restore", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Restore", arg0, arg1)
}

func (_m *MockVolumeDriver) Set(_param0 string, _param1 *api.VolumeLocator, _param2 *api.VolumeSpec) error {
	ret := _m.ctrl.Call(_m, "Set", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) Shutdown() {
	_m.ctrl.Call(_m, "Shutdown")
}

func (_mr *_MockVolumeDriverRecorder) Shutdown() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shutdown")
}

func (_m *MockVolumeDriver) SnapEnumerate(_param0 []string, _param1 map[string]string) ([]*api.Volume, error) {
	ret := _m.ctrl.Call(_m, "SnapEnumerate", _param0, _param1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) SnapEnumerate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SnapEnumerate", arg0, arg1)
}

func (_m *MockVolumeDriver) Snapshot(_param0 string, _param1 bool, _param2 *api.VolumeLocator, _param3 bool) (string, error) {
	ret := _m.ctrl.Call(_m, "Snapshot", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Snapshot(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Snapshot", arg0, arg1, arg2, arg3)
}

func (_m *MockVolumeDriver) SnapshotGroup(_param0 string, _param1 map[string]string, _param2 []string, _param3 bool) (*api.GroupSnapCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "SnapshotGroup", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(*api.GroupSnapCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) SnapshotGroup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SnapshotGroup", arg0, arg1, arg2, arg3)
}

func (_m *MockVolumeDriver) Stats(_param0 string, _param1 bool) (*api.Stats, error) {
	ret := _m.ctrl.Call(_m, "Stats", _param0, _param1)
	ret0, _ := ret[0].(*api.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stats", arg0, arg1)
}

func (_m *MockVolumeDriver) Status() [][2]string {
	ret := _m.ctrl.Call(_m, "Status")
	ret0, _ := ret[0].([][2]string)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Status() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Status")
}

func (_m *MockVolumeDriver) Type() api.DriverType {
	ret := _m.ctrl.Call(_m, "Type")
	ret0, _ := ret[0].(api.DriverType)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Type() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Type")
}

func (_m *MockVolumeDriver) Unmount(_param0 string, _param1 string, _param2 map[string]string) error {
	ret := _m.ctrl.Call(_m, "Unmount", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Unmount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unmount", arg0, arg1, arg2)
}

func (_m *MockVolumeDriver) Unquiesce(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Unquiesce", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVolumeDriverRecorder) Unquiesce(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unquiesce", arg0)
}

func (_m *MockVolumeDriver) UsedSize(_param0 string) (uint64, error) {
	ret := _m.ctrl.Call(_m, "UsedSize", _param0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) UsedSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UsedSize", arg0)
}

func (_m *MockVolumeDriver) Version() (*api.StorageVersion, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(*api.StorageVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockVolumeDriver) VolService(_param0 string, _param1 *api.VolumeServiceRequest) (*api.VolumeServiceResponse, error) {
	ret := _m.ctrl.Call(_m, "VolService", _param0, _param1)
	ret0, _ := ret[0].(*api.VolumeServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) VolService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolService", arg0, arg1)
}

func (_m *MockVolumeDriver) Write(_param0 string, _param1 []byte, _param2 uint64, _param3 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVolumeDriverRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3)
}
