//go:build wireinject
// +build wireinject

// この2つがないとパッケージ内で競合する

package di

import (
	"backend-app/controller"
	"backend-app/repository"
	"backend-app/service"
	"database/sql"

	"github.com/google/wire"
)

// repository
var repositorySet = wire.NewSet(
	repository.NewUserRepository,
)

// service
var serviceSet = wire.NewSet(
	service.NewUserService,
)

// controller
var controllerSet = wire.NewSet(
	controller.NewUserController,
)

type ControllersSet struct {
	IUserController controller.IUserController
}

func InitializeControllers(db *sql.DB) (*ControllersSet, error) {
	wire.Build(
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
