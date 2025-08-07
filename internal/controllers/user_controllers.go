package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService usecases.UserService
}

func NewUserHandler(userService usecases.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type userLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) Login(c echo.Context) error {
	var req userLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		var validationErrors []map[string]string

		for _, e := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(e.Field())
			friendlyMessage := utils.GetFriendlyErrorMessage(e)

			validationErrors = append(validationErrors, map[string]string{
				fieldName: friendlyMessage,
			})
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": validationErrors,
		})
	}

	user, token, err := h.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Println(err)
		if err.Error() == "user not found" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid email or password",
				"details": nil,
			})
		}
		if err.Error() == "unauthorize" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid email or password",
				"details": nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"details": nil,
		})
	}

	user.PasswordHash = ""

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"token":   token,
		"data":    user,
	})
}
