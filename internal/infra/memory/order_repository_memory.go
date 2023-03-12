package memory

import "github.com/devfullcycle/gointensivo2/internal/entity"

type OrderRepositoryMemory struct {
	Orders []*entity.Order
}

func NewOrderRepositoryMemory() *OrderRepositoryMemory {
	var orders []*entity.Order
	order, _ := entity.NewOrder("1")
	order.AddItem(&entity.Item{
		ID:    "1",
		Name:  "Item 1",
		Price: 10.00,
	}, 1)
	orders = append(orders, order)
	return &OrderRepositoryMemory{
		Orders: orders,
	}
}

func (o *OrderRepositoryMemory) Save(order *entity.Order) error {
	o.Orders = append(o.Orders, order)
	return nil
}

func (o *OrderRepositoryMemory) GetTotal() (int, error) {
	var total int
	for _, order := range o.Orders {
		total += int(order.CalculateFinalPrice())
	}
	return total, nil
}
