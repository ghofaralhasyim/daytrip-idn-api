package usecases

import (
	"errors"
	"fmt"

	error_app "github.com/daytrip-idn-api/internal/error"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Authenticate(ctx echo.Context, email string, password string) (*models.User, string, error)
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (s *userUsecase) Authenticate(ctx echo.Context, email string, password string) (*models.User, string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx.Request().Context(), email)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "sql: no rows in result set" {
			return nil, "", error_app.NewAppError(error_app.InvalidCredentials, errors.New("invalid credentials"))
		}
		return nil, "", error_app.NewAppError(error_app.RepositoryGetError, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", error_app.NewAppError(error_app.InvalidCredentials, errors.New("invalid credentials"))
	}

	jwt, err := utils.GenerateSessionToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, jwt, nil
}
