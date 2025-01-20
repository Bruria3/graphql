package resolver

import (
	"context"
	"go-react-graphql-orders/service"
	"log"

	"github.com/leandro-lugaresi/hub"
)

func (r *Resolver) OrdersUpdated(ctx context.Context) (chan *[]*orderResolver, error) {
	c := make(chan *[]*orderResolver)
	orderService := ctx.Value(service.KeyOrderService).(*service.OrderService)

	go func() {
		sub := orderService.Hub.NonBlockingSubscribe(10, service.KeyTplOrderChanged)
		orderService.Hub.Publish(hub.Message{Name: service.KeyTplOrderChanged})

		defer func() {
			orderService.Hub.Unsubscribe(sub)
			close(c)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case <-sub.Receiver:
				orders, err := orderService.GetAll()
				if err != nil {
					log.Println(err)
					continue
				}

				resolvers := make([]*orderResolver, len(orders))
				for i, order := range orders {
					resolvers[i] = &orderResolver{o: order}
				}

				c <- &resolvers
			}
		}
	}()

	return c, nil
}
