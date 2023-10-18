// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	registry "github.com/simbou2000/goharbor-client/v5/apiv2/internal/api/client/registry"
	mock "github.com/stretchr/testify/mock"
)

// MockRegistryClientService is an autogenerated mock type for the ClientService type
type MockRegistryClientService struct {
	mock.Mock
}

// CreateRegistry provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) CreateRegistry(params *registry.CreateRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*registry.CreateRegistryCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.CreateRegistryCreated
	if rf, ok := ret.Get(0).(func(*registry.CreateRegistryParams, runtime.ClientAuthInfoWriter) *registry.CreateRegistryCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.CreateRegistryCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.CreateRegistryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRegistry provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) DeleteRegistry(params *registry.DeleteRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*registry.DeleteRegistryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.DeleteRegistryOK
	if rf, ok := ret.Get(0).(func(*registry.DeleteRegistryParams, runtime.ClientAuthInfoWriter) *registry.DeleteRegistryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.DeleteRegistryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.DeleteRegistryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistry provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) GetRegistry(params *registry.GetRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*registry.GetRegistryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.GetRegistryOK
	if rf, ok := ret.Get(0).(func(*registry.GetRegistryParams, runtime.ClientAuthInfoWriter) *registry.GetRegistryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.GetRegistryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.GetRegistryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistryInfo provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) GetRegistryInfo(params *registry.GetRegistryInfoParams, authInfo runtime.ClientAuthInfoWriter) (*registry.GetRegistryInfoOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.GetRegistryInfoOK
	if rf, ok := ret.Get(0).(func(*registry.GetRegistryInfoParams, runtime.ClientAuthInfoWriter) *registry.GetRegistryInfoOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.GetRegistryInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.GetRegistryInfoParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRegistries provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) ListRegistries(params *registry.ListRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*registry.ListRegistriesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.ListRegistriesOK
	if rf, ok := ret.Get(0).(func(*registry.ListRegistriesParams, runtime.ClientAuthInfoWriter) *registry.ListRegistriesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.ListRegistriesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.ListRegistriesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRegistryProviderInfos provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) ListRegistryProviderInfos(params *registry.ListRegistryProviderInfosParams, authInfo runtime.ClientAuthInfoWriter) (*registry.ListRegistryProviderInfosOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.ListRegistryProviderInfosOK
	if rf, ok := ret.Get(0).(func(*registry.ListRegistryProviderInfosParams, runtime.ClientAuthInfoWriter) *registry.ListRegistryProviderInfosOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.ListRegistryProviderInfosOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.ListRegistryProviderInfosParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRegistryProviderTypes provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) ListRegistryProviderTypes(params *registry.ListRegistryProviderTypesParams, authInfo runtime.ClientAuthInfoWriter) (*registry.ListRegistryProviderTypesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.ListRegistryProviderTypesOK
	if rf, ok := ret.Get(0).(func(*registry.ListRegistryProviderTypesParams, runtime.ClientAuthInfoWriter) *registry.ListRegistryProviderTypesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.ListRegistryProviderTypesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.ListRegistryProviderTypesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingRegistry provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) PingRegistry(params *registry.PingRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*registry.PingRegistryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.PingRegistryOK
	if rf, ok := ret.Get(0).(func(*registry.PingRegistryParams, runtime.ClientAuthInfoWriter) *registry.PingRegistryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.PingRegistryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.PingRegistryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRegistryClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateRegistry provides a mock function with given fields: params, authInfo
func (_m *MockRegistryClientService) UpdateRegistry(params *registry.UpdateRegistryParams, authInfo runtime.ClientAuthInfoWriter) (*registry.UpdateRegistryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *registry.UpdateRegistryOK
	if rf, ok := ret.Get(0).(func(*registry.UpdateRegistryParams, runtime.ClientAuthInfoWriter) *registry.UpdateRegistryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registry.UpdateRegistryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*registry.UpdateRegistryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRegistryClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRegistryClientService creates a new instance of MockRegistryClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRegistryClientService(t mockConstructorTestingTNewMockRegistryClientService) *MockRegistryClientService {
	mock := &MockRegistryClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
