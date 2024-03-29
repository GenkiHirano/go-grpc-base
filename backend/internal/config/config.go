package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment Environment
	App         App
	DB          DB
	AWS         AWS
}

func Init(ctx context.Context) (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		// TODO: 実装する
		// errMsg := "failed to envconfig process"
		// logging.FatalWithCtxAndError(ctx, err, errMsg)
		// return nil, apperror.WrapWithMessage(err, apperror.NewDetail(apperror.CodeInternal, ""))
		return nil, err
	}

	if err := cfg.Environment.validateContains(ctx); err != nil {
		return nil, err
	}

	return &cfg, nil
}
