package routes

import (
	controllers "flowech-device-server/presentation/controllers/health_check"

	"github.com/gin-gonic/gin"
)

type Config struct {
	App App `yaml:"app"`
}

type App struct {
	MongoDB MongoDB `yaml:"mongodb"`
}

type MongoDB struct {
	URI string `yaml:"uri"`
}

type HealthCheckHttpRoutes struct {
	HealthCheckController controllers.HealthCheckController
}

func ApplyHealthCheckRoutes(router *gin.Engine, httpRoutes *HealthCheckHttpRoutes) {
	healthCheckRoutes := router.Group("/health")
	{
		healthCheckRoutes.GET("/ping", httpRoutes.HealthCheckController.Ping)
		healthCheckRoutes.GET("/device-db", httpRoutes.HealthCheckController.GetDeviceDbNames)
	}
}
