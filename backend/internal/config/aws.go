package config

type AWSConfig struct {
	AccessKey       string `envconfig:"AWS_ACCESS_KEY" required:"true"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region          string `envconfig:"AWS_REGION" required:"true"`
}
