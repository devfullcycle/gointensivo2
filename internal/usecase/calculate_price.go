package usecase

import "github.com/devfullcycle/gointensivo2/internal/entity"

type OrderInputDTO struct {
	ID string
}

type OrderOutputDTO struct {
	ID         string
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrice) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	existingOrder, err := c.OrderRepository.Find(input.ID)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{
		ID:         existingOrder.ID,
		FinalPrice: existingOrder.CalculateFinalPrice(),
	}, nil
}
