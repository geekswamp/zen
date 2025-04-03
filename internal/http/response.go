package http

import (
	"io"
	"net/http"

	"github.com/geekswamp/zen/internal/core"
	"github.com/gin-gonic/gin"
)

// BaseResponse serves as a foundational structure for API responses.
type BaseResponse struct{}

// Response represents a standardized API response structure.
type Response struct {
	RequestID string `json:"request_id"`
	Error     *Error `json:"error"`
	Result    any    `json:"result"`
}

// Error represents a standard error response structure.
type Error struct {
	Code   Errno  `json:"code"`
	Reason string `json:"reason"`
}

// Entries represents a paginated collection of items of type T.
type Entries[T any] struct {
	Entries       []T   `json:"entries"`
	TotalItems    int64 `json:"total_items"`
	TotalPages    int64 `json:"total_pages"`
	HasReachedMax bool  `json:"has_reached_max"`
}

// New creates and returns a new empty instance of BaseResponse.
// This function serves as a constructor for initializing a BaseResponse object
// with its zero values.
func New() BaseResponse {
	return BaseResponse{}
}

// NewJSON creates and sends a JSON response using the provided gin.Context.
func NewJSON(c *gin.Context, httpCode int, message, err *Error, data any) {
	newResponse(c, httpCode, err, data)
}

// NewEntries creates a new Entries instance with pagination metadata.
func NewEntries[T any](entries []T, totalItems, totalPages int64, hasReachedMax bool) Entries[T] {
	return Entries[T]{
		Entries:       entries,
		TotalItems:    totalItems,
		TotalPages:    totalPages,
		HasReachedMax: hasReachedMax,
	}
}

// Created sends a JSON response with HTTP 201 Created status code.
func (b BaseResponse) Created(c *gin.Context, data any) {
	newResponse(c, http.StatusCreated, nil, data)
}

// Success sends a JSON response with HTTP 200 Ok status code.
func (b BaseResponse) Success(c *gin.Context, data any) {
	newResponse(c, http.StatusOK, nil, data)
}

// BadRequest sends a JSON response with HTTP 400 Bad Request status code.
func (b BaseResponse) BadRequest(c *gin.Context, err Error) {
	newResponse(c, http.StatusBadRequest, &err, nil)
}

// Unauthorized sends a JSON response with HTTP 401 Unauthorized status code.
func (b BaseResponse) Unauthorized(c *gin.Context, err Error) {
	newResponse(c, http.StatusUnauthorized, &err, nil)
}

// TMR sends a JSON response with HTTP 429 Too Many Requests status code.
func (b BaseResponse) TMR(c *gin.Context) {
	newResponse(c, http.StatusTooManyRequests, &Error{Code: TooManyReqs.Code(), Reason: TooManyReqs.Detail()}, nil)
}

// ISE sends a JSON response with HTTP 500 Internal Server Error status code.
func (b BaseResponse) ISE(c *gin.Context, err *Error) {
	newResponse(c, http.StatusInternalServerError, err, nil)
}

// Error handles and formats error responses for HTTP requests.
// It accepts a gin.Context and any error parameter, processing different error types:
//   - For custom Error type: Responds with BadRequest
//   - For standard error type: Processes specific cases like io.EOF with appropriate status codes
func (b BaseResponse) Error(c *gin.Context, errParam any) {
	switch err := errParam.(type) {
	case Error:
		b.BadRequest(c, err)

	case error:
		var code Errno
		var msg string
		var httpCode int

		switch {
		case err == io.EOF:
			code = NotValidJSONFormat.Code()
			msg = NotValidJSONFormat.Detail()
			httpCode = http.StatusBadRequest

		default:
			code = SystemError.Code()
			msg = SystemError.Detail()
			httpCode = http.StatusInternalServerError
		}

		newResponse(c, httpCode, &Error{Code: code, Reason: msg}, nil)
	}
}

func newResponse(c *gin.Context, code int, err *Error, data any) {
	ctx := core.NewContext(c)

	c.JSON(code, Response{RequestID: *ctx.GetRequestID(), Error: err, Result: data})
}
