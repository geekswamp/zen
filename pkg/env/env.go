package env

import (
	"flag"
	"fmt"
	"strings"
)

const (
	// Dev is development environment mode
	Dev string = "dev"

	// Pro is production environment mode
	Pro string = "pro"
)

var (
	_      Environment = (*environment)(nil)
	active Environment
	devEnv Environment = &environment{value: Dev}
	proEnv Environment = &environment{value: Pro}
)

type (
	Environment interface {
		Value() string
		IsDev() bool
		IsPro() bool
		t()
	}

	environment struct{ value string }
)

func Active() Environment {
	return active
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == Dev
}

func (e *environment) IsPro() bool {
	return e.value == Pro
}

func (e *environment) t() {}

func init() {
	env := flag.String("env", "", "Please enter the environment (dev or pro)")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case Pro:
		active = proEnv
	default:
		active = devEnv
		fmt.Println("Warning: '-env' not found or is invalid. Defaulting to 'dev'.")
	}
}
