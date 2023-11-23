package cache

import (
	"fmt"
	"nats-streaming-service/internal/model"
)

type Cache struct {
	cache map[string]model.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]model.Order)}
}

func (c *Cache) GetByUid(uId string) (*model.Order, error) {
	order, ok := c.cache[uId]
	if !ok {
		return nil, fmt.Errorf("Order not found")
	}
	return &order, nil
}

func (c *Cache) GetOrders() ([]model.Order, error) {
	if len(c.cache) == 0 {
		return nil, fmt.Errorf("No orders in cache")
	}

	var orders []model.Order
	for _, order := range c.cache {
		orders = append(orders, order)
	}
	return orders, nil
}

func (c *Cache) AddOrders(orders []model.Order) error {
	for _, order := range orders {
		c.cache[order.OrderUID] = order
	}
	return nil
}

func (c *Cache) AddOrder(order model.Order) error {
	if _, exists := c.cache[order.OrderUID]; exists {
		return fmt.Errorf("Order with orderUID %s is already exists", order.OrderUID)
	}
	c.cache[order.OrderUID] = order
	return nil
}
