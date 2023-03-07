package memory

import (
	"fmt"
	"testing"

	"github.com/devfullcycle/gointensivo2/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestShouldSaveOrder(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 1.0)
	assert.NoError(t, err)
	orderRepository := NewOrderRepositoryMemory()
	orderRepository.Save(order)
	total, _ := orderRepository.GetTotal()
	fmt.Println(total)
	assert.Equal(t, 11.0, float64(total))
}
