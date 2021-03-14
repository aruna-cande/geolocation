// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/geolocation/service/geolocationData.go

// Package mock_service is a generated GoMock package.
package mock

import (
	domain "Geolocation/pkg/geolocation/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGeolocationDataService is a mock of GeolocationDataService interface.
type MockGeolocationDataService struct {
	ctrl     *gomock.Controller
	recorder *MockGeolocationDataServiceMockRecorder
}

// MockGeolocationDataServiceMockRecorder is the mock recorder for MockGeolocationDataService.
type MockGeolocationDataServiceMockRecorder struct {
	mock *MockGeolocationDataService
}

// NewMockGeolocationDataService creates a new mock instance.
func NewMockGeolocationDataService(ctrl *gomock.Controller) *MockGeolocationDataService {
	mock := &MockGeolocationDataService{ctrl: ctrl}
	mock.recorder = &MockGeolocationDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeolocationDataService) EXPECT() *MockGeolocationDataServiceMockRecorder {
	return m.recorder
}

// GetGeolocationByIp mocks base method.
func (m *MockGeolocationDataService) GetGeolocationByIp(ipAdrress string) (*domain.Geolocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeolocationByIp", ipAdrress)
	ret0, _ := ret[0].(*domain.Geolocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeolocationByIp indicates an expected call of GetGeolocationByIp.
func (mr *MockGeolocationDataServiceMockRecorder) GetGeolocationByIp(ipAdrress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeolocationByIp", reflect.TypeOf((*MockGeolocationDataService)(nil).GetGeolocationByIp), ipAdrress)
}
