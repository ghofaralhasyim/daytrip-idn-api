package usecases

import (
	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/labstack/echo/v4"
)

type (
	InvitationUsecase interface {
	}

	invitationUsecase struct {
		invitationRepository repositories.InvitationRepository
		imageStorage         ImageStorageUsecase
	}
)

func NewInvitationUsecase(
	invitationRepository repositories.InvitationRepository,
	imageStorage ImageStorageUsecase,
) InvitationUsecase {
	return &invitationUsecase{
		invitationRepository: invitationRepository,
		imageStorage:         imageStorage,
	}
}

func (u *invitationUsecase) CreateInvitation(ctx echo.Context, e entities.InvitationEntity) (int64, error) {
	if e.Slug == "" {
		e.Slug = helpers.GenerateSlug(e.Title)
	}

	var assets []entities.InvitationAssetEntity
	for _, asset := range e.Assets {
		publicURL, err := u.imageStorage.Save(&asset.FileHeader)
		if err != nil {
			return 0, err
		}
		asset.AssetUrl = publicURL
		assets = append(assets, asset)
	}

	return u.invitationRepository.Create(ctx.Request().Context(), &e, assets)
}
