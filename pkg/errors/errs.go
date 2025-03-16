package errors

import "errors"

var (
	ErrNotAFile                  = errors.New("Not a file")
	ErrFailedToConnectDB         = errors.New("Failed to connect to database")
	ErrFailedLoadLocal           = errors.New("Failed to load local timezone")
	ErrInvalidHashFormat         = errors.New("Invalid hash format")
	ErrIncompatibleArgon2Version = errors.New("Incompatible Argon2 version")
	ErrFailedToDecodeHash        = errors.New("Failed to decode hash")
	ErrFailedToSignToken         = errors.New("Failed to sign token")
	ErrFailedGenRandomBytes      = errors.New("Failed to generate random bytes")
	ErrFailedToDecodeStr         = errors.New("Failed to decode string")
)
