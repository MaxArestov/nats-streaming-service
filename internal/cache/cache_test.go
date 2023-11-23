package cache

import (
	"github.com/stretchr/testify/assert"
	"nats-streaming-service/internal/model"
	"testing"
)

func TestCache_NewCache(t *testing.T) {
	c := NewCache()
	assert.NotNil(t, c, "New cache should be not nil")
	assert.Emptyf(t, c.cache, "New cache should be empty")
}

func TestCache_GetByUid(t *testing.T) {
	var expectedUid string = "123"
	c := NewCache()
	order := model.Order{OrderUID: "123"}
	err := c.AddOrder(order)

	assert.Nil(t, err, "Error should be empty with Adding order")

	assert.Equal(t, order.OrderUID, expectedUid, "Expected and actual UID should be equals")
}

func TestCache_AddOrder(t *testing.T) {
	c := NewCache()

	err := c.AddOrder(model.Order{OrderUID: "2tou4og"})
	assert.Nil(t, err, "Error should be empty with adding order")

	err = c.AddOrder(model.Order{OrderUID: "2tou4og"})
	assert.Error(t, err, "AddOrder should return error if order is duplicated")
}
