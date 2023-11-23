package nats_streaming

import (
	"log"
	"nats-streaming-service/internal/client"

	"github.com/nats-io/stan.go"
)

type NatsStreamingSubscriber struct {
	nastClient stan.Conn
	channel    string
	client     *client.Client
}

func NewNSSubscriber(channel string, nastClient stan.Conn, client *client.Client) *NatsStreamingSubscriber {
	return &NatsStreamingSubscriber{
		nastClient: nastClient,
		channel:    channel,
		client:     client,
	}
}

func (ns *NatsStreamingSubscriber) NatsStreamingSubscribeSubscriber() {
	_, err := ns.nastClient.Subscribe(
		ns.channel, func(m *stan.Msg) {
			err := ns.client.AddOrder(m.Data)
			if err != nil {
				log.Printf("Failed to add order from streaming: %v", err)
			}
		})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
}
