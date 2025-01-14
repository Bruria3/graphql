package resolver

import (
	"go-react-graphql-orders/model"

	"github.com/graph-gophers/graphql-go"
)

type orderResolver struct {
	o *model.Order
}

// additional types

type OrderInput struct {
	Name      *string
	Status    *string
	CreatedAt *string
	Quantity  *int32
}

func (i *OrderInput) ToEntity() *model.Order {
	return &model.Order{
		Name:      stringOrDefault(i.Name),
		Status:    stringOrDefault(i.Status),
		Quantity:  intOrDefault(i.Quantity),
		CreatedAt: stringOrDefault(i.CreatedAt),
	}
}

func stringOrDefault(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func intOrDefault(s *int32) int32 {
	if s == nil {
		return 1
	}
	return *s
}

type OrderUpdateInput struct {
	Id       string
	Quantity *int32
}

func (i *OrderUpdateInput) ToEntity() *model.Order {
	return &model.Order{
		ID:       i.Id,
		Quantity: intOrDefault(i.Quantity),
	}
}

func (r *orderResolver) Id() graphql.ID {
	return graphql.ID(r.o.ID)
}

func (r *orderResolver) Quantity() int32 {
	s := int32(r.o.Quantity)
	return s
}

func (r *orderResolver) Name() *string {
	if r.o.Name == "" {
		return nil
	}
	return &r.o.Name
}

func (r *orderResolver) Status() *string {
	if r.o.Status == "" {
		return nil
	}
	return &r.o.Status
}

func (r *orderResolver) CreatedAt() *string {
	if r.o.CreatedAt == "" {
		return nil
	}
	return &r.o.CreatedAt
}
