// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	recaptcha "github.com/hiroaki-yamamoto/recaptcha"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockIHttpClient is a mock of IHttpClient interface
type MockIHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockIHttpClientMockRecorder
}

// MockIHttpClientMockRecorder is the mock recorder for MockIHttpClient
type MockIHttpClientMockRecorder struct {
	mock *MockIHttpClient
}

// NewMockIHttpClient creates a new mock instance
func NewMockIHttpClient(ctrl *gomock.Controller) *MockIHttpClient {
	mock := &MockIHttpClient{ctrl: ctrl}
	mock.recorder = &MockIHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHttpClient) EXPECT() *MockIHttpClientMockRecorder {
	return m.recorder
}

// PostForm mocks base method
func (m *MockIHttpClient) PostForm(url string, data url.Values) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", url, data)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostForm indicates an expected call of PostForm
func (mr *MockIHttpClientMockRecorder) PostForm(url, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockIHttpClient)(nil).PostForm), url, data)
}

// MockIRecaptcha is a mock of IRecaptcha interface
type MockIRecaptcha struct {
	ctrl     *gomock.Controller
	recorder *MockIRecaptchaMockRecorder
}

// MockIRecaptchaMockRecorder is the mock recorder for MockIRecaptcha
type MockIRecaptchaMockRecorder struct {
	mock *MockIRecaptcha
}

// NewMockIRecaptcha creates a new mock instance
func NewMockIRecaptcha(ctrl *gomock.Controller) *MockIRecaptcha {
	mock := &MockIRecaptcha{ctrl: ctrl}
	mock.recorder = &MockIRecaptchaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRecaptcha) EXPECT() *MockIRecaptchaMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockIRecaptcha) Check(remoteIP, response string) (*recaptcha.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", remoteIP, response)
	ret0, _ := ret[0].(*recaptcha.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check
func (mr *MockIRecaptchaMockRecorder) Check(remoteIP, response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIRecaptcha)(nil).Check), remoteIP, response)
}
