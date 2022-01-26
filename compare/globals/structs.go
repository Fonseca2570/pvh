package globals

import "github.com/go-playground/validator/v10"

type Config struct {
	AppPort int    `toml:"APP_PORT"`
	AppName string `toml:"APP_NAME"`
}

type Application struct {
	Config    Config
	Validator *validator.Validate
}
