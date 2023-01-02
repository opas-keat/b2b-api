package dealer

import (
	"context"
	"member/entities"
	"models/dealer"
)

type Service interface {
	CreateBatch(ctx context.Context, d *[]entities.Dealer) (*[]entities.Dealer, error)
	GetDealerByCode(ctx context.Context, code string) (*dealer.Dealer, error)
}
