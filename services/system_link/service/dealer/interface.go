package dealer

import (
	"context"
	"models"
	"models/dealer"
)

type Service interface {
	GetDealerByCode(ctx context.Context, code string) (*dealer.Dealer, error)
	ListDealerByCode(ctx context.Context, filter dealer.ListDealerRequest, pagination models.Pagination) (*models.ListResponse[dealer.Dealer], error)
	ListDealerByCodeInternal(ctx context.Context, filter dealer.ListDealerRequest, pagination models.Pagination) (*[]dealer.Dealer, error)
}
