package uberfx

import (
	"github.com/raufhm/learning-uberfx/internal/handler"
	"github.com/raufhm/learning-uberfx/internal/repository"
	"github.com/raufhm/learning-uberfx/internal/service"
	"go.uber.org/fx"
)

func provideHandler() fx.Option {
	return fx.Provide(handler.NewUserHandler)
}

func provideService() fx.Option {
	return fx.Provide(service.NewUserService)
}

func provideRepository() fx.Option {
	return fx.Provide(repository.NewUserRepository)
}
