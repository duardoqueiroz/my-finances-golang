package config

import (
	"fmt"
	"log"
	"os"
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
	} `mapstructure:"server"`
}

var C config

func ReadConfig() {
	config := &C
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			fmt.Println(e)
		}
		os.Exit(1)
	}

	spew.Dump(C)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
