package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Port     string
		Name     string
		Host     string
	}
	Server struct {
		Port string
	}
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

	spew.Dump(C)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
