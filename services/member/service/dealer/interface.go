package dealer

import (
	"context"
	"member/entities"
)

type Service interface {
	CreateBatch(ctx context.Context, d *[]entities.Dealer) (*[]entities.Dealer, error)
}
