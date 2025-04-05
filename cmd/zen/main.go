package main

import (
	"fmt"
	"time"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/di"
	"github.com/geekswamp/zen/internal/router"
	"github.com/geekswamp/zen/internal/storage/seed"
	"github.com/geekswamp/zen/pkg/http/middleware"
	"github.com/geekswamp/zen/pkg/http/middleware/cors"
	"github.com/geekswamp/zen/pkg/http/server"
)

func main() {
	c := configs.Get()
	s := server.New(
		fmt.Sprintf("%s:%d", c.App.Host, c.App.Port),
		server.SetMode(c.App.Mode),
		server.Middlewares(cors.Default(), middleware.RequestID()),
		server.ReadTimeout(30*time.Second),
		server.WriteTimeout(30*time.Second),
		server.RegisterRouter(router.RegisterRouter),
	)

	if err := seed.RunSeeders(di.ProvidePostgres()); err != nil {
		panic(err)
	}

	s.Start()
}
