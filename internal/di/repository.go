package di

import (
	"github.com/geekswamp/zen/internal/base"
	"github.com/geekswamp/zen/internal/repository"
	"github.com/google/wire"
)

var UserRepositorySet = wire.NewSet(
	PostgresSet,
	base.NewRepo,
	repository.NewUserRepo,
)
