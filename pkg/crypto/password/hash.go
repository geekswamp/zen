package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/geekswamp/zen/pkg/errors"
	"golang.org/x/crypto/argon2"
)

func (a *passHash) Generate(text []byte) (hash string, err error) {
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

func (a *passHash) Verify(text []byte, hash string) (bool, error) {
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

func decodeHash(encodedHash string) (ph *passHash, salt, b []byte, err error) {
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

	ph = &passHash{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &ph.memory, &ph.iterations, &ph.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	ph.saltLength = uint32(len(salt))

	b, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}

	ph.keyLength = uint32(len(b))

	return ph, salt, b, nil
}
