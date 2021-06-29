// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scan "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/scan"
)

// MockScanClientService is an autogenerated mock type for the ClientService type
type MockScanClientService struct {
	mock.Mock
}

// GetReportLog provides a mock function with given fields: params, authInfo, opts
func (_m *MockScanClientService) GetReportLog(params *scan.GetReportLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...scan.ClientOption) (*scan.GetReportLogOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scan.GetReportLogOK
	if rf, ok := ret.Get(0).(func(*scan.GetReportLogParams, runtime.ClientAuthInfoWriter, ...scan.ClientOption) *scan.GetReportLogOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.GetReportLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.GetReportLogParams, runtime.ClientAuthInfoWriter, ...scan.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanArtifact provides a mock function with given fields: params, authInfo, opts
func (_m *MockScanClientService) ScanArtifact(params *scan.ScanArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...scan.ClientOption) (*scan.ScanArtifactAccepted, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scan.ScanArtifactAccepted
	if rf, ok := ret.Get(0).(func(*scan.ScanArtifactParams, runtime.ClientAuthInfoWriter, ...scan.ClientOption) *scan.ScanArtifactAccepted); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.ScanArtifactAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.ScanArtifactParams, runtime.ClientAuthInfoWriter, ...scan.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScanClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
