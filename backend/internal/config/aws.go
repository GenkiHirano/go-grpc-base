package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

type AWS struct {
	AccessKey       string `envconfig:"AWS_ACCESS_KEY" required:"true"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region          string `envconfig:"AWS_REGION" required:"true"`
	Endpoint        string `envconfig:"AWS_ENDPOINT" required:"true"`
}

func (a *AWS) Init(ctx context.Context, env string) (*aws.Config, error) {
	if env == Local {
		customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...any) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               a.Endpoint,
				HostnameImmutable: true,
			}, nil
		})

		cfg, err := awsconfig.LoadDefaultConfig(ctx,
			awsconfig.WithRegion(a.Region),
			awsconfig.WithEndpointResolverWithOptions(customResolver),
		)
		if err != nil {
			// TODO: カスタムエラーで返す
			return nil, err
		}

		return &cfg, nil
	}

	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(a.Region))
	if err != nil {
		// TODO: カスタムエラーで返す
		return nil, err
	}

	return &cfg, nil
}
