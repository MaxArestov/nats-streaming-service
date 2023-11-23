package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"nats-streaming-service/internal/http"
	"os"
)

type Config struct {
	Database      DbConfig            `yaml:"db"`
	NastStreaming NatsStreamingConfig `yaml:"nats-streaming"`
	Http          http.Config         `yaml:"http"`
}

func NewConfig(configPath string) (*Config, error) {
	yamlReader, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: %v", err)
	}
	cfg := &Config{}

	if err = yaml.Unmarshal(yamlReader, cfg); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %v", err)
	}
	return cfg, nil
}
