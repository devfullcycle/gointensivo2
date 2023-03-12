package memory

import (
	"testing"

	"github.com/devfullcycle/gointensivo2/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestShouldSaveOrder(t *testing.T) {
	order, err := entity.NewOrder("123")
	assert.NoError(t, err)
	orderRepository := NewOrderRepositoryMemory()
	orderRepository.Save(order)
	total, _ := orderRepository.GetTotal()
	assert.Equal(t, 10.00, float64(total))
}
