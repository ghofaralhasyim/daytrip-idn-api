package routes

import (
	"github.com/daytrip-idn-api/internal/middleware"
	"github.com/daytrip-idn-api/internal/modules"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, m *modules.AppModules) {
	apiv1 := e.Group("/v1")

	apiv1.POST("/login", m.Controllers.User.Login)

	apiv1.GET("/banners", m.Controllers.Banner.GetBanners)

	apiv1.GET("/destinations", m.Controllers.Destination.GetDestinations)

	apiv1.GET("/messages",
		m.Controllers.Message.GetMessages,
		middleware.AuthMiddleware,
	)
	apiv1.POST("/messages",
		m.Controllers.Message.InsertMessage,
		middleware.AuthMiddleware,
	)
	apiv1.DELETE("/message/:messageId",
		m.Controllers.Message.DeleteMessage,
		middleware.AuthMiddleware,
	)
}
