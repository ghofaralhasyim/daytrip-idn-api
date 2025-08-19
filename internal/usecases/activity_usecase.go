package usecases

import (
	"github.com/daytrip-idn-api/internal/entities"
	app_error "github.com/daytrip-idn-api/internal/error"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type (
	ActivityUsecase interface {
		GetActivities(ctx echo.Context) ([]entities.ActivityEntity, error)
	}

	activityUsecase struct {
		activityRepository repositories.ActivityRepository
	}
)

func NewActivityUsecase(
	activityRepository repositories.ActivityRepository,
) ActivityUsecase {
	return &activityUsecase{
		activityRepository: activityRepository,
	}
}

func (u *activityUsecase) GetActivities(ctx echo.Context) (
	[]entities.ActivityEntity, error,
) {
	activities, err := u.activityRepository.GetActivities(ctx.Request().Context())
	if err != nil {
		return nil, app_error.NewAppError(app_error.RepositoryGetError, err)
	}

	return activities, nil
}
