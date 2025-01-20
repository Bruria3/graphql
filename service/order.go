package service

import (
	"go-react-graphql-orders/config"
	"go-react-graphql-orders/model"
	"go-react-graphql-orders/utils"
	"log"
	"time"

	"github.com/leandro-lugaresi/hub"
)

type OrderService struct {
	*ServiceConfig
}

var lastId int

const (
	KeyOrderService    = "orderService"
	KeyTplOrderChanged = "order.update"
)

func NewOrderService() *OrderService {
	return &OrderService{
		&ServiceConfig{
			config.GetEventHub(),
			config.Get(),
		},
	}
}

func logAPIUsage(apiName string, details map[string]interface{}) {
	log.Printf("[%s] API Called: %s, Details: %v", time.Now().Format("2006-01-02 15:04:05"), apiName, details)
}

func (s *OrderService) GetAll() ([]*model.Order, error) {
	logAPIUsage("GetAll", nil)
	orders, err := utils.LoadOrdersFromFile("utils/orders.json")
	if err != nil {
		log.Printf("Error in GetAll: %v", err)
	}
	return orders, err
}

func (s *OrderService) Get(id string) (*model.Order, error) {
	logAPIUsage("Get", map[string]interface{}{"id": id})
	orders, err := utils.LoadOrdersFromFile("utils/orders.json")
	if err != nil {
		log.Printf("Error in Get: %v", err)
		return nil, err
	}
	if len(orders) == 0 {
		log.Printf("No orders found in Get")
		return nil, nil
	}
	return orders[0], nil
}

func (s *OrderService) Create(order *model.Order) (*model.Order, error) {
	logAPIUsage("Create", map[string]interface{}{"orderID": order.ID})
	s.Hub.Publish(hub.Message{Name: KeyTplOrderChanged})
	return order, nil
}

func (s *OrderService) Update(order *model.Order) (*model.Order, error) {
	logAPIUsage("Update", map[string]interface{}{"orderID": order.ID})
	_, err := utils.UpdateOrderInFile(order.ID, order, "utils/orders.json")
	if err != nil {
		log.Printf("Error in Update: %v", err)
		return nil, err
	}
	s.Hub.Publish(hub.Message{Name: KeyTplOrderChanged})
	return order, nil
}
