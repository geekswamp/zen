package di

import (
	"github.com/geekswamp/zen/internal/handler/v1/user"
	"github.com/geekswamp/zen/internal/http"
	"github.com/google/wire"
)

var UserHandlerSet = wire.NewSet(
	http.New,
	user.New,
)
