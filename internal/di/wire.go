//go:build wireinject
// +build wireinject

package di

import (
	"github.com/geekswamp/zen/internal/handler/v1/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserHandler() user.UserHandler {
	wire.Build(
		UserRepositorySet,
		UserServiceSet,
		UserHandlerSet,
	)

	return user.UserHandler{}
}

func ProvidePostgres() *gorm.DB {
	wire.Build(PostgresSet)

	return &gorm.DB{}
}
