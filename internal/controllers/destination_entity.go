package controllers

import (
	"net/http"

	response_transform "github.com/daytrip-idn-api/internal/rest/transform"
	"github.com/daytrip-idn-api/internal/usecases"
	"github.com/daytrip-idn-api/pkg/utils"
	"github.com/labstack/echo/v4"
)

type DestinationController struct {
	destinationUsecase usecases.DestinationUsecase
}

func NewDestinationController(
	destinationUsecase usecases.DestinationUsecase,
) *DestinationController {
	return &DestinationController{
		destinationUsecase: destinationUsecase,
	}
}

func (c *DestinationController) GetDestinations(ctx echo.Context) error {
	destination, err := c.destinationUsecase.GetDestinations(ctx)
	if err != nil {
		return utils.EchoError(ctx, err)
	}

	response := response_transform.TransformListDestinationResponse(destination)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"data":    response,
	})
}
