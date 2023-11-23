package main

import (
	"log"
	"nats-streaming-service/internal/config"
	"nats-streaming-service/internal/nats_streaming"
)

func main() {
	configPath := "../../config/config.publisher.yaml"
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	natsClient, err := config.ConnectToNatsStreaming(cfg.NastStreaming)
	if err != nil {
		log.Fatalf("Failed to connect to nats-streaming: %v", err)
	}
	log.Println("Successfully connected to nats-streaming!")

	defer func() {
		if err := natsClient.Close(); err != nil {
			log.Fatalf("Failed to close nats-streaming: %v", err)
		}
	}()

	natsStreaming := nats_streaming.CreateNSPublisher(cfg.NastStreaming.Channel, natsClient)
	natsStreaming.NatsStreamingSubscribePublisher()
}
