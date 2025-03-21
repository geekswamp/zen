package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config defines the configuration options for CORS middleware.
type Config struct {
	AllowAllOrigins           bool
	AllowOrigins              []string
	AllowMethods              []string
	AllowHeaders              []string
	CustomSchemas             []string
	ExposeHeaders             []string
	AllowCredentials          bool
	MaxAge                    time.Duration
	AllowWildcard             bool
	AllowBrowserExtensions    bool
	AllowWebSockets           bool
	AllowFiles                bool
	OptionsResponseStatusCode int
}

// DefaultConfig returns a default CORS configuration with predefined settings.
func DefaultConfig() Config {
	return Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

// Default creates a handler function with default CORS configuration settings.
func Default() gin.HandlerFunc {
	config := DefaultConfig()
	return New(config)
}

// New creates a new CORS middleware handler with the specified configuration.
func New(conf Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:           conf.AllowAllOrigins,
		AllowOrigins:              conf.AllowOrigins,
		AllowMethods:              conf.AllowMethods,
		AllowHeaders:              conf.AllowHeaders,
		AllowCredentials:          conf.AllowCredentials,
		AllowWildcard:             conf.AllowWildcard,
		AllowBrowserExtensions:    conf.AllowBrowserExtensions,
		AllowWebSockets:           conf.AllowWebSockets,
		AllowFiles:                conf.AllowFiles,
		ExposeHeaders:             conf.ExposeHeaders,
		MaxAge:                    conf.MaxAge,
		OptionsResponseStatusCode: conf.OptionsResponseStatusCode,
		CustomSchemas:             conf.CustomSchemas,
	})
}
