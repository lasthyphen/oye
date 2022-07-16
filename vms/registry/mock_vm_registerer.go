// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: vms/registry/vm_registerer.go

// Package registry is a generated GoMock package.
package registry

import (
	reflect "reflect"

	ids "github.com/lasthyphen/beacongo/ids"
	vms "github.com/lasthyphen/beacongo/vms"
	gomock "github.com/golang/mock/gomock"
)

// MockVMRegisterer is a mock of VMRegisterer interface.
type MockVMRegisterer struct {
	ctrl     *gomock.Controller
	recorder *MockVMRegistererMockRecorder
}

// MockVMRegistererMockRecorder is the mock recorder for MockVMRegisterer.
type MockVMRegistererMockRecorder struct {
	mock *MockVMRegisterer
}

// NewMockVMRegisterer creates a new mock instance.
func NewMockVMRegisterer(ctrl *gomock.Controller) *MockVMRegisterer {
	mock := &MockVMRegisterer{ctrl: ctrl}
	mock.recorder = &MockVMRegistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMRegisterer) EXPECT() *MockVMRegistererMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockVMRegisterer) Register(arg0 ids.ID, arg1 vms.Factory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockVMRegistererMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockVMRegisterer)(nil).Register), arg0, arg1)
}

// RegisterWithReadLock mocks base method.
func (m *MockVMRegisterer) RegisterWithReadLock(arg0 ids.ID, arg1 vms.Factory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterWithReadLock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterWithReadLock indicates an expected call of RegisterWithReadLock.
func (mr *MockVMRegistererMockRecorder) RegisterWithReadLock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterWithReadLock", reflect.TypeOf((*MockVMRegisterer)(nil).RegisterWithReadLock), arg0, arg1)
}

// Mockregisterer is a mock of registerer interface.
type Mockregisterer struct {
	ctrl     *gomock.Controller
	recorder *MockregistererMockRecorder
}

// MockregistererMockRecorder is the mock recorder for Mockregisterer.
type MockregistererMockRecorder struct {
	mock *Mockregisterer
}

// NewMockregisterer creates a new mock instance.
func NewMockregisterer(ctrl *gomock.Controller) *Mockregisterer {
	mock := &Mockregisterer{ctrl: ctrl}
	mock.recorder = &MockregistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockregisterer) EXPECT() *MockregistererMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *Mockregisterer) Register(arg0 ids.ID, arg1 vms.Factory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockregistererMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*Mockregisterer)(nil).Register), arg0, arg1)
}
