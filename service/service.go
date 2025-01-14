package service

import (
	"go-react-graphql-orders/config"

	eventhub "github.com/leandro-lugaresi/hub"
)

type Service interface{}

type ServiceConfig struct {
	Hub *eventhub.Hub
	cfg *config.Config
}
