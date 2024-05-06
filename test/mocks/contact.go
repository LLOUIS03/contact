// Code generated by MockGen. DO NOT EDIT.
// Source: domain/services/contact/contact.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	repos "github.com/contact/infraestructure/db/repos"
	gomock "github.com/golang/mock/gomock"
)

// MockContact is a mock of Contact interface.
type MockContact struct {
	ctrl     *gomock.Controller
	recorder *MockContactMockRecorder
}

// MockContactMockRecorder is the mock recorder for MockContact.
type MockContactMockRecorder struct {
	mock *MockContact
}

// NewMockContact creates a new mock instance.
func NewMockContact(ctrl *gomock.Controller) *MockContact {
	mock := &MockContact{ctrl: ctrl}
	mock.recorder = &MockContactMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContact) EXPECT() *MockContactMockRecorder {
	return m.recorder
}

// AddContact mocks base method.
func (m *MockContact) AddContact(arg0 context.Context, arg1 repos.Contact) (*repos.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddContact", arg0, arg1)
	ret0, _ := ret[0].(*repos.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddContact indicates an expected call of AddContact.
func (mr *MockContactMockRecorder) AddContact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContact", reflect.TypeOf((*MockContact)(nil).AddContact), arg0, arg1)
}

// DeleteContact mocks base method.
func (m *MockContact) DeleteContact(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContact", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContact indicates an expected call of DeleteContact.
func (mr *MockContactMockRecorder) DeleteContact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContact", reflect.TypeOf((*MockContact)(nil).DeleteContact), arg0, arg1)
}

// GetContactByID mocks base method.
func (m *MockContact) GetContactByID(arg0 context.Context, arg1 string) (*repos.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactByID", arg0, arg1)
	ret0, _ := ret[0].(*repos.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactByID indicates an expected call of GetContactByID.
func (mr *MockContactMockRecorder) GetContactByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactByID", reflect.TypeOf((*MockContact)(nil).GetContactByID), arg0, arg1)
}

// GetContacts mocks base method.
func (m *MockContact) GetContacts(arg0 context.Context) ([]repos.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContacts", arg0)
	ret0, _ := ret[0].([]repos.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContacts indicates an expected call of GetContacts.
func (mr *MockContactMockRecorder) GetContacts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContacts", reflect.TypeOf((*MockContact)(nil).GetContacts), arg0)
}

// UpdateContact mocks base method.
func (m *MockContact) UpdateContact(arg0 context.Context, arg1 repos.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContact", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContact indicates an expected call of UpdateContact.
func (mr *MockContactMockRecorder) UpdateContact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContact", reflect.TypeOf((*MockContact)(nil).UpdateContact), arg0, arg1)
}
