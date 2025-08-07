package usecases

import (
	"errors"
	"strconv"

	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserById(userId string) (*models.User, error)
	Authenticate(email string, password string) (*models.User, string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUserById(id string) (*models.User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return s.userRepository.GetUserById(userId)
}

func (s *userService) Authenticate(email string, password string) (*models.User, string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, "", errors.New("user not found")
		}
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", errors.New("unauthorize")
	}

	jwt, err := utils.GenerateSessionToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, jwt, nil
}
