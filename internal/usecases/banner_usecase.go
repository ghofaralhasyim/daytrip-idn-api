package usecases

import (
	"github.com/daytrip-idn-api/internal/entities"
	app_error "github.com/daytrip-idn-api/internal/error/app"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type (
	BannerUsecase interface {
		GetBanners(ctx echo.Context) ([]entities.BannerEntity, error)
	}

	bannerUsecase struct {
		bannerRepository repositories.BannerRepository
	}
)

func NewBannerUsecase(
	bannerRepository repositories.BannerRepository,
) BannerUsecase {
	return &bannerUsecase{
		bannerRepository: bannerRepository,
	}
}

func (u *bannerUsecase) GetBanners(ctx echo.Context) ([]entities.BannerEntity, error) {
	banner, err := u.bannerRepository.GetBanners(ctx.Request().Context())
	if err != nil {
		return nil, app_error.NewAppError(app_error.RepositoryGetError, err)
	}

	return banner, nil
}
