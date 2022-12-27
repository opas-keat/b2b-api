package user

type MeResponse struct {
	ID         string `json:"id"`
	RoleName   string `json:"role_name"`
	Email      string `json:"email"`
	DealerCode string `json:"dealer_code"`
	//Birthdate           *string                `json:"birthdate,omitempty"`
	//Phone               string                 `json:"phone"`
	//MasterID            *string                `json:"master_id,omitempty"`
	//AgentID             *string                `json:"agent_id,omitempty"`
	//AssistantPermission map[string]interface{} `json:"assistant_permission,omitempty"`
	//BookBanks           []MeBookBank           `json:"book_banks"`
	//ReferralCode        string                 `json:"referralCode"`
	//CreateType          string                 `json:"create_type"`
	//InviterMemberID     *string                `json:"inviter_member_id,omitempty"`
}
