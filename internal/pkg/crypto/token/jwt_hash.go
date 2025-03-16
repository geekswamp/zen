package token

import (
	"time"

	"github.com/geekswamp/zen/internal/logger"
	"github.com/geekswamp/zen/internal/pkg/crypto/key"
	"github.com/geekswamp/zen/internal/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTHash struct {
	jwt.RegisteredClaims
	key.RSAKeyProvider
}

func New(iss, sub string, aud jwt.ClaimStrings, exp time.Duration, provider key.RSAKeyProvider) JWTProvider {
	now := time.Now()

	return &JWTHash{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: jwt.NewNumericDate(now.Add(exp)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
		RSAKeyProvider: provider,
	}
}

func (j *JWTHash) Generate() (hash string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &j.RegisteredClaims)

	hash, err = token.SignedString(j.GetPrivateKey())
	if err != nil {
		log.Error(errors.ErrFailedToSignToken.Error(), logger.ErrDetails(err))
		return "", errors.ErrFailedToSignToken
	}

	return hash, nil
}

func (j *JWTHash) Verify(tokenStr string) (claims *jwt.RegisteredClaims, err error) {
	pubKey := j.GetPublicKey()
	if pubKey == nil {
		log.Error(errors.ErrNilPubKey.Error())
		return nil, errors.ErrInvalidToken
	}

	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Error(errors.ErrFailedToSignToken.Error(), logger.ErrDetails(err))
			return nil, errors.ErrFailedToSignToken
		}
		return pubKey, nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, err
		}

		log.Error(errors.ErrFailedTokenParsing.Error(), logger.ErrDetails(err))
		return nil, errors.ErrInvalidToken
	}

	if !token.Valid {
		log.Error(errors.ErrInvalidToken.Error())
		return nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		log.Error(errors.ErrInvalidToken.Error(), logger.ErrDetails(err))
		return nil, errors.ErrInvalidToken
	}

	return claims, nil
}
