package http

type Errno string

const (
	TooManyReqs        Errno = "REQ-000001"
	NotValidJSONFormat Errno = "ERR-PA40001"
	InputNotValid      Errno = "ERR-PA40002"
	Forbidden          Errno = "ERR-PA40003"
	NotValidQuery      Errno = "ERR-PA40004"
	SystemError        Errno = "ERR-SY50001"
	NotValidJSON       Errno = "ERR-SY50002"
)

func Text(code Errno) string {
	return text[code]
}
