package routes

import (
	"database/sql"

	"github.com/daytrip-idn-api/internal/http"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/internal/services"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	apiv1 := e.Group("/v1")

	apiv1.Static("/public/images/articles", "public/images/articles")

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)
	apiv1.POST("/login", userHandler.Login)

}
