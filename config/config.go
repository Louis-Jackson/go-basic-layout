package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func LoadConfig() Config {
	var config Config
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Error reading config.yaml: %v", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error parsing config.yaml: %v", err)
	}

	return config
}
