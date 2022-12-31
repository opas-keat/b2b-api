package dealerconnector

import (
	"context"
	"fmt"
	"models"
	"models/dealer"
	"shareerrors"
	"shareutils"
)

func (i impl) FindDealers(ctx context.Context, dealerCode string) (*[]dealer.Dealer, error) {
	//reqBody := dealer.ListDealerRequest{
	//	FTCustCode: dealerCode,
	//	Pagination: models.Pagination{
	//		Offset: 0,
	//		Limit:  50,
	//	},
	//}
	reqBody := models.ListRequest[dealer.ListDealerRequest]{
		Criteria: &dealer.ListDealerRequest{
			FTCustCode: dealerCode,
		},
		Pagination: models.Pagination{
			Offset: 0,
			Limit:  50,
		},
	}
	resp, err := shareutils.Post[models.ListRequest[dealer.ListDealerRequest], []dealer.Dealer](ctx, fmt.Sprintf("%v/api/v1/dealers/internal", i.url), &reqBody)
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp.Data)
	if !resp.IsSuccess() {
		return nil, shareerrors.NewError(resp.StatusCode, resp.Message).SetData(resp.Data)
	}
	//fmt.Println(resp.Data)
	return resp.Data, nil
}
