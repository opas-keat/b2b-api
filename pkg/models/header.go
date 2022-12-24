package models

const (
	HeaderXUsername  = "X-Username"
	HeaderXMemberID  = "X-Member-ID"
	HeaderXRole      = "X-Role"
	HeaderXAgentID   = "X-Agent-ID"
	HeaderXMasterID  = "X-Master-ID"
	HeaderIPAddress  = "X-IP-Address"
	HeaderReferrer   = "X-Referrer"
	HeaderXRequestID = "X-Request-ID"

	HeaderXInternalCallerID = "X-Internal-Caller-ID"
)

var ArrHeader = []string{
	HeaderXUsername,
	HeaderXMemberID,
	HeaderXRole,
	HeaderXAgentID,
	HeaderXMasterID,
	HeaderIPAddress,
	HeaderReferrer,
}
