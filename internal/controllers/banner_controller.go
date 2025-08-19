package controllers

import (
	"net/http"

	response_transform "github.com/daytrip-idn-api/internal/rest/transform"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
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
