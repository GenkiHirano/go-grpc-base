package config

type App struct {
	Key  string `envconfig:"APP_KEY" required:"true"`
	Name string `envconfig:"APP_NAME" required:"true"`
	Port string `envconfig:"APP_PORT" required:"true" default:"8080"`
}
