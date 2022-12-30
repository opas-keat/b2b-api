package user

import (
	"context"
	shareModels "models/status_code"
	"models/user"
	"shareerrors"
	"user/entities"
	"user/utils"
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

	return &user.VerifyEmailResponse{
		Message: "Email verified successfully",
	}, nil
}
