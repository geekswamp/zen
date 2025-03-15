package env

import (
	"flag"
	"fmt"
	"strings"
)

type Environment struct {
	value string
}

func Active() EnvironmentTypes {
	return active
}

func (e *Environment) Value() string {
	return e.value
}

func (e *Environment) IsDev() bool {
	return e.value == Dev
}

func (e *Environment) IsPro() bool {
	return e.value == Pro
}

func init() {
	env := flag.String("env", "dev", "Your environment mode (dev or pro)")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case Pro:
		active = proEnv
	default:
		active = devEnv
		fmt.Println("Warning: '-env' not found or is invalid. Defaulting to 'dev'.")
	}
}
