// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/chroma/chroma-coordinator/internal/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// INotificationDb is an autogenerated mock type for the INotificationDb type
type INotificationDb struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *INotificationDb) Delete(id []int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAll provides a mock function with given fields:
func (_m *INotificationDb) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPendingNotifications provides a mock function with given fields:
func (_m *INotificationDb) GetAllPendingNotifications() ([]*dbmodel.Notification, error) {
	ret := _m.Called()

	var r0 []*dbmodel.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*dbmodel.Notification, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*dbmodel.Notification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationByCollectionID provides a mock function with given fields: collectionID
func (_m *INotificationDb) GetNotificationByCollectionID(collectionID string) ([]*dbmodel.Notification, error) {
	ret := _m.Called(collectionID)

	var r0 []*dbmodel.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*dbmodel.Notification, error)); ok {
		return rf(collectionID)
	}
	if rf, ok := ret.Get(0).(func(string) []*dbmodel.Notification); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *INotificationDb) Insert(in *dbmodel.Notification) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Notification) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewINotificationDb creates a new instance of INotificationDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewINotificationDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *INotificationDb {
	mock := &INotificationDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
