package usecases

import "github.com/gin-gonic/gin"

type DeviceUsecase interface {
	GetDeviceDbNames(ctx *gin.Context)
}

type deviceUsecase struct{}

func NewDeviceUsecase() DeviceUsecase {
	return &deviceUsecase{}
}
