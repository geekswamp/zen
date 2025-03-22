package errors

import "errors"

var (
	ErrNotAFile                  = errors.New("not a file")
	ErrFailedToConnectDB         = errors.New("failed to connect to database")
	ErrFailedLoadLocal           = errors.New("failed to load local timezone")
	ErrInvalidHashFormat         = errors.New("invalid hash format")
	ErrPasswordTooShort          = errors.New("password too short")
	ErrIncompatibleArgon2Version = errors.New("incompatible Argon2 version")
	ErrFailedToDecodeHash        = errors.New("failed to decode hash")
	ErrFailedToSignToken         = errors.New("failed to sign token")
	ErrFailedGenRandomBytes      = errors.New("failed to generate random bytes")
	ErrFailedToDecodeStr         = errors.New("failed to decode string")
	ErrInvalidToken              = errors.New("invalid token")
	ErrNilPubKey                 = errors.New("public key is nil")
	ErrFailedTokenParsing        = errors.New("failed parsing token")
	ErrFailedRunningSeeder       = errors.New("failed running seeder")
	ErrValidatorTrans            = errors.New("error validator translation")
	ErrInvalidErrCode            = errors.New("error code does not exist, please change one")
	ErrInvalidMode               = errors.New("the 'mode' only supports 'debug' and 'release'. Please update your config file accordingly")
)
