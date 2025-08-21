package usecases

import (
	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/labstack/echo/v4"
)

type (
	InvitationUsecase interface {
		CreateInvitation(ctx echo.Context, e entities.InvitationEntity) (int64, error)
		GetInvitations(ctx echo.Context) ([]entities.InvitationEntity, error)
		GetInvitationBySlug(ctx echo.Context, slug string) (*entities.InvitationEntity, error)
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

	path, err := u.imageStorage.Save("/invitations/", e.ImageFile)
	if err != nil {
		return 0, err
	}
	e.Image = &path

	path, err = u.imageStorage.Save("/invitations/", e.ImageFile1)
	if err != nil {
		return 0, err
	}
	e.Image1 = &path

	return u.invitationRepository.Create(ctx.Request().Context(), &e)
}

func (u *invitationUsecase) GetInvitations(ctx echo.Context) (
	[]entities.InvitationEntity, error,
) {
	return u.invitationRepository.GetInvitations(ctx.Request().Context())
}

func (u *invitationUsecase) GetInvitationBySlug(ctx echo.Context, slug string) (
	*entities.InvitationEntity, error,
) {
	return u.invitationRepository.GetBySlug(ctx.Request().Context(), slug)
}
