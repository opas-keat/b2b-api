package dealerconnector

import (
	"context"
	"models/dealer"
)

type DealerConnector interface {
	GetDealer(ctx context.Context, dealerCode string) (*dealer.DealerResponse, error)
}
