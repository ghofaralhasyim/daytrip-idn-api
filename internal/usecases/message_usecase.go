package usecases

import (
	"fmt"

	"github.com/daytrip-idn-api/internal/entities"
	error_app "github.com/daytrip-idn-api/internal/error"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

type (
	MessageUsecase interface {
		GetMessages(ctx echo.Context) ([]entities.MessageEntity, error)
		InsertMessage(ctx echo.Context, data entities.MessageEntity) (*entities.MessageEntity, error)
		DeleteMessage(ctx echo.Context, messageId int64) error
	}

	messageUsecase struct {
		messageRepository repositories.MessageRepository
	}
)

func NewMessageUsecase(
	messageRepository repositories.MessageRepository,
) MessageUsecase {
	return &messageUsecase{
		messageRepository: messageRepository,
	}
}

func (u *messageUsecase) GetMessages(
	ctx echo.Context,
) ([]entities.MessageEntity, error) {
	messages, err := u.messageRepository.GetMessages(ctx.Request().Context())
	if err != nil {
		fmt.Println("messageUsecase GetMessages err: ", err)
		return nil, error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	return messages, nil
}

func (u *messageUsecase) InsertMessage(
	ctx echo.Context, data entities.MessageEntity,
) (*entities.MessageEntity, error) {
	result, err := u.messageRepository.InsertMessage(
		ctx.Request().Context(), data,
	)
	if err != nil {
		fmt.Println("messageUsecase InsertMessage err: ", err)
		return nil, error_app.NewAppError(error_app.RepositorySaveError, err)
	}

	return &result, nil
}

func (u *messageUsecase) DeleteMessage(
	ctx echo.Context, messageId int64,
) error {
	err := u.messageRepository.DeleteMessage(ctx.Request().Context(), messageId)
	if err != nil {
		fmt.Println("messageUsecase InsertMessage err: ", err)
		return error_app.NewAppError(error_app.RepositoryDeleteError, err)
	}

	return nil
}
