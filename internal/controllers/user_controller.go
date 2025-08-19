package controllers

import (
	"fmt"
	"net/http"
	"strings"

	rest_request "github.com/daytrip-idn-api/internal/rest/request"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) Login(ctx echo.Context) error {
	var req rest_request.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := ctx.Validate(&req); err != nil {
		var validationErrors []map[string]string

		for _, e := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(e.Field())
			friendlyMessage := helpers.GetFriendlyErrorMessage(e)

			validationErrors = append(validationErrors, map[string]string{
				fieldName: friendlyMessage,
			})
		}

		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": validationErrors,
		})
	}

	user, token, err := c.userUsecase.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
		return helpers.EchoError(ctx, err)
	}

	user.PasswordHash = ""

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"token":   token,
		"data":    user,
	})
}
