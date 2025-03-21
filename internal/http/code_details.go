package http

var text = map[Errno]string{
	TooManyReqs:        "Too many requests, please try again later",
	NotValidJSONFormat: "Payload not valid JSON format",
	InputNotValid:      "The provided input is not valid",
	Forbidden:          "Access to this resource is forbidden",
	SystemError:        "A system error has occurred, please try again later",
	NotValidJSON:       "The provided JSON format is invalid",
}
