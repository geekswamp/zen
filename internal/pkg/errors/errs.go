package errors

import "errors"

var (
	ErrNotAFile                  = errors.New("not a file")
	ErrFailedToConnectDB         = errors.New("failed to connect to database")
	ErrFailedLoadLocal           = errors.New("failed to load local timezone")
	ErrInvalidHashFormat         = errors.New("invalid hash format")
	ErrIncompatibleArgon2Version = errors.New("incompatible Argon2 version")
	ErrFailedToDecodeHash        = errors.New("failed to decode hash")
	ErrFailedToSignToken         = errors.New("failed to sign token")
	ErrFailedGenRandomBytes      = errors.New("failed to generate random bytes")
	ErrFailedToDecodeStr         = errors.New("failed to decode string")
	ErrInvalidToken              = errors.New("invalid token")
	ErrNilPubKey                 = errors.New("public key is nil")
	ErrFailedTokenParsing        = errors.New("failed parsing token")
)
