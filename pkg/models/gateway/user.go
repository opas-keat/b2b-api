package gateway

import (
	"github.com/AlekSi/pointer"
	"models"
)

type User struct {
	Username         string      `json:"username"`
	AgentID          *string     `json:"agent_id"`
	MasterID         *string     `json:"master_id"`
	ReferrerMemberID string      `json:"referrer_member_id"`
	IPAddress        string      `json:"ip_address"`
	MemberID         string      `json:"memberId"`
	Role             models.Role `json:"role"`
	WalletID         string      `json:"wallet_id"`
	RoleName         string      `json:"role_name"`
}

func (u User) GetActionBy() string {
	switch u.Role {
	case models.AdminAssistant:
		return models.AdminMemberID
	case models.MasterAssistant, models.AgentAssistant:
		return u.ReferrerMemberID
	default:
		return u.MemberID
	}
}

func (u User) GetAssistanceActionBy() *string {
	if u.IsAssistant() {
		return pointer.ToString(u.MemberID)
	}
	return nil
}

func (u User) IsAssistant() bool {
	return u.Role == models.AdminAssistant || u.Role == models.MasterAssistant || u.Role == models.AgentAssistant
}

func (u User) IsAdminRef() bool {
	return u.Role == models.Admin || u.Role == models.AdminAssistant
}

func (u User) IsMasterRef() bool {
	return u.Role == models.Master || u.Role == models.MasterAssistant
}

func (u User) IsAgentRef() bool {
	return u.Role == models.Agent || u.Role == models.AgentAssistant
}

func (u User) IsMember() bool {
	return u.Role == models.Member
}
