package models

type Role string

const (
	Admin           Role = "ADMIN"
	AdminAssistant  Role = "ADMIN_ASSISTANT"
	Master          Role = "MASTER"
	MasterAssistant Role = "MASTER_ASSISTANT"
	Agent           Role = "AGENT"
	AgentAssistant  Role = "AGENT_ASSISTANT"
	Member          Role = "MEMBER"
)

const AdminMemberID = "u_admin"

var (
	AdminRole  = []Role{Admin, AdminAssistant}
	MasterRole = []Role{Master, MasterAssistant}
	AgentRole  = []Role{Agent, AgentAssistant}
	NonMember  = append(AdminRole, append(MasterRole, AgentRole...)...)
)

func (r Role) String() string {
	return string(r)
}

var MappingToRole = map[string]Role{
	"ADMIN":            Admin,
	"ADMIN_ASSISTANT":  AdminAssistant,
	"MASTER":           Master,
	"MASTER_ASSISTANT": MasterAssistant,
	"AGENT":            Agent,
	"AGENT_ASSISTANT":  AgentAssistant,
	"MEMBER":           Member,
}
