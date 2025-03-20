package password

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/crypto/rand"
	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/logger"
	"golang.org/x/crypto/argon2"
)

var log = logger.New()

type Raw struct {
	Config Config
	Salt   []byte
	Hash   []byte
}

type Config struct {
	pepper      string
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func New(pepper string, memory, iterations, saltLength, keyLength uint32, parallelism uint8) Config {
	return Config{
		pepper:      pepper,
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltLength:  saltLength,
		keyLength:   keyLength,
	}
}

func NewFromConfig() Config {
	cfg := configs.Get().Password

	return Config{
		pepper:      cfg.Pepper,
		memory:      cfg.Argon2.Memory,
		iterations:  cfg.Argon2.Iterations,
		parallelism: cfg.Argon2.Parallelism,
		saltLength:  cfg.Argon2.SaltLength,
		keyLength:   cfg.Argon2.KeyLength,
	}
}

func (a *Config) Hash(text, salt []byte) (*Raw, error) {
	pepperedText := append(text, []byte(a.pepper)...)

	if text == nil {
		return nil, errors.ErrPasswordTooShort
	}

	if salt == nil {
		var err error
		salt, err = rand.GenerateRandomBytes(a.saltLength)
		if err != nil {
			return nil, err
		}
	}

	hash := argon2.IDKey(pepperedText, salt, a.iterations, a.memory, a.parallelism, a.keyLength)

	return &Raw{
		Config: *a,
		Salt:   salt,
		Hash:   hash,
	}, nil
}

func (a *Config) Generate(text []byte) (hash string, err error) {
	raw, err := a.Hash(text, nil)
	if err != nil {
		log.Error(errors.ErrFailedGenRandomBytes.Error(), logger.ErrDetails(err))
		return "", err
	}

	b64Salt := encodeToString(raw.Salt)
	b64Hash := encodeToString(raw.Hash)

	hash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.iterations, a.parallelism, b64Salt, b64Hash)

	return hash, nil
}

func (r *Raw) Verify(text []byte, hash string) (bool, error) {
	raw, err := r.Config.Hash(text, r.Salt)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(r.Hash, raw.Hash) == 1, nil
}

func encodeToString(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}
