package http

import "github.com/geekswamp/zen/internal/errors"

// Errno represents a custom error number type as a string.
type Errno string

// ErrorCode represents an error with a specific error code and detailed message.
type ErrorCode struct {
	code   Errno
	detail string
}

var errCodes = map[Errno]struct{}{}

// NewErrorCode creates a new error code with the given code and detail.
// It panics if the error code already exists in the global errCodes map.
func NewErrorCode(code Errno, detail string) *ErrorCode {
	if _, ok := errCodes[code]; ok {
		panic(errors.ErrInvalidErrCode)
	}

	errCodes[code] = struct{}{}

	return &ErrorCode{code: code, detail: detail}
}

// Code returns the error code number (Errno) associated with this ErrorCode.
func (e *ErrorCode) Code() Errno {
	return e.code
}

// Detail returns the error detail message. The detail provides specific information about the error
// that occurred, offering more context than the error code alone.
func (e *ErrorCode) Detail() string {
	return e.detail
}
