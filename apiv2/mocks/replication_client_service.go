// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	replication "github.com/simbou2000/goharbor-client/v5/apiv2/internal/api/client/replication"
	mock "github.com/stretchr/testify/mock"
)

// MockReplicationClientService is an autogenerated mock type for the ClientService type
type MockReplicationClientService struct {
	mock.Mock
}

// CreateReplicationPolicy provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) CreateReplicationPolicy(params *replication.CreateReplicationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*replication.CreateReplicationPolicyCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.CreateReplicationPolicyCreated
	if rf, ok := ret.Get(0).(func(*replication.CreateReplicationPolicyParams, runtime.ClientAuthInfoWriter) *replication.CreateReplicationPolicyCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.CreateReplicationPolicyCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.CreateReplicationPolicyParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReplicationPolicy provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) DeleteReplicationPolicy(params *replication.DeleteReplicationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*replication.DeleteReplicationPolicyOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.DeleteReplicationPolicyOK
	if rf, ok := ret.Get(0).(func(*replication.DeleteReplicationPolicyParams, runtime.ClientAuthInfoWriter) *replication.DeleteReplicationPolicyOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.DeleteReplicationPolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.DeleteReplicationPolicyParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationExecution provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) GetReplicationExecution(params *replication.GetReplicationExecutionParams, authInfo runtime.ClientAuthInfoWriter) (*replication.GetReplicationExecutionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.GetReplicationExecutionOK
	if rf, ok := ret.Get(0).(func(*replication.GetReplicationExecutionParams, runtime.ClientAuthInfoWriter) *replication.GetReplicationExecutionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.GetReplicationExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.GetReplicationExecutionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationLog provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) GetReplicationLog(params *replication.GetReplicationLogParams, authInfo runtime.ClientAuthInfoWriter) (*replication.GetReplicationLogOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.GetReplicationLogOK
	if rf, ok := ret.Get(0).(func(*replication.GetReplicationLogParams, runtime.ClientAuthInfoWriter) *replication.GetReplicationLogOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.GetReplicationLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.GetReplicationLogParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationPolicy provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) GetReplicationPolicy(params *replication.GetReplicationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*replication.GetReplicationPolicyOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.GetReplicationPolicyOK
	if rf, ok := ret.Get(0).(func(*replication.GetReplicationPolicyParams, runtime.ClientAuthInfoWriter) *replication.GetReplicationPolicyOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.GetReplicationPolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.GetReplicationPolicyParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationExecutions provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) ListReplicationExecutions(params *replication.ListReplicationExecutionsParams, authInfo runtime.ClientAuthInfoWriter) (*replication.ListReplicationExecutionsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.ListReplicationExecutionsOK
	if rf, ok := ret.Get(0).(func(*replication.ListReplicationExecutionsParams, runtime.ClientAuthInfoWriter) *replication.ListReplicationExecutionsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.ListReplicationExecutionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.ListReplicationExecutionsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationPolicies provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) ListReplicationPolicies(params *replication.ListReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*replication.ListReplicationPoliciesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.ListReplicationPoliciesOK
	if rf, ok := ret.Get(0).(func(*replication.ListReplicationPoliciesParams, runtime.ClientAuthInfoWriter) *replication.ListReplicationPoliciesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.ListReplicationPoliciesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.ListReplicationPoliciesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationTasks provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) ListReplicationTasks(params *replication.ListReplicationTasksParams, authInfo runtime.ClientAuthInfoWriter) (*replication.ListReplicationTasksOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.ListReplicationTasksOK
	if rf, ok := ret.Get(0).(func(*replication.ListReplicationTasksParams, runtime.ClientAuthInfoWriter) *replication.ListReplicationTasksOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.ListReplicationTasksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.ListReplicationTasksParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockReplicationClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// StartReplication provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) StartReplication(params *replication.StartReplicationParams, authInfo runtime.ClientAuthInfoWriter) (*replication.StartReplicationCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.StartReplicationCreated
	if rf, ok := ret.Get(0).(func(*replication.StartReplicationParams, runtime.ClientAuthInfoWriter) *replication.StartReplicationCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.StartReplicationCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.StartReplicationParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopReplication provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) StopReplication(params *replication.StopReplicationParams, authInfo runtime.ClientAuthInfoWriter) (*replication.StopReplicationOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.StopReplicationOK
	if rf, ok := ret.Get(0).(func(*replication.StopReplicationParams, runtime.ClientAuthInfoWriter) *replication.StopReplicationOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.StopReplicationOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.StopReplicationParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReplicationPolicy provides a mock function with given fields: params, authInfo
func (_m *MockReplicationClientService) UpdateReplicationPolicy(params *replication.UpdateReplicationPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*replication.UpdateReplicationPolicyOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *replication.UpdateReplicationPolicyOK
	if rf, ok := ret.Get(0).(func(*replication.UpdateReplicationPolicyParams, runtime.ClientAuthInfoWriter) *replication.UpdateReplicationPolicyOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.UpdateReplicationPolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.UpdateReplicationPolicyParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockReplicationClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockReplicationClientService creates a new instance of MockReplicationClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockReplicationClientService(t mockConstructorTestingTNewMockReplicationClientService) *MockReplicationClientService {
	mock := &MockReplicationClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
