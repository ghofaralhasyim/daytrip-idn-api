package main

import (
	"log"
	"os"

	"github.com/daytrip-idn-api/internal/middleware"
	"github.com/daytrip-idn-api/internal/modules"
	"github.com/daytrip-idn-api/internal/routes"
	"github.com/daytrip-idn-api/pkg/database"
	"github.com/daytrip-idn-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	echo := echo.New()

	v := validator.New()
	echo.Validator = &CustomValidator{validator: v}

	echo.Use(middleware.CORSMiddleware)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDbConnection()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	modules := modules.NewAppModules(db)

	routes.SetupRoutes(echo, modules)

	if os.Getenv("STAGE_STATUS") == "production" {
		utils.StartServerWithGracefulShutdown(echo)
	} else {
		utils.StartServer(echo)
	}
}
