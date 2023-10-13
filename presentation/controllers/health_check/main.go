package controllers

import (
	usecase "flowech-device-server/domain/usecases/device"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	Ping(ctx *gin.Context)
	GetDeviceDbNames(ctx *gin.Context)
}

type healthCheckController struct {
	DeviceUsecase usecase.DeviceUsecase
}

func NewHealthCheckController(deviceUsecase usecase.DeviceUsecase) HealthCheckController {
	return &healthCheckController{
		DeviceUsecase: deviceUsecase,
	}
}
