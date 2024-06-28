package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Type     string `mapstructure:"type" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Port     string `mapstructure:"port" validate:"required"`
		Name     string `mapstructure:"name" validate:"required"`
		Host     string `mapstructure:"host" validate:"required"`
	} `mapstructure:"database"`
	Server struct {
		Port string `mapstructure:"port" validate:"required"`
		Host string `mapstructure:"host" validate:"required"`
	} `mapstructure:"server"`
}

func ReadConfig() (*config, error) {
	config := &config{}
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			fmt.Println(e)
		}
		return nil, fmt.Errorf("error validating config, %v", err)
	}

	spew.Dump(config)
	return config, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
