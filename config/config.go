package config

import (
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Writetimeout    time.Duration
	Readtimeout     time.Duration
	Idletimeout     time.Duration
	Gracefultimeout time.Duration
	Environment     string
	Debug           bool
	Db              map[string]string
	URL             string
}

func GetConfig(configPath string) (*Config, error) {
	var config Config

	splits := strings.Split(filepath.Base(configPath), ".")
	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(configPath))
	viper.AddConfigPath("/Users/reddaemon/go/src/github.com/reddaemon/apiproject")
	viper.AddConfigPath("/etc")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to read config %s", err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}

	if config.Environment == "" {
		log.Fatal("Unable to find environment parameter in config")
	}

	return &config, nil

}
