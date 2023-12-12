package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env            string `yaml:"env"`
	StorageConfig  StorageConfig
	HTTPServer     HTTPServer
	ConfigPath     string `yaml:"config_path"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	// User        string        `yaml:"user" env-required:"true"`
    // Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type StorageConfig struct {
	DbHost string `yaml:"DbHost"`
	DbPort string `yaml:"DbPort"`
	DbName string `yaml:"DbName"`
	DbUser string `yaml:"DbUser"`
	DbPass string `yaml:"DbPass"`
}

// func MustLoad() *Config {
// 	return &Config{
// 		Env: os.Getenv("ENV"),
// 		HTTPServer: os.Getenv("HTTP_HOST"),
// 		ConfigPath: os.Getenv("CONFIG_PATH"),
// 	}
// }

// func DbConfig() *StorageConfig{
// 	return &StorageConfig{
// 		DbHost: os.Getenv("DB_HOST"),
// 		DbPort: os.Getenv("DB_PORT"),
// 		DbName: os.Getenv("DB_NAME"),
// 		DbUser: os.Getenv("DB_USER"),
// 		DbPass: os.Getenv("DB_PASS"),
// 	}
// }

func LoadConfig() *Config{
	file, err := os.ReadFile("local.yaml")
	if err != nil{
		fmt.Println("Error",err)
	}

	cfg := &Config{}

	dec := yaml.Unmarshal(file,cfg)
	if err != nil{
		fmt.Println(dec, err)
	}
	return cfg
}
