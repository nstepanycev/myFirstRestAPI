package config

import (
	"os"
	"time"
	
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer  string `yaml:"http_server"`
	ConfigPath  string `yaml:"config_path"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type StorageConfig struct {
	DbHost string `yaml:"DbHost"`
	DbPort string `yaml:"DbPort"`
	DbName string `yaml:"DbName"`
	DbUser string `yaml:"DbUser"`
	DbPass string `yaml:"DbPass"`
}

func MustLoad() *Config {
	return &Config{
		Env: os.Getenv("ENV"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		HTTPServer: os.Getenv("HOST"),
		ConfigPath: os.Getenv("CONFIG_PATH"),
	}
}

func DbConfig() *StorageConfig{
	return &StorageConfig{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
	}
}
