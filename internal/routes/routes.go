package routes

import (
	"database/sql"

	"github.com/daytrip-idn-api/internal/controllers"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	apiv1 := e.Group("/v1")

	userRepo := repositories.NewUserRepository(db)
	userService := usecases.NewUserService(userRepo)
	userHandler := controllers.NewUserHandler(userService)
	apiv1.POST("/login", userHandler.Login)

}
