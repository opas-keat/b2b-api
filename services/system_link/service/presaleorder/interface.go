package presaleorder

import (
	"context"
	"models/presaleorder"
)

type Service interface {
	CreatePreSaleOrder(ctx context.Context, req presaleorder.PreSaleOrderRequest) (*presaleorder.PreSaleOrderRequest, error)
}
