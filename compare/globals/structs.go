package globals

type Config struct {
	AppPort          int    `toml:"APP_PORT"`
	AppName          string `toml:"APP_NAME"`
}
