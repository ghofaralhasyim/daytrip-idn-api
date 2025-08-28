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

type BannerController struct {
	bannerUsecase usecases.BannerUsecase
}

func NewBannerController(
	bannerUsecase usecases.BannerUsecase,
) *BannerController {
	return &BannerController{
		bannerUsecase: bannerUsecase,
	}
}

func (c *BannerController) GetBanners(ctx echo.Context) error {
	banners, err := c.bannerUsecase.GetBanners(ctx)
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformListBannerResponse(banners)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}

func (c *BannerController) CreateBanners(ctx echo.Context) error {
	var req rest_request.InsertBannerRequest
	if err := ctx.Bind(&req); err != nil {
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

	desktopImage, err := ctx.FormFile("desktopImage")
	if err != nil || desktopImage == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": "desktop image required",
		})
	}

	mobileImage, err := ctx.FormFile("mobileImage")
	if err != nil || mobileImage == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request - validation failed",
			"details": "mobile image required",
		})
	}

	reqEntity := entities.BannerEntity{
		Title:            req.Title,
		Description:      req.Description,
		Cta:              req.Cta,
		CtaUrl:           req.CtaUrl,
		MobileImageFile:  mobileImage,
		DesktopImageFile: desktopImage,
	}

	result, err := c.bannerUsecase.InsertBanner(ctx, reqEntity)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformBannerResponse(result)

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "created",
		"data":    response,
	})
}

func (c *BannerController) UpdateBanner(ctx echo.Context) error {
	var req rest_request.UpdateBannerRequest
	if err := ctx.Bind(&req); err != nil {
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

	reqEntity := entities.BannerEntity{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Cta:         req.Cta,
		CtaUrl:      req.CtaUrl,
	}

	desktopImage, _ := ctx.FormFile("desktopImage")
	mobileImage, _ := ctx.FormFile("mobileImage")

	if desktopImage != nil {
		reqEntity.DesktopImageFile = desktopImage
	}

	if mobileImage != nil {
		reqEntity.MobileImageFile = mobileImage
	}

	result, err := c.bannerUsecase.UpdateBanner(ctx, reqEntity)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformBannerResponse(result)

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "updated",
		"data":    response,
	})
}

func (c *BannerController) DeleteBanner(ctx echo.Context) error {
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

	err = c.bannerUsecase.DeleteBanner(ctx, idInt)
	if err != nil {
		log.Println(err)
		return helpers.EchoError(ctx, err)
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "deleted",
	})
}
