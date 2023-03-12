package entity

import (
	"errors"
)

type Order struct {
	ID         string
	OrderItems []*OrderItem
}

func NewOrder(id string) (*Order, error) {
	order := &Order{
		ID:         id,
		OrderItems: make([]*OrderItem, 0),
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

func (o *Order) AddItem(item *Item, quantity int) {
	orderItem, err := NewOrderItem(item.ID, item.Price, quantity)
	if err != nil {
		panic(err)
	}
	o.OrderItems = append(o.OrderItems, orderItem)
}

func (o *Order) CalculateFinalPrice() float64 {
	finalPrice := 0.0
	for _, orderItem := range o.OrderItems {
		finalPrice += orderItem.CalculateTotal()
	}
	return finalPrice
}
