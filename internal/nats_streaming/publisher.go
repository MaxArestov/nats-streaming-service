package nats_streaming

import (
	"encoding/json"
	"fmt"
	"log"
	"nats-streaming-service/internal/model"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

type NatsStreamingPublisher struct {
	natsClient stan.Conn
	channel    string
}

func CreateNSPublisher(channel string, client stan.Conn) *NatsStreamingPublisher {
	return &NatsStreamingPublisher{
		natsClient: client,
		channel:    channel,
	}
}

func (ns *NatsStreamingPublisher) NatsStreamingSubscribePublisher() {
	jsonModel, err := os.ReadFile("../../internal/nats_streaming/model.json")
	if err != nil {
		log.Fatalf("Failed to read json: %v", err)
	}
	var order model.Order

	err = json.Unmarshal(jsonModel, &order)
	if err != nil {
		log.Fatalf("Failed to unmarshal model: %v", err)
	}

	for {
		order.OrderUID = fmt.Sprintf("%d", int(time.Now().Unix()))
		if err != nil {
			log.Println("Failed to generate UID")
		}

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Printf("Failed to marshal json: %v", err)
			continue
		}

		err = ns.natsClient.Publish(ns.channel, bytes)
		if err != nil {
			log.Printf("Failed to publish msg: %v", err)
			continue
		}
		time.Sleep(5 * time.Second)
	}
}
