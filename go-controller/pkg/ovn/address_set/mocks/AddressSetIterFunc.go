// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AddressSetIterFunc is an autogenerated mock type for the AddressSetIterFunc type
type AddressSetIterFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: hashedName, name
func (_m *AddressSetIterFunc) Execute(hashedName string, name string) error {
	ret := _m.Called(hashedName, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashedName, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAddressSetIterFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewAddressSetIterFunc creates a new instance of AddressSetIterFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAddressSetIterFunc(t mockConstructorTestingTNewAddressSetIterFunc) *AddressSetIterFunc {
	mock := &AddressSetIterFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}