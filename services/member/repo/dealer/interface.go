package dealer

import (
	"context"
	"member/entities"
)

type Repo interface {
	Create(ctx context.Context, member entities.Dealer) (*entities.Dealer, error)
	Get(ctx context.Context, filter entities.Dealer) (*entities.Dealer, error)
	Update(ctx context.Context, dealerId string, updateModel entities.Dealer) error
	CreateBatch(ctx context.Context, dealers []entities.Dealer) error
}
