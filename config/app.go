package config

type AppConfig struct {
	AppName string `env:"APP_NAME env-default:"OwO"`
	AppEnv string `env:"APP_ENV" env-default:"DEV"`
	Port string `env:"APP_PORT" env-default:":8080"`
	Host string `env:"HOST" env-default:"localhost"`
}
