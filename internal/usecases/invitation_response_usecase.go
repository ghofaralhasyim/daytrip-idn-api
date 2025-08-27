package usecases

import (
	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type InvitationResponseUsecase interface {
	SubmitResponse(ctx echo.Context, e entities.InvitationResponseEntity) (int64, error)
	GetInvitationResponse(ctx echo.Context) ([]entities.InvitationResponseEntity, error)
	GetInvitationResponseBySlug(ctx echo.Context, slug string) ([]entities.InvitationResponseEntity, error)
}

type invitationResponseUsecase struct {
	responseRepo repositories.InvitationResponseRepository
}

func NewInvitationResponseUsecase(
	responseRepo repositories.InvitationResponseRepository,
) InvitationResponseUsecase {
	return &invitationResponseUsecase{
		responseRepo: responseRepo,
	}
}

func (u *invitationResponseUsecase) SubmitResponse(
	ctx echo.Context,
	e entities.InvitationResponseEntity,
) (int64, error) {
	return u.responseRepo.Create(ctx.Request().Context(), &e)
}

func (u *invitationResponseUsecase) GetInvitationResponse(ctx echo.Context) ([]entities.InvitationResponseEntity, error) {
	return u.responseRepo.GetInvitationResponse(ctx.Request().Context())
}

func (u *invitationResponseUsecase) GetInvitationResponseBySlug(ctx echo.Context, slug string) ([]entities.InvitationResponseEntity, error) {
	return u.responseRepo.GetInvitationResponseBySlug(ctx.Request().Context(), slug)
}
