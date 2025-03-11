package errors

import "errors"

var (
	ErrNotAFile          = errors.New("Not a file")
	ErrFailedToConnectDB = errors.New("Failed to connect to database")
	ErrFailedLoadLocal   = errors.New("Failed to load local timezone")
)
