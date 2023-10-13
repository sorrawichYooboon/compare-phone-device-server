package server

import (
	usecases "flowech-device-server/domain/usecases/device"
	controllers "flowech-device-server/presentation/controllers/health_check"
	"flowech-device-server/presentation/routes"

	"github.com/gin-gonic/gin"
)

func Start() error {
	app := gin.Default()

	deviceUsecase := usecases.NewDeviceUsecase()

	healthCheckController := controllers.NewHealthCheckController(deviceUsecase)

	routes.ApplyHealthCheckRoutes(app, &routes.HealthCheckHttpRoutes{
		HealthCheckController: healthCheckController,
	})
	// routes.ApplyDeviceHealthCheckRoutes(app)

	return app.Run(":8080")
}
