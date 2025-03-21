package http

const (
	TooManyReqs        string = "REQ-000001"
	NotValidJSONFormat string = "ERR-PA40001"
	InputNotValid      string = "ERR-PA40002"
	Forbidden          string = "ERR-PA40003"
	NotValidQuery      string = "ERR-PA40004"
	SystemError        string = "ERR-SY50001"
	NotValidJSON       string = "ERR-SY50002"
)

func Text(code string) string {
	return text[code]
}
