package controllers

import (
	"net/http"

	response_transform "github.com/daytrip-idn-api/internal/rest/transform"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	activityUsecase usecases.ActivityUsecase
}

func NewActivityController(
	activityUsecase usecases.ActivityUsecase,
) *ActivityController {
	return &ActivityController{
		activityUsecase: activityUsecase,
	}
}

func (c *ActivityController) GetActivities(ctx echo.Context) error {
	activities, err := c.activityUsecase.GetActivities(ctx)
	if err != nil {
		return helpers.EchoError(ctx, err)
	}

	response := response_transform.TransformListActivityResponse(activities)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}
