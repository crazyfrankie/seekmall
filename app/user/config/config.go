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
	MySQL  MySQL  `yaml:"mysql"`
	JWT    JWT    `yaml:"jwt"`
	SMS    SMS    `yaml:"sm"`
	ETCD   ETCD   `yaml:"etcd"`
	System System `yaml:"system"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type JWT struct {
	SecretKey string `yaml:"secret_key"`
}

type SMS struct {
	TemplateID string `yaml:"template_id"`
}

type ETCD struct {
	Addr string `yaml:"addr"`
}

type Minio struct {
	EndPoint   string `yaml:"endPoint"`
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	BucketName string `yaml:"bucketName"`
}

type System struct {
	DefaultAvatar string `yaml:"defaultAvatar"`
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
