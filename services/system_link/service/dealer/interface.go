package dealer

import (
	"context"
	"models"
	"models/dealer"
)

type Service interface {
	GetDealerByCode(ctx context.Context, code string) (*dealer.DealerResponse, error)
	ListDealerByCode(ctx context.Context, filter dealer.ListDealerRequest, pagination models.Pagination) (*models.ListResponse[dealer.DealerResponse], error)
}
