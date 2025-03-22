package server

import (
	"context"
	"net/http"
	"time"

	"github.com/geekswamp/zen/internal/errors"
	"github.com/gin-gonic/gin"
)

type Router func(engine *gin.Engine)

type Server interface {
	Start() error
	Stop() error
}

type Config struct {
	*http.Server
	mode        string
	middlewares []gin.HandlerFunc
	router      Router
}

func New(addr, mode string, router Router) Server {
	config := &Config{
		Server: &http.Server{Addr: addr},
		mode:   mode,
		router: router,
	}

	config.Handler = config.handler()

	return config
}

func (c *Config) Start() error {
	err := c.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		return err
	}

	return nil
}

func (c *Config) Stop() error {
	ctx, stop := context.WithTimeout(context.Background(), time.Second*3)
	defer stop()

	return c.Server.Shutdown(ctx)
}

func (c *Config) handler() *gin.Engine {
	server := gin.New()

	switch c.mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		panic(errors.ErrInvalidMode)
	}

	server.Use(c.middlewares...)
	c.router(server)

	return server
}

func (c *Config) ReadTimeout(timeout time.Duration) {
	c.Server.ReadTimeout = timeout
}

func (c *Config) WriteTimeout(timeout time.Duration) {
	c.Server.WriteTimeout = timeout
}

func (c *Config) Middleware(middlewares ...gin.HandlerFunc) {
	c.middlewares = append(c.middlewares, middlewares...)
}
