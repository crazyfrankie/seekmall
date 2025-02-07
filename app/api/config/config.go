package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Env    string
	Server Server `yaml:"server"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
}
type JWT struct {
	SecretKey string `yaml:"secret_key"`
}

func GetConf() *Config {
	once.Do(initConf)
	return config
}

func initConf() {
	prefix := "config"
	filePath := filepath.Join(prefix, filepath.Join(getServerEnv(), "config.yaml"))
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	config = new(Config)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	config.Env = getServerEnv()
	fmt.Printf("%#v", config)
}

func getServerEnv() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "test"
	}

	return env
}
