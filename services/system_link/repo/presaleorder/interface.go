package presaleorder

import (
	"context"
	"systemlink/entities"
)

type Repo interface {
	GetPreSaleOrderNo(ctx context.Context, orderNo string) (string, error)
	Create(ctx context.Context, preSaleOrder entities.PreSaleOrder) (*entities.PreSaleOrder, error)
}
