package http

import "github.com/geekswamp/zen/internal/errors"

type Errno string

type ErrorCode struct {
	code   Errno
	detail string
}

var errCodes = map[Errno]struct{}{}

func NewErrorCode(code Errno, detail string) *ErrorCode {
	if _, ok := errCodes[code]; ok {
		panic(errors.ErrInvalidErrCode)
	}

	errCodes[code] = struct{}{}

	return &ErrorCode{code: code, detail: detail}
}

func (e *ErrorCode) Code() Errno {
	return e.code
}

func (e *ErrorCode) Detail() string {
	return e.detail
}
