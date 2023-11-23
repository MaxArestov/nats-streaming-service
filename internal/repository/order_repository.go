package repository

import (
	"encoding/json"
	"gorm.io/gorm"
	"nats-streaming-service/internal/model"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (rep *OrderRepository) AddOrder(order model.Order) error {
	orderJsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	var orderJson model.OrderJson

	orderJson.OrderUID = order.OrderUID
	orderJson.Data = json.RawMessage(orderJsonData)
	return rep.db.Create(orderJson).Error
}

func (rep *OrderRepository) GetAll() ([]model.Order, error) {
	var ordersJson []model.OrderJson
	result := rep.db.Find(&ordersJson)
	if result.Error != nil {
		return nil, result.Error
	}

	orders := make([]model.Order, 0, len(ordersJson))

	for _, record := range ordersJson {
		var order model.Order
		err := json.Unmarshal(record.Data, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (rep *OrderRepository) GetById(orderUID string) (model.Order, error) {
	var orderJson model.OrderJson
	result := rep.db.Where("order_uid= ?", orderUID).First(&orderJson)
	if result.Error != nil {
		return model.Order{}, result.Error
	}

	var order model.Order
	err := json.Unmarshal(orderJson.Data, &order)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (rep *OrderRepository) Delete(orderUID string) error {
	return rep.db.Where("order_uid= ?", orderUID).Delete(&model.OrderJson{}).Error
}

func (rep *OrderRepository) AutoMigrate() error {
	return rep.db.AutoMigrate(&model.OrderJson{})
}
