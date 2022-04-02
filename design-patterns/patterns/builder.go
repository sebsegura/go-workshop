package main

import (
	"fmt"
)

const (
	Pending   = "PENDING"
	Processed = "PROCESSED"
	Ecommerce = "ECOMMERCE"
)

type Order interface {
	SetProcessed()
}

type order struct {
	OrderID string  `json:"orderId"`
	Amount  float64 `json:"amount"`
	Origin  string  `json:"origin"`
	Status  string  `json:"status"`
}

func (o *order) SetProcessed() {
	o.Status = Processed
}

type OrderBuilder interface {
	SetOrderId(id string) OrderBuilder
	SetAmount(amount float64) OrderBuilder
	SetOrigin(origin string) OrderBuilder
	Build() Order
}

type orderBuilder struct {
	id     string
	amount float64
	origin string
}

func (b *orderBuilder) SetOrderId(id string) OrderBuilder {
	b.id = id
	return b
}

func (b *orderBuilder) SetAmount(amount float64) OrderBuilder {
	b.amount = amount
	return b
}

func (b *orderBuilder) SetOrigin(origin string) OrderBuilder {
	b.origin = origin
	return b
}

func (b *orderBuilder) Build() Order {
	return &order{
		OrderID: b.id,
		Amount:  b.amount,
		Origin:  b.origin,
		Status:  Pending,
	}
}

func NewOrderBuilder() OrderBuilder {
	return &orderBuilder{}
}

func main() {
	builder := NewOrderBuilder()
	order := builder.SetOrderId("1").SetAmount(100.9).SetOrigin(Ecommerce).Build()
	fmt.Println(order)
	order.SetProcessed()
	fmt.Println(order)
}
