package resolver

import (
	"go-react-graphql-orders/service"

	"golang.org/x/net/context"
)

func (r *Resolver) Order(ctx context.Context, args struct {
	Id string
}) (*orderResolver, error) {
	order, err := ctx.Value(service.KeyOrderService).(*service.OrderService).Get(args.Id)
	if err != nil {
		return nil, err
	}

	return &orderResolver{order}, nil
}

func (r *Resolver) Orders(ctx context.Context) (*[]*orderResolver, error) {
	orders, err := ctx.Value(service.KeyOrderService).(*service.OrderService).GetAll()

	if err != nil {
		return nil, err
	}

	l := make([]*orderResolver, len(orders))
	for i := range l {
		l[i] = &orderResolver{o: orders[i]}
	}

	return &l, nil
}
