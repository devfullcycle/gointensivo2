package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/gointensivo2/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(36) primary key)")
	db.Exec("CREATE TABLE order_item(order_id varchar(36) not null, item_id varchar(36) not null, price float not null, quantity integer not null, total float not null)")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestSavingOrder() {
	order, err := entity.NewOrder("123")
	order.AddItem(&entity.Item{
		ID: "1",
		Name: "Test",
		Price: 100.00,
	}, 1)
	suite.NoError(err)
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("select id from orders where id = ?",
		order.ID).Scan(&orderResult.ID)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(100.00, order.CalculateFinalPrice())
}
