package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnExceptionIfEmptyID(t *testing.T) {
	order, err := NewOrder("")
	assert.EqualError(t, err, "id is required")
	assert.Nil(t, order)
}

func TestShouldCreateOrderWithItem(t *testing.T) {
	order, err := NewOrder("123")
	assert.NoError(t, err)
	order.AddItem(&Item{
		ID:    "1",
		Name:  "Product 1",
		Price: 10.00,
	}, 2)
	order.AddItem(&Item{
		ID:    "2",
		Name:  "Product 2",
		Price: 15.00,
	}, 1)
	assert.Len(t, order.OrderItems, 2)
}

func TestShouldCreateOrderAndCalculateTotal(t *testing.T) {
	order, err := NewOrder("123")
	assert.NoError(t, err)
	order.AddItem(&Item{
		ID:    "1",
		Name:  "Product 1",
		Price: 10.00,
	}, 2)
	order.AddItem(&Item{
		ID:    "2",
		Name:  "Product 2",
		Price: 15.00,
	}, 1)
	assert.Equal(t, 35.00, order.CalculateFinalPrice())
}
