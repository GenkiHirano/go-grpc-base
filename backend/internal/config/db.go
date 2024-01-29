package config

type DBConfig struct {
	Driver   string `envconfig:"DB_DRIVER" required:"true" default:"mysql"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true" default:"3306"`
	Username string `envconfig:"DB_USERNAME" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true" redact:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}
