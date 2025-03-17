package password

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/geekswamp/zen/internal/pkg/crypto/rand"
	"github.com/geekswamp/zen/internal/pkg/errors"
	"golang.org/x/crypto/argon2"
)

type Config struct {
	pepper      string
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func New(pepper string, memory, iterations, saltLength, keyLength uint32, parallelism uint8) Hash {
	return &Config{
		pepper:      pepper,
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltLength:  saltLength,
		keyLength:   keyLength,
	}
}

func NewDefault() Hash {
	cfg := configs.Get().Password

	return &Config{
		pepper:      cfg.Pepper,
		memory:      cfg.Argon2.Memory,
		iterations:  cfg.Argon2.Iterations,
		parallelism: cfg.Argon2.Parallelism,
		saltLength:  cfg.Argon2.SaltLength,
		keyLength:   cfg.Argon2.KeyLength,
	}
}

func (a *Config) Generate(text []byte) (hash string, err error) {
	pepperedText := append(text, []byte(a.pepper)...)

	salt, err := rand.GenerateRandomBytes(a.saltLength)
	if err != nil {
		log.Error(errors.ErrFailedGenRandomBytes.Error(), logger.ErrDetails(err))
		return "", errors.ErrFailedGenRandomBytes
	}

	hasher := argon2.IDKey(pepperedText, salt, a.iterations, a.memory, a.parallelism, a.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hasher)

	hash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.iterations, a.parallelism, b64Salt, b64Hash)

	return hash, nil
}

func (a *Config) Verify(text []byte, hash string) (bool, error) {
	pepperedText := append(text, []byte(a.pepper)...)

	h, salt, hashed, err := decodeHash(hash)
	if err != nil {
		log.Error(err.Error(), logger.ErrDetails(err))
		return false, errors.ErrFailedToDecodeHash
	}

	otherHash := argon2.IDKey(pepperedText, salt, h.iterations, h.memory, h.parallelism, h.keyLength)
	if subtle.ConstantTimeCompare(hashed, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (ph *Config, salt, b []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.ErrInvalidHashFormat
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.ErrIncompatibleArgon2Version
	}

	ph = &Config{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &ph.memory, &ph.iterations, &ph.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, errors.ErrFailedToDecodeStr
	}

	ph.saltLength = uint32(len(salt))

	b, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, errors.ErrFailedToDecodeStr
	}

	ph.keyLength = uint32(len(b))

	return ph, salt, b, nil
}
