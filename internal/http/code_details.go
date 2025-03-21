package http

type codeMap map[string]string

var text = codeMap{
	TooManyReqs:          "Too many requests, please try again later",
	PANotValidJSONFormat: "Payload not valid JSON format",
	PAInputNotValid:      "The provided input is not valid",
	PAForbidden:          "Access to this resource is forbidden",
	SYSSystemError:       "A system error has occurred, please try again later",
	SYSNotValidJSON:      "The provided JSON format is invalid",
}
