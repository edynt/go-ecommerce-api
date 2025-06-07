//go:build wireinject

package wire

import (
	"github.com/edynt/go-ecommerce-api/internal/controllers"
	"github.com/edynt/go-ecommerce-api/internal/repo"
	"github.com/edynt/go-ecommerce-api/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		services.NewUserService,
		controllers.NewUserController,
	)

	return new(controllers.UserController), nil
}
