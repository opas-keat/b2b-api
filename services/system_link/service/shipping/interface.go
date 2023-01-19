package shipping

import (
	"context"
	"models"
	"models/shipping"
)

type Service interface {
	GetShippingByCode(ctx context.Context, code string) (*shipping.Shipping, error)
	ListShippingByRegion(ctx context.Context, filter shipping.ListShippingRequest, pagination models.Pagination) (*models.ListResponse[shipping.Shipping], error)
	// ListShippingByCodeInternal(ctx context.Context, filter shipping.ListShippingRequest, pagination models.Pagination) (*[]shipping.Shipping, error)
}
