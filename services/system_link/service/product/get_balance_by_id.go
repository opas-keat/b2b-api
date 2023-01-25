package product

import (
	"context"
	"errors"
	"fmt"
)

func (s ServiceImpl) GetBalanceById(ctx context.Context, id string) (float64, error) {
	balance, err := s.productRepo.GetBalance(ctx, id)
	if err != nil {
		return 0, errors.New("get year by product id not found")
	}
	fmt.Println("---- GetBalanceById--- balance = " + fmt.Sprintf("%f", balance))
	return balance, nil
}
