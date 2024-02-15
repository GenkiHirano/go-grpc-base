package config

import (
	"context"
)

type Environment struct {
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
}

const (
	Local       = "local"
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

var validEnvironments = map[string]bool{
	Local:       true,
	Development: true,
	Staging:     true,
	Production:  true,
}

func (e *Environment) validateContains(ctx context.Context) error {
	if _, ok := validEnvironments[e.Environment]; !ok {
		// TODO: 実装する
		// msg := fmt.Sprintf("the value '%s' set for the environment variable 'ENVIRONMENT' is invalid. Valid values are 'local', 'development', 'staging', or 'production'. please check your configuration.", e.Environment)
		// logging.FatalWithCtx(ctx, msg)
		// return apperror.NewError(apperror.NewDetail(apperror.CodeInternal, ""))
		return nil
	}
	return nil
}
