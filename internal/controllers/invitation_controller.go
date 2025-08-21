package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/daytrip-idn-api/internal/entities"
	error_app "github.com/daytrip-idn-api/internal/error/app"
	error_data "github.com/daytrip-idn-api/internal/error/data"
	rest_request "github.com/daytrip-idn-api/internal/rest/request"
	response_transform "github.com/daytrip-idn-api/internal/rest/transform"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type InvitationController struct {
	invitationResponseUsecase usecases.InvitationResponseUsecase
	invitationUsecase         usecases.InvitationUsecase
}

func NewInvitationController(
	invitationResponseUsecase usecases.InvitationResponseUsecase,
	invitationUsecase usecases.InvitationUsecase,
) *InvitationController {
	return &InvitationController{
		invitationResponseUsecase: invitationResponseUsecase,
		invitationUsecase:         invitationUsecase,
	}
}

func (c *InvitationController) InsertAttendance(ctx echo.Context) error {
	var req rest_request.InvitationResponseRequest
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

	reqEntity := entities.InvitationResponseEntity{
		InvitationId: req.InvitationId,
		IsAttending:  req.IsAttending,
		Message:      req.Message,
		Name:         req.Name,
	}

	id, err := c.invitationResponseUsecase.SubmitResponse(ctx, reqEntity)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	reqEntity.Id = id

	response := response_transform.TransformInvitationResponseResponse(&reqEntity)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *InvitationController) GetAttendance(ctx echo.Context) error {
	result, err := c.invitationResponseUsecase.GetInvitationResponse(ctx)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformListInvitationResponseResponse(result)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *InvitationController) GetInvitations(ctx echo.Context) error {
	result, err := c.invitationUsecase.GetInvitations(ctx)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformListInvitationResponse(result)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *InvitationController) GetInvitationBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")
	if slug == "" {
		err := fmt.Errorf("invalid request")
		return helpers.EchoError(ctx, error_app.NewAppError(error_data.InvalidDataRequest, err))
	}

	result, err := c.invitationUsecase.GetInvitationBySlug(ctx, slug)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformInvitationResponse(result)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}
