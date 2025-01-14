package service

import (
	"go-react-graphql-orders/utils"

	"go-react-graphql-orders/config"
	"go-react-graphql-orders/model"
)

type OrderService struct {
	*ServiceConfig
}

var lastId int

const (
	KeyOrderService    = "orderService"
	KeyOrderCreated    = "order.create"
	KeyTplOrderChanged = "order.update.%s"
)

func NewOrderService() *OrderService {
	return &OrderService{
		&ServiceConfig{
			config.GetEventHub(),
			config.Get(),
		},
	}
}

func (s *OrderService) GetAll() (order []*model.Order, err error) {
	orders, err := utils.LoadOrdersFromFile("utils/orders.json")
	return orders, err
}

func (s *OrderService) Get(id string) (order *model.Order, err error) {
	orders, err := utils.LoadOrdersFromFile("utils/orders.json")
	return orders[0], err
}

func (s *OrderService) Create(order *model.Order) (*model.Order, error) {
	return order, nil
}

func (s *OrderService) Update(order *model.Order) (*model.Order, error) {
	utils.UpdateOrderInFile(order.ID, order, "utils/orders.json")
	return order, nil
}
