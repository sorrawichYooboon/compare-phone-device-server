package controllers

import (
	"github.com/gin-gonic/gin"
)

func (c *healthCheckController) GetDeviceDbNames(ctx *gin.Context) {
	c.DeviceUsecase.GetDeviceDbNames(ctx)
}
