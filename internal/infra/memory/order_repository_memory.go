package memory

import "github.com/devfullcycle/gointensivo2/internal/entity"

type OrderRepositoryMemory struct {
	Orders []*entity.Order
}

func NewOrderRepositoryMemory() *OrderRepositoryMemory {
	return &OrderRepositoryMemory{
		make([]*entity.Order, 0),
	}
}

func (o *OrderRepositoryMemory) Save(order *entity.Order) error {
	o.Orders = append(o.Orders, order)
	return nil
}

func (o *OrderRepositoryMemory) GetTotal() (int, error) {
	var total int
	for _, order := range o.Orders {
		order.CalculateFinalPrice()
		total = int(order.FinalPrice)
	}
	return total, nil
}
