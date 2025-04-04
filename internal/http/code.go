package http

var (
	TooManyReqs        = NewErrorCode("REQ-000001", "Too many requests, please try again later")
	NotValidJSONFormat = NewErrorCode("ERR-PA40001", "Payload not valid JSON format")
	InputNotValid      = NewErrorCode("ERR-PA40002", "The provided input is not valid")
	Forbidden          = NewErrorCode("ERR-PA40003", "Access to this resource is forbidden")
	NotValidQuery      = NewErrorCode("ERR-PA40004", "The provided URL Query is not valid")
	UserAlreadyExists  = NewErrorCode("ERR-PA40005", "User already exists. Please use a different email or phone number")
	NotFound           = NewErrorCode("ERR-PA40006", "The requested resource was not found")
	InvalidRequestID   = NewErrorCode("ERR-HR40001", "Invalid X-Request-ID format. It must be a valid UUID")
	SystemError        = NewErrorCode("ERR-SY50001", "A system error has occurred, please try again later")
)
