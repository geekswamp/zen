package key

import "crypto/rsa"

type RSAKeyProvider interface {
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
}
