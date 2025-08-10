package modules

import (
	"database/sql"

	"github.com/daytrip-idn-api/internal/controllers"
	"github.com/daytrip-idn-api/internal/repositories"
	"github.com/daytrip-idn-api/internal/usecases"
)

type Repositories struct {
	User        repositories.UserRepository
	Banner      repositories.BannerRepository
	Destination repositories.DestinationRepository
	Message     repositories.MessageRepository
}

type Usecases struct {
	User        usecases.UserUsecase
	Banner      usecases.BannerUsecase
	Destination usecases.DestinationUsecase
	Message     usecases.MessageUsecase
}

type Controllers struct {
	User        *controllers.UserController
	Banner      *controllers.BannerController
	Destination *controllers.DestinationController
	Message     *controllers.MessageController
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
		User:        repositories.NewUserRepository(db),
		Banner:      repositories.NewBannerRepository(db),
		Destination: repositories.NewDestinationRepository(db),
		Message:     repositories.NewMessageRepository(db),
	}

	modules.Usecases = Usecases{
		User:        usecases.NewUserUsecase(modules.Repositories.User),
		Banner:      usecases.NewBannerUsecase(modules.Repositories.Banner),
		Destination: usecases.NewDestinationUsecase(modules.Repositories.Destination),
		Message:     usecases.NewMessageUsecase(modules.Repositories.Message),
	}

	modules.Controllers = Controllers{
		User:        controllers.NewUserController(modules.Usecases.User),
		Banner:      controllers.NewBannerController(modules.Usecases.Banner),
		Destination: controllers.NewDestinationController(modules.Usecases.Destination),
		Message:     controllers.NewMessageController(modules.Usecases.Message),
	}

	return &modules
}
