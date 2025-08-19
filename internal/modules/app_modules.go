package modules

import (
	"database/sql"

	"github.com/daytrip-idn-api/internal/controllers"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/internal/usecases"
)

type Repositories struct {
	User               repositories.UserRepository
	Banner             repositories.BannerRepository
	Destination        repositories.DestinationRepository
	Message            repositories.MessageRepository
	Activity           repositories.ActivityRepository
	InvitationResponse repositories.InvitationResponseRepository
}

type Usecases struct {
	User               usecases.UserUsecase
	Banner             usecases.BannerUsecase
	Destination        usecases.DestinationUsecase
	Message            usecases.MessageUsecase
	Activity           usecases.ActivityUsecase
	InvitationResponse usecases.InvitationResponseUsecase
}

type Controllers struct {
	User        *controllers.UserController
	Banner      *controllers.BannerController
	Destination *controllers.DestinationController
	Message     *controllers.MessageController
	Activity    *controllers.ActivityController
	Invitation  *controllers.InvitationController
}

type AppModules struct {
	Repositories
	Usecases
	Controllers
}

func NewAppModules(
	db *sql.DB,
) *AppModules {
	modules := AppModules{}

	modules.Repositories = Repositories{
		User:               repositories.NewUserRepository(db),
		Banner:             repositories.NewBannerRepository(db),
		Destination:        repositories.NewDestinationRepository(db),
		Message:            repositories.NewMessageRepository(db),
		Activity:           repositories.NewActivityRepository(db),
		InvitationResponse: repositories.NewInvitationResponseRepository(db),
	}

	modules.Usecases = Usecases{
		User:               usecases.NewUserUsecase(modules.Repositories.User),
		Banner:             usecases.NewBannerUsecase(modules.Repositories.Banner),
		Destination:        usecases.NewDestinationUsecase(modules.Repositories.Destination),
		Message:            usecases.NewMessageUsecase(modules.Repositories.Message),
		Activity:           usecases.NewActivityUsecase(modules.Repositories.Activity),
		InvitationResponse: usecases.NewInvitationResponseUsecase(modules.Repositories.InvitationResponse),
	}

	modules.Controllers = Controllers{
		User:        controllers.NewUserController(modules.Usecases.User),
		Banner:      controllers.NewBannerController(modules.Usecases.Banner),
		Destination: controllers.NewDestinationController(modules.Usecases.Destination),
		Message:     controllers.NewMessageController(modules.Usecases.Message),
		Activity:    controllers.NewActivityController(modules.Usecases.Activity),
		Invitation:  controllers.NewInvitationController(modules.Usecases.InvitationResponse),
	}

	return &modules
}
