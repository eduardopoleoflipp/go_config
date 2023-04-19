package config

var (
	Env     Config
	EnvName string
)

type Config struct {
	name string
	url  string
}
