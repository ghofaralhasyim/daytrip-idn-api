package usecases

import (
	"strings"

	"github.com/daytrip-idn-api/internal/entities"
	app_error "github.com/daytrip-idn-api/internal/error/app"
	error_app "github.com/daytrip-idn-api/internal/error/app"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type (
	BannerUsecase interface {
		GetBanners(ctx echo.Context) ([]entities.BannerEntity, error)
		InsertBanner(ctx echo.Context, e entities.BannerEntity) (*entities.BannerEntity, error)
		UpdateBanner(ctx echo.Context, e entities.BannerEntity) (*entities.BannerEntity, error)
		DeleteBanner(ctx echo.Context, id int) error
	}

	bannerUsecase struct {
		bannerRepository repositories.BannerRepository
		imageStorage     ImageStorageUsecase
	}
)

func NewBannerUsecase(
	bannerRepository repositories.BannerRepository,
	imageStorage ImageStorageUsecase,
) BannerUsecase {
	return &bannerUsecase{
		bannerRepository: bannerRepository,
		imageStorage:     imageStorage,
	}
}

func (u *bannerUsecase) GetBanners(ctx echo.Context) ([]entities.BannerEntity, error) {
	banner, err := u.bannerRepository.GetBanners(ctx.Request().Context())
	if err != nil {
		return nil, app_error.NewAppError(app_error.RepositoryGetError, err)
	}

	return banner, nil
}

func (u *bannerUsecase) InsertBanner(ctx echo.Context, e entities.BannerEntity) (*entities.BannerEntity, error) {
	if e.DesktopImageFile != nil {
		path, err := u.imageStorage.Save("/banners/", e.DesktopImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}

		e.DesktopImage = path
	}

	if e.MobileImageFile != nil {
		path, err := u.imageStorage.Save("/banners/", e.MobileImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}

		e.MobileImage = path
	}

	data, err := u.bannerRepository.InsertBanner(ctx.Request().Context(), &e)
	if err != nil {
		return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
	}

	return &data, nil
}

func (u *bannerUsecase) UpdateBanner(ctx echo.Context, e entities.BannerEntity) (*entities.BannerEntity, error) {
	data, err := u.bannerRepository.GetBannerById(ctx.Request().Context(), int(e.Id))
	if err != nil {
		return nil, error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	if e.DesktopImageFile != nil {
		path, err := u.imageStorage.Save("/banners/", e.DesktopImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}
		err = u.imageStorage.Delete(data.DesktopImage)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return nil, error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
			err = nil
		}

		e.DesktopImage = path
	} else {
		e.DesktopImage = data.DesktopImage
	}

	if e.MobileImageFile != nil {
		path, err := u.imageStorage.Save("/banners/", e.MobileImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}
		err = u.imageStorage.Delete(data.MobileImage)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return nil, error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
			err = nil
		}

		e.MobileImage = path
	} else {
		e.MobileImage = data.MobileImage
	}

	updatedData, err := u.bannerRepository.UpdateBanner(ctx.Request().Context(), &e)
	if err != nil {
		return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
	}

	return &updatedData, nil
}

func (u *bannerUsecase) DeleteBanner(ctx echo.Context, id int) error {
	data, err := u.bannerRepository.GetBannerById(ctx.Request().Context(), int(id))
	if err != nil {
		return error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	if data.DesktopImage != "" {
		err = u.imageStorage.Delete(data.DesktopImage)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
		}
	}

	if data.MobileImage != "" {
		err = u.imageStorage.Delete(data.MobileImage)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
		}
	}

	err = u.bannerRepository.DeleteBanner(ctx.Request().Context(), id)
	if err != nil {
		return error_app.NewAppError(error_app.RepositoryDeleteError, err)
	}

	return nil
}
