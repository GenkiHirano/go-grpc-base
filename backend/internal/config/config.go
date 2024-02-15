package config

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment Environment
	App         App
	DB          DB
	AWS         AWS
}

func Init(ctx context.Context) (*Config, error) {
	fmt.Println("config-Init-1")
	cfg := Config{}
	fmt.Println("config-Init-2")
	if err := envconfig.Process("", &cfg); err != nil {
		// TODO: 実装する
		// errMsg := "failed to envconfig process"
		// logging.FatalWithCtxAndError(ctx, err, errMsg)
		// return nil, apperror.WrapWithMessage(err, apperror.NewDetail(apperror.CodeInternal, ""))
		return nil, err
	}

	fmt.Println("config-Init-3")

	if err := cfg.Environment.validateContains(ctx); err != nil {
		return nil, err
	}

	fmt.Println("config-Init-4")

	return &cfg, nil
}
