package token

import (
	"github.com/geekswamp/zen/internal/logger"
	"github.com/golang-jwt/jwt/v5"
)

var log = logger.New()

type JWTProvider interface {
	Generate() (hash string, err error)
	Verify(tokenStr string) (claims *jwt.RegisteredClaims, err error)
}
