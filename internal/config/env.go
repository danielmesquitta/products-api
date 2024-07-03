package config

import (
	"github.com/danielmesquitta/products-api/pkg/validator"
	"github.com/spf13/viper"
)

type Environment string

const (
	DevelopmentEnv Environment = "development"
	TestEnv        Environment = "test"
	ProductionEnv  Environment = "production"
)

type Env struct {
	validator *validator.Validator

	Environment  Environment `mapstructure:"ENVIRONMENT"`
	Port         string      `mapstructure:"PORT"`
	DBConnection string      `mapstructure:"DB_CONNECTION" validate:"required"`
}

func (e *Env) validate() error {
	if err := e.validator.Validate(e); err != nil {
		return err
	}

	if e.Environment == "" {
		e.Environment = DevelopmentEnv
	}
	if e.Port == "" {
		e.Port = "8080"
	}
	return nil
}

func LoadEnv(validator *validator.Validator) *Env {
	env := &Env{
		validator: validator,
	}

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	if err := env.validate(); err != nil {
		panic(err)
	}

	return env
}
