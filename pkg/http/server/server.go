package server

import (
	"context"
	"net/http"
	"time"

	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.New()

type Router func(engine *gin.Engine)

type Server interface {
	Start() error
	Stop() error
}

type Config struct {
	*http.Server
	mode        string
	middlewares []gin.HandlerFunc
	routerFunc  Router
}

type Option func(c *Config)

func New(addr string, opts ...Option) Server {
	config := &Config{
		Server: &http.Server{Addr: addr},
	}

	for _, opt := range opts {
		opt(config)
	}

	config.Handler = config.handler()

	return config
}

func (c *Config) Start() error {
	log.Info("Starting server", logger.Server(c.Server.Addr))

	if err := c.Server.ListenAndServe(); err == http.ErrServerClosed {
		return err
	}

	return nil
}

func (c *Config) Stop() error {
	log.Info("Stopping server", logger.Server(c.Server.Addr))

	ctx, stop := context.WithTimeout(context.Background(), time.Second*3)
	defer stop()

	return c.Server.Shutdown(ctx)
}

func (c *Config) handler() *gin.Engine {
	g := gin.New()

	switch c.mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		panic(errors.ErrInvalidMode)
	}

	g.Use(c.middlewares...)

	c.routerFunc(g)

	return g
}

func ReadTimeout(timout time.Duration) Option {
	return func(c *Config) { c.Server.ReadTimeout = timout }
}

func WriteTimeout(timout time.Duration) Option {
	return func(c *Config) { c.Server.WriteTimeout = timout }
}

func SetMode(mode string) Option {
	return func(c *Config) { c.mode = mode }
}

func Middlewares(middlewares ...gin.HandlerFunc) Option {
	return func(c *Config) { c.middlewares = append(c.middlewares, middlewares...) }
}

func RegisterRouter(router Router) Option {
	return func(c *Config) { c.routerFunc = router }
}
