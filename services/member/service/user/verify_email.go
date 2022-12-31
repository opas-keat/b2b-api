package user

import (
	"context"
	"member/entities"
	"member/utils"
	shareModels "models/status_code"
	"models/user"
	"shareerrors"
)

func (s ServiceImpl) VerifyEmail(ctx context.Context, verificationCode string) (*user.VerifyEmailResponse, error) {
	encodeVC := utils.Encode(verificationCode)
	updatedUser, err := s.userRepo.Get(ctx, entities.User{VerificationCode: encodeVC})
	if err != nil {
		return nil, shareerrors.NewError(shareModels.Duplicate, "Invalid verification code or user doesn't exists")
	}

	if updatedUser.Verified {
		return nil, shareerrors.NewError(shareModels.Duplicate, "User already verified")
	}

	em := entities.User{
		VerificationCode: "-",
		Verified:         true,
	}
	if err := s.userRepo.Update(ctx, updatedUser.ID, em); err != nil {
		return nil, shareerrors.NewError(shareModels.NotFound, "Update verification code fail")
	}

	dealerListResp, err := s.dealerConnector.FindDealers(ctx, updatedUser.DealerCode)
	if err != nil {
		return nil, err
	}
	//println(dealerListResp)

	var dealers []entities.Dealer
	for _, dealer := range *dealerListResp {
		dealers = append(dealers, entities.Dealer{
			DealerCode: dealer.Code,
			Address:    dealer.Address,
			Phone:      dealer.Phone,
			Name:       dealer.Name,
			LinkId:     dealer.ID,
		})
	}
	err = s.dealerRepo.CreateBatch(ctx, dealers)
	if err != nil {
		return nil, err
	}

	return &user.VerifyEmailResponse{
		Message: "ยืนยันอีเมลเรียบร้อย",
		//Message: "Email verified successfully",
	}, nil
}
