package controllers

import (
	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

func (c *healthCheckController) Ping(ctx *gin.Context) {
	ctx.JSON(200, PingResponse{
		Message: "pong",
	})
}
