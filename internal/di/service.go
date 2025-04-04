package di

import (
	"github.com/geekswamp/zen/internal/service"
	"github.com/google/wire"
)

var UserServiceSet = wire.NewSet(service.NewUserService)
