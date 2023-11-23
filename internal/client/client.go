package client

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"nats-streaming-service/internal/cache"
	"nats-streaming-service/internal/model"
	"nats-streaming-service/internal/repository"
)

type Client struct {
	db              *gorm.DB
	cache           *cache.Cache
	orderRepository *repository.OrderRepository
}

func NewClient(db *gorm.DB) *Client {
	return &Client{
		db:              db,
		cache:           cache.NewCache(),
		orderRepository: repository.NewOrderRepository(db),
	}
}

func (client *Client) Start() error {
	err := client.orderRepository.AutoMigrate()
	if err != nil {
		return fmt.Errorf("Failed to automigrate: %v", err)
	}

	orders, err := client.orderRepository.GetAll()
	if err != nil {
		return fmt.Errorf("Failed to add orders: %v", err)
	}
	client.cache.AddOrders(orders)
	return nil
}

func (c *Client) GetOrder(orderUid string) (*model.Order, error) {
	order, err := c.cache.GetByUid(orderUid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order, nil
}

func (c *Client) GetAllOrders() (*[]model.Order, error) {
	orders, err := c.cache.GetOrders()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &orders, nil
}

func (c *Client) AddOrder(data []byte) error {
	var order model.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		log.Printf("Failed to unmarshal data: %v", err)
		return err
	}
	log.Printf("Getting order with uid %s", order.OrderUID)

	if _, err := c.cache.GetByUid(order.OrderUID); err != nil {
		if err := c.orderRepository.AddOrder(order); err != nil {
			log.Printf("Failed to create order: %v", err)
			return err
		}
		err := c.cache.AddOrder(order)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("Order with Uid %s is already exists", order.OrderUID)
	}
}
