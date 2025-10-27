package usecases

import (
	"fmt"
	"log"
	"strings"

	"github.com/daytrip-idn-api/internal/entities"
	error_app "github.com/daytrip-idn-api/internal/error/app"
	error_data "github.com/daytrip-idn-api/internal/error/data"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
	"github.com/labstack/echo/v4"
)

type (
	InvitationUsecase interface {
		CreateInvitation(ctx echo.Context, e entities.InvitationEntity) (*entities.InvitationEntity, error)
		GetInvitations(ctx echo.Context) ([]entities.InvitationEntity, error)
		GetInvitationBySlug(ctx echo.Context, slug string) (*entities.InvitationEntity, error)
		UpdateInvitation(ctx echo.Context, e entities.InvitationEntity) (*entities.InvitationEntity, error)
		DeleteInvitation(ctx echo.Context, id int) error
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

func (u *invitationUsecase) CreateInvitation(ctx echo.Context, e entities.InvitationEntity) (*entities.InvitationEntity, error) {
	if e.Slug == "" {
		e.Slug = helpers.GenerateSlug(e.Title)
	}

	ctr := 1
	for {
		data, err := u.invitationRepository.GetBySlug(ctx.Request().Context(), e.slug)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositoryGetError, err)
		}
		if data == nil {
			break
		}

		ctr++
		e.Slug += fmt.Sprintf("%s-%d", e.Slug, ctr)
	}

	if e.TemplateId == 2 {
		e.Slug += "-birthday-party"
	}

	if e.ImageFile != nil {
		path, err := u.imageStorage.Save("/invitations/", e.ImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}

		e.Image = &path
	}

	if e.ImageFile1 != nil {
		path, err := u.imageStorage.Save("/invitations/", e.ImageFile1)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}

		e.Image1 = &path
	}

	id, err := u.invitationRepository.Create(ctx.Request().Context(), &e)
	if err != nil {
		return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
	}

	e.Id = id

	return &e, nil
}

func (u *invitationUsecase) UpdateInvitation(ctx echo.Context, e entities.InvitationEntity) (*entities.InvitationEntity, error) {
	if e.Slug == "" {
		return nil, error_data.NewAppError(error_data.InvalidDataRequest, fmt.Errorf("slug is required"))
	}

	data, err := u.invitationRepository.GetBySlug(ctx.Request().Context(), e.Slug)
	if err != nil || data == nil {
		return nil, error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	if e.Title != data.Title {
		e.Slug = helpers.GenerateSlug(e.Title)
	}

	if e.ImageFile != nil {
		path, err := u.imageStorage.Save("/invitations/", e.ImageFile)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}
		err = u.imageStorage.Delete(*data.Image)
		if err != nil {
			log.Println(err.Error())
		}
		e.Image = &path
	} else {
		e.Image = data.Image
	}

	if e.ImageFile1 != nil {
		path, err := u.imageStorage.Save("/invitations/", e.ImageFile1)
		if err != nil {
			return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
		}
		err = u.imageStorage.Delete(*data.Image1)
		if err != nil {
			if err != nil {
				log.Println(err.Error())
			}
			err = nil
		}
		e.Image1 = &path
	} else {
		e.Image1 = data.Image1
	}

	e.Id = data.Id

	if e.BirthdayVal == nil {
		var val int64
		e.BirthdayVal = &val
	}

	err = u.invitationRepository.Update(ctx.Request().Context(), &e)
	if err != nil {
		return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
	}

	return &e, nil
}

func (u *invitationUsecase) DeleteInvitation(ctx echo.Context, id int) error {
	data, err := u.invitationRepository.GetById(ctx.Request().Context(), id)
	if err != nil {
		return error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	tx, err := u.invitationRepository.BeginTx()
	if err != nil {
		return error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Println("Recovered from panic: %v", p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = u.invitationRepository.DeleteResponseInvitation(ctx.Request().Context(), tx, id)
	if err != nil {
		log.Println(err)
		return error_app.NewAppError(error_app.RepositoryDeleteError, err)
	}

	err = u.invitationRepository.DeleteInvitation(ctx.Request().Context(), tx, id)
	if err != nil {
		log.Println(err)
		return error_app.NewAppError(error_app.RepositoryDeleteError, err)
	}

	if data.Image != nil && *data.Image != "" {
		err = u.imageStorage.Delete(*data.Image)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
			err = nil
		}
	}

	if data.Image1 != nil && *data.Image1 != "" {
		err = u.imageStorage.Delete(*data.Image1)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				return error_app.NewAppError(error_app.RepositoryDeleteError, err)
			}
			err = nil
		}
	}

	return nil
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
