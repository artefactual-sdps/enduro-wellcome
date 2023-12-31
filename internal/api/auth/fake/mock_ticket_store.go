// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/api/auth (interfaces: TicketStore)
//
// Generated by this command:
//
//	mockgen -typed -destination=./internal/api/auth/fake/mock_ticket_store.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TicketStore
//
// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockTicketStore is a mock of TicketStore interface.
type MockTicketStore struct {
	ctrl     *gomock.Controller
	recorder *MockTicketStoreMockRecorder
}

// MockTicketStoreMockRecorder is the mock recorder for MockTicketStore.
type MockTicketStoreMockRecorder struct {
	mock *MockTicketStore
}

// NewMockTicketStore creates a new mock instance.
func NewMockTicketStore(ctrl *gomock.Controller) *MockTicketStore {
	mock := &MockTicketStore{ctrl: ctrl}
	mock.recorder = &MockTicketStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketStore) EXPECT() *MockTicketStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTicketStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTicketStoreMockRecorder) Close() *TicketStoreCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTicketStore)(nil).Close))
	return &TicketStoreCloseCall{Call: call}
}

// TicketStoreCloseCall wrap *gomock.Call
type TicketStoreCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TicketStoreCloseCall) Return(arg0 error) *TicketStoreCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TicketStoreCloseCall) Do(f func() error) *TicketStoreCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TicketStoreCloseCall) DoAndReturn(f func() error) *TicketStoreCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetDel mocks base method.
func (m *MockTicketStore) GetDel(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetDel indicates an expected call of GetDel.
func (mr *MockTicketStoreMockRecorder) GetDel(arg0, arg1 any) *TicketStoreGetDelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockTicketStore)(nil).GetDel), arg0, arg1)
	return &TicketStoreGetDelCall{Call: call}
}

// TicketStoreGetDelCall wrap *gomock.Call
type TicketStoreGetDelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TicketStoreGetDelCall) Return(arg0 error) *TicketStoreGetDelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TicketStoreGetDelCall) Do(f func(context.Context, string) error) *TicketStoreGetDelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TicketStoreGetDelCall) DoAndReturn(f func(context.Context, string) error) *TicketStoreGetDelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetEx mocks base method.
func (m *MockTicketStore) SetEx(arg0 context.Context, arg1 string, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEx indicates an expected call of SetEx.
func (mr *MockTicketStoreMockRecorder) SetEx(arg0, arg1, arg2 any) *TicketStoreSetExCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEx", reflect.TypeOf((*MockTicketStore)(nil).SetEx), arg0, arg1, arg2)
	return &TicketStoreSetExCall{Call: call}
}

// TicketStoreSetExCall wrap *gomock.Call
type TicketStoreSetExCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TicketStoreSetExCall) Return(arg0 error) *TicketStoreSetExCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TicketStoreSetExCall) Do(f func(context.Context, string, time.Duration) error) *TicketStoreSetExCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TicketStoreSetExCall) DoAndReturn(f func(context.Context, string, time.Duration) error) *TicketStoreSetExCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
