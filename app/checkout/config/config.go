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
	once sync.Once
	conf *Config
)

type Config struct {
	Env    string
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
	ETCD   ETCD   `yaml:"etcd"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type ETCD struct {
	Addr string `yaml:"addr"`
}

type JWT struct {
	SecretKey string `yaml:"secretKey"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "config"
	filePath := filepath.Join(prefix, filepath.Join(getServerEnv(), "config.yaml"))
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	conf = new(Config)
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	conf.Env = getServerEnv()
	fmt.Printf("%#v", conf)
}

func getServerEnv() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "test"
	}

	return env
}
