package password

import "github.com/geekswamp/zen/internal/logger"

var log = logger.New()

type Hash interface {
	Generate(text []byte) (hash string, err error)
	Verify(text []byte, hash string) (bool, error)
}
