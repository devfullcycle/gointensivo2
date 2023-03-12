package entity

import "errors"

type OrderItem struct {
	ItemID   string
	Price    float64
	Quantity int
}

func NewOrderItem(itemID string, price float64, quantity int) (*OrderItem, error) {
	orderItem := &OrderItem{
		ItemID:   itemID,
		Price:    price,
		Quantity: quantity,
	}
	return orderItem, nil
}

func (o *OrderItem) Validate() error {
	if o.ItemID == "" {
		return errors.New("ItemID is required")
	}
	if (o.Price) <= 0 {
		return errors.New("price should greate than zero")
	}
	if (o.Quantity) <= 0 {
		return errors.New("quantity should greate than zero")
	}
	return nil
}

func (oi *OrderItem) CalculateTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}
