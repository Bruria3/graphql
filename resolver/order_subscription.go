package resolver

import (
	"context"
	"fmt"
	"go-react-graphql-orders/service"
	"log"

	"github.com/leandro-lugaresi/hub"
)

func (r *Resolver) OrdersUpdated(ctx context.Context) (chan *[]*orderResolver, error) {
	c := make(chan *[]*orderResolver)
	orderService := ctx.Value(service.KeyOrderService).(*service.OrderService)

	go func() {
		sub := orderService.Hub.NonBlockingSubscribe(10, "application.ordersUpdated")
		orderService.Hub.Publish(hub.Message{Name: "application.ordersUpdated"})

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

func (r *Resolver) OrderCreated(ctx context.Context) (chan *orderResolver, error) {
	c := make(chan *orderResolver)
	go subscribeOrder(service.KeyOrderCreated, ctx, c)
	return c, nil
}

func (r *Resolver) OrderChanged(ctx context.Context, args struct{ Id string }) (chan *orderResolver, error) {
	c := make(chan *orderResolver)
	go subscribeOrder(fmt.Sprintf(service.KeyTplOrderChanged, args.Id), ctx, c)
	return c, nil
}

func subscribeOrder(key string, ctx context.Context, c chan *orderResolver) {
	srv := ctx.Value(service.KeyOrderService).(*service.OrderService)
	sub := srv.Hub.NonBlockingSubscribe(10, key)

	defer func() {
		srv.Hub.Unsubscribe(sub)
		close(c)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case m := <-sub.Receiver:
			if u, err := srv.Get(m.Fields["id"].(string)); err != nil {
				log.Println(err)
			} else {
				c <- &orderResolver{u}
			}
		}
	}
}
