package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (c *InvitationController) GetAttendanceBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")
	if slug == "" {
		err := fmt.Errorf("invalid request")
		return helpers.EchoError(ctx, error_app.NewAppError(error_data.InvalidDataRequest, err))
	}

	result, err := c.invitationResponseUsecase.GetInvitationResponseBySlug(ctx, slug)
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

	result.KeyPass = ""

	response := response_transform.TransformInvitationResponse(result)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *InvitationController) GetAdminInvitationBySlug(ctx echo.Context) error {
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

func (c *InvitationController) DeleteInvitation(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		err := fmt.Errorf("invalid request")
		return helpers.EchoError(ctx, error_app.NewAppError(error_data.InvalidDataRequest, err))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		err := fmt.Errorf("invalid request")
		return helpers.EchoError(ctx, error_app.NewAppError(error_data.InvalidDataRequest, err))
	}

	err = c.invitationUsecase.DeleteInvitation(ctx, idInt)
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})

}

func (c *InvitationController) UpdateInvitation(ctx echo.Context) error {
	slug := ctx.Param("slug")
	if slug == "" {
		err := fmt.Errorf("invalid request")
		return helpers.EchoError(ctx, error_app.NewAppError(error_data.InvalidDataRequest, err))
	}

	var req rest_request.InvtationRequest
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

	imageFile, _ := ctx.FormFile("image")
	imageFile1, _ := ctx.FormFile("image1")

	reqEntity := entities.InvitationEntity{
		Title:       req.Title,
		Slug:        slug,
		Description: &req.Description,
		TemplateId:  &req.TemplateId,
		StartDate:   &req.StartDate,
		EndDate:     &req.EndDate,
		MapsUrl:     &req.MapsUrl,
		Address:     &req.Address,
		Location:    &req.Location,
		DressCode:   &req.DressCode,
		KeyPass:     req.KeyPass,
		ImageFile:   imageFile,
		ImageFile1:  imageFile1,
	}

	if req.TemplateId == 2 {
		reqEntity.BirthdayVal = &req.BirthdayVal
	}

	data, err := c.invitationUsecase.UpdateInvitation(ctx, reqEntity)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformInvitationResponse(data)

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "updated",
		"data":    response,
	})
}

func (c *InvitationController) CreateInvitation(ctx echo.Context) error {
	var req rest_request.InvtationRequest
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

	imageFile, err := ctx.FormFile("image")
	if err != nil || imageFile == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": "image required",
		})
	}
	imageFile1, err := ctx.FormFile("image1")
	if err != nil || imageFile1 == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": "image required",
		})
	}

	reqEntity := entities.InvitationEntity{
		Title:       req.Title,
		Description: &req.Description,
		TemplateId:  &req.TemplateId,
		StartDate:   &req.StartDate,
		EndDate:     &req.EndDate,
		MapsUrl:     &req.MapsUrl,
		Address:     &req.Address,
		Location:    &req.Location,
		DressCode:   &req.DressCode,
		KeyPass:     req.KeyPass,
		ImageFile:   imageFile,
		ImageFile1:  imageFile1,
	}

	if req.TemplateId == 2 {
		reqEntity.BirthdayVal = &req.BirthdayVal
	} else {
		var val int64
		reqEntity.BirthdayVal = &val
	}

	data, err := c.invitationUsecase.CreateInvitation(ctx, reqEntity)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformInvitationResponse(data)

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "created",
		"data":    response,
	})
}
