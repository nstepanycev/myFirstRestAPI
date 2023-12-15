package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env            string        `yaml:"env"`
	StorageConfig  StorageConfig `yaml:"db"`
	HTTP_Server    HTTPServer    `yaml:"http_server"`
	ConfigPath     string        `yaml:"config_path"`
}

type HTTPServer struct {
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	// Auth later fix
	// User        string        `yaml:"user" env-required:"true"`
    // Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type StorageConfig struct {
	DbHost string `yaml:"dbHost"`
	DbPort string `yaml:"dbPort"`
	DbName string `yaml:"dbName"`
	DbUser string `yaml:"dbUser"`
	DbPass string `yaml:"dbPass"`
}

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
