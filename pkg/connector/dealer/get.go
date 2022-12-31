package dealerconnector

import (
	"context"
	"fmt"
	"models/dealer"
	"shareerrors"
	"shareutils"
)

func (i impl) GetDealer(ctx context.Context, dealerCode string) (*dealer.Dealer, error) {
	resp, err := shareutils.Get[dealer.Dealer](ctx, fmt.Sprintf("%v/api/v1/dealers/%v", i.url, dealerCode))
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, shareerrors.NewError(resp.StatusCode, resp.Message).SetData(resp.Data)
	}

	return resp.Data, nil
}
