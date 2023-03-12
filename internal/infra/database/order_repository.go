package database

import (
	"database/sql"

	"github.com/devfullcycle/gointensivo2/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("insert into orders (id) Values(?)", order.ID)
	if err != nil {
		return err
	}
	for _, orderItem := range order.OrderItems {
		_, err := r.Db.Exec("insert into order_item(order_id, item_id, price, quantity, total) values(?, ?, ?, ?, ?)",
			order.ID, orderItem.ItemID, orderItem.Price, orderItem.Quantity, order.CalculateFinalPrice())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) Find(OrderID string) (*entity.Order, error) {
	var order *entity.Order
	err := r.Db.QueryRow("select id from orders where id = ?", OrderID).Scan(&order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
