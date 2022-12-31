package dealerconnector

import (
	"context"
	"fmt"
	"models/dealer"
	"shareerrors"
	"shareutils"
)

func (i impl) GetDealer(ctx context.Context, dealerCode string) (*dealer.DealerResponse, error) {
	resp, err := shareutils.Get[dealer.DealerResponse](ctx, fmt.Sprintf("%v/api/v1/dealers/%v", i.url, dealerCode))
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, shareerrors.NewError(resp.StatusCode, resp.Message).SetData(resp.Data)
	}

	return resp.Data, nil
}
