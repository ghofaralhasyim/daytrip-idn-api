package usecases

import (
	"fmt"

	"github.com/daytrip-idn-api/internal/entities"
	app_error "github.com/daytrip-idn-api/internal/error/app"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type (
	DestinationUsecase interface {
		GetDestinations(ctx echo.Context) ([]entities.DestinationEntity, error)
	}

	destinationUsecase struct {
		destinationRepository repositories.DestinationRepository
	}
)

func NewDestinationUsecase(
	destinationRepository repositories.DestinationRepository,
) DestinationUsecase {
	return &destinationUsecase{
		destinationRepository: destinationRepository,
	}
}

func (u *destinationUsecase) GetDestinations(
	ctx echo.Context,
) ([]entities.DestinationEntity, error) {
	destinations, err := u.destinationRepository.GetDestinations(ctx.Request().Context())
	if err != nil {
		fmt.Println("destinationUsecase GetDestinations err:", err)
		return nil, app_error.NewAppError(app_error.RepositoryGetError, err)
	}

	return destinations, nil
}
