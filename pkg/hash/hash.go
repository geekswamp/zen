package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/geekswamp/zen/pkg/errors"
	"golang.org/x/crypto/argon2"
)

var _ Hash = (*hash)(nil)

type (
	Hash interface {
		Generate(text []byte) (hash string, err error)
		Verify(text []byte, hash string) (bool, error)
		t()
	}

	hash struct {
		pepper      string
		memory      uint32
		iterations  uint32
		parallelism uint8
		saltLength  uint32
		keyLength   uint32
	}
)

func New(pepper string, memory, iterations, saltLength, keyLength uint32, parallelism uint8) Hash {
	return &hash{
		pepper:      pepper,
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltLength:  saltLength,
		keyLength:   keyLength,
	}
}

func (a *hash) t() {}

func (a *hash) Generate(text []byte) (hash string, err error) {
	pepperedText := append(text, []byte(a.pepper)...)

	salt, err := generateRandomBytes(a.saltLength)
	if err != nil {
		return "", err
	}

	hasher := argon2.IDKey(pepperedText, salt, a.iterations, a.memory, a.parallelism, a.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hasher)

	hash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.iterations, a.parallelism, b64Salt, b64Hash)

	return hash, nil
}

func (a *hash) Verify(text []byte, hash string) (bool, error) {
	pepperedText := append(text, []byte(a.pepper)...)

	h, salt, hashed, err := decodeHash(hash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey(pepperedText, salt, h.iterations, h.memory, h.parallelism, h.keyLength)
	if subtle.ConstantTimeCompare(hashed, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func decodeHash(encodedHash string) (h *hash, salt, b []byte, err error) {
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

	h = &hash{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &h.memory, &h.iterations, &h.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	h.saltLength = uint32(len(salt))

	b, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	h.keyLength = uint32(len(b))

	return h, salt, b, nil
}
