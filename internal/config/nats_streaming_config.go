package config

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterId string `yaml:"cluster_id"`
	ClientId  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

func ConnectToNatsStreaming(cfg NatsStreamingConfig) (stan.Conn, error) {
	host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	sc, err := stan.Connect(cfg.ClusterId+"", cfg.ClientId, stan.NatsURL(host))
	if err != nil {
		return nil, err
	}
	return sc, nil
}
