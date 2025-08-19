package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/daytrip-idn-api/internal/entities"
	error_app "github.com/daytrip-idn-api/internal/error"
	rest_request "github.com/daytrip-idn-api/internal/rest/request"
	response_transform "github.com/daytrip-idn-api/internal/rest/transform"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type MessageController struct {
	messageUsecase usecases.MessageUsecase
}

func NewMessageController(
	messageUsecase usecases.MessageUsecase,
) *MessageController {
	return &MessageController{
		messageUsecase: messageUsecase,
	}
}

func (c *MessageController) GetMessages(ctx echo.Context) error {
	messages, err := c.messageUsecase.GetMessages(ctx)
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformListMessageResponse(messages)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *MessageController) InsertMessage(ctx echo.Context) error {
	var req rest_request.MessageRequest
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

	reqEntity := entities.MessageEntity{
		Phone:       req.Phone,
		Email:       req.Email,
		PackageName: req.PackageName,
		Message:     req.Message,
	}

	result, err := c.messageUsecase.InsertMessage(ctx, reqEntity)
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformMessageResponse(result)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *MessageController) DeleteMessage(ctx echo.Context) error {
	messageId := ctx.Param("messageId")
	messageIdInt, err := strconv.Atoi(messageId)
	if err != nil {
		appErr := error_app.NewAppError(error_app.UsecaseValidateError, err)
		return helpers.EchoError(ctx, appErr)
	}

	err = c.messageUsecase.DeleteMessage(ctx, int64(messageIdInt))
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})
}
