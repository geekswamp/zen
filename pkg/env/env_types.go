package env

const (
	// Dev is development environment mode
	Dev string = "dev"

	// Pro is production environment mode
	Pro string = "pro"
)

var (
	active EnvironmentTypes
	devEnv EnvironmentTypes = &Environment{value: Dev}
	proEnv EnvironmentTypes = &Environment{value: Pro}
)

type EnvironmentTypes interface {
	Value() string
	IsDev() bool
	IsPro() bool
}
