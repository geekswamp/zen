package key

import (
	"crypto/rsa"
	"os"

	"github.com/geekswamp/zen/configs"
	"github.com/golang-jwt/jwt/v5"
)

type RSAKeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func New() (RSAKeyProvider, error) {
	cfg := configs.Get()

	privateKey, err := loadPrivateKey(cfg)
	if err != nil {
		return nil, err
	}

	publicKey, err := loadPublicKey(cfg)
	if err != nil {
		return nil, err
	}

	return &RSAKeyPair{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (r *RSAKeyPair) GetPrivateKey() *rsa.PrivateKey {
	return r.privateKey
}

func (r *RSAKeyPair) GetPublicKey() *rsa.PublicKey {
	return r.publicKey
}

func loadPrivateKey(config configs.Config) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(config.JWT.PrivKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKey(config configs.Config) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(config.JWT.PubKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
