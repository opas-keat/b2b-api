package user

import (
	"context"
	"errors"
	"fmt"
	"models/gateway"
	models "models/user"
	"user/constant"
	"user/entities"
	"user/utils"
)

func (s ServiceImpl) Me(ctx context.Context, userDetail gateway.User) (*models.MeResponse, error) {
	id, err := utils.ValidateToken(userDetail.MemberID, constant.PrivateKey)
	if err != nil {
		return nil, err
	}
	m, err := s.userRepo.Get(ctx, entities.User{ID: fmt.Sprint(id)})
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//meResponse := models.MeResponse{
	//	//MemberID:        member.MemberID,
	//	//RoleName:        string(member.RoleName),
	//	//Username:        member.Username,
	//	//Birthdate:       member.Birthdate,
	//	//Phone:           member.Phone,
	//	//MasterID:        member.MasterID,
	//	//AgentID:         member.AgentID,
	//	//ReferralCode:    member.ReferralCode,
	//	//InviterMemberID: member.InviterMemberID,
	//}
	return &models.MeResponse{
		ID:         m.ID,
		RoleName:   m.RoleName,
		Email:      m.Email,
		DealerCode: m.DealerCode,
	}, nil
}
