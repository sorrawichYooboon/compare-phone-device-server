package mocks

import (
	"github.com/gin-gonic/gin"
)

type MockDeviceUsecase struct {
	PingIsCalled             bool
	GetDeviceDbNamesIsCalled bool
}

func (m *MockDeviceUsecase) Ping(ctx *gin.Context) {
	m.PingIsCalled = true
}

func (m *MockDeviceUsecase) GetDeviceDbNames(ctx *gin.Context) {
	m.GetDeviceDbNamesIsCalled = true
}
