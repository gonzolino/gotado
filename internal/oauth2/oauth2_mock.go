// Code generated by MockGen. DO NOT EDIT.
// Source: oauth2.go

// Package oauth2 is a generated GoMock package.
package oauth2

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
	http "net/http"
	reflect "reflect"
)

// MockConfigInterface is a mock of ConfigInterface interface
type MockConfigInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigInterfaceMockRecorder
}

// MockConfigInterfaceMockRecorder is the mock recorder for MockConfigInterface
type MockConfigInterfaceMockRecorder struct {
	mock *MockConfigInterface
}

// NewMockConfigInterface creates a new mock instance
func NewMockConfigInterface(ctrl *gomock.Controller) *MockConfigInterface {
	mock := &MockConfigInterface{ctrl: ctrl}
	mock.recorder = &MockConfigInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigInterface) EXPECT() *MockConfigInterfaceMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method
func (m *MockConfigInterface) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{state}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthCodeURL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthCodeURL indicates an expected call of AuthCodeURL
func (mr *MockConfigInterfaceMockRecorder) AuthCodeURL(state interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{state}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*MockConfigInterface)(nil).AuthCodeURL), varargs...)
}

// PasswordCredentialsToken mocks base method
func (m *MockConfigInterface) PasswordCredentialsToken(ctx context.Context, username, password string) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordCredentialsToken", ctx, username, password)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordCredentialsToken indicates an expected call of PasswordCredentialsToken
func (mr *MockConfigInterfaceMockRecorder) PasswordCredentialsToken(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordCredentialsToken", reflect.TypeOf((*MockConfigInterface)(nil).PasswordCredentialsToken), ctx, username, password)
}

// Exchange mocks base method
func (m *MockConfigInterface) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, code}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exchange", varargs...)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange
func (mr *MockConfigInterfaceMockRecorder) Exchange(ctx, code interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, code}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockConfigInterface)(nil).Exchange), varargs...)
}

// Client mocks base method
func (m *MockConfigInterface) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client", ctx, t)
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockConfigInterfaceMockRecorder) Client(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockConfigInterface)(nil).Client), ctx, t)
}

// TokenSource mocks base method
func (m *MockConfigInterface) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenSource", ctx, t)
	ret0, _ := ret[0].(oauth2.TokenSource)
	return ret0
}

// TokenSource indicates an expected call of TokenSource
func (mr *MockConfigInterfaceMockRecorder) TokenSource(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenSource", reflect.TypeOf((*MockConfigInterface)(nil).TokenSource), ctx, t)
}
