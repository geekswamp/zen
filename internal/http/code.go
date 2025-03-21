package http

const (
	TooManyReqs          string = "REQ-000001"
	PANotValidJSONFormat string = "ERR-PA40001"
	PAInputNotValid      string = "ERR-PA40002"
	PAForbidden          string = "ERR-PA40003"
	PANotValidQuery      string = "ERR-PA40004"
	SYSSystemError       string = "ERR-SY50001"
	SYSNotValidJSON      string = "ERR-SY50002"
)

func Text(code string) string {
	return text[code]
}
