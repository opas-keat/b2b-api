package entities

import (
	"log"
)

type Response struct {
	Status    string   `json:"status"`
	Desc      string   `json:"desc"`
	Balance   *float64 `json:"balance,omitempty"`
	BalanceTs *string  `json:"balanceTs,omitempty"`
	UserID    *string  `json:"userId,omitempty"`
}

func (receiver *Response) Create(statusCode string, desc error) *Response {
	log.Println(desc)
	if desc != nil {
		receiver.Desc = desc.Error()
	} else {
		receiver.Desc = ErrorCodeMessageTable.GetMessage(statusCode)
	}
	receiver.Status = statusCode
	return receiver
}

//func (receiver *Response) SetBalance(balance decimal.Decimal) *Response {
//	_b, _ := balance.RoundBank(3).Float64()
//	receiver.Balance = pointer.ToFloat64(_b)
//	receiver.BalanceTs = pointer.ToString(time.Now().Format("2006-01-02T15:04:05.000-07:00"))
//	return receiver
//}

//func (receiver *Response) SetUserID(memberID string) *Response {
//	receiver.UserID = pointer.ToString(memberID)
//	return receiver
//}

type ErrorCodeTable map[string]string

var ErrorCodeMessageTable = ErrorCodeTable{
	"9998": "System Busy",
	"9999": "Fail",
	"0000": "Success",
	//"11":   "Do not have this platform under your agent.",
	//"1000": "Invalid user Id",
	//"1001": "Account existed",
	//"1002": "Account is not exists",
	//"1004": "Invalid Currency",
	//"1005": "language is not exists",
	//"1006": "PT Setting is empty!",
	//"1007": "Invalid PT setting with parent!",
	//"1008": "Invalid token!",
	//"1009": "Invalid timeZone",
	//"1010": "Invalid amount",
	//"1011": "Invalid txCode",
	//"1012": "Has Pending Transfer",
	//"1013": "Account is Lock",
	//"1014": "Account is Suspend",
	//"1016": "TxCode already operation!",
	//"1017": "TxCode is not exist",
	//"1018": "Not Enough Amount",
	//"1019": "No Data",
	//"1024": "Invalid date time format",
	//"1025": "Invalid transaction status",
	//"1026": "Invalid bet limit setting",
	//"1027": "Invalid Certificate",
	//"1028": "Unable to proceed. please try again later.",
	//"1029": "invalid IP address.It might be IP address did not whitelist yet, or AWC cannot identify your AgentID",
	//"1030": "invalid Device to call API.(Ex.IE)",
	//"1031": "System is under maintenance.",
	//"1032": "Duplicate login.",
	//"1033": "Invalid gameCode",
	//"1034": "Time does not meet.",
	//"1035": "Invalid Agent Id.",
	//"1036": "Invalid parameters.",
	//"1037": "Invalid customer setting.It might be your target CallbackURL not yet passed AWC test, please contact us for assistant.",
	//"1038": "Duplicate transaction.",
	//"1039": "Transaction not found.",
	//"1040": "Request timeout.",
	//"1041": "HTTP Status error.",
	//"1042": "HTTP Response is empty.",
	//"1043": "Bet has canceled.",
	//"1044": "Invalid bet.",
	//"1045": "Add account statement failed.",
	//"1046": "Transfer Failed! Please contact customer support immediately. Sorry for any inconvenience caused.",
	//"1047": "Game is under maintenance.",
}

func (receiver ErrorCodeTable) GetMessage(code string) string {
	if message, ok := receiver[code]; ok {
		return message
	}
	return "undefined"
}

const (
	//SystemBusy                          = "9998"
	Fail = "9999"
	//Success                             = "0000"
	//Donothavethisplatformunderyouragent = "11"
	//InvaliduserId                       = "1000"
	//Accountexisted                      = "1001"
	//Accountisnotexists                  = "1002"
	//InvalidCurrency                     = "1004"
	//languageisnotexists                 = "1005"
	//PTSettingisempty                    = "1006"
	//InvalidPTsettingwithparent          = "1007"
	//Invalidtoken                        = "1008"
	//InvalidtimeZone                     = "1009"
	//Invalidamount                       = "1010"
	//InvalidtxCode                       = "1011"
	//HasPendingTransfer                  = "1012"
	//AccountisLock                       = "1013"
	//AccountisSuspend                    = "1014"
	//TxCodealreadyoperation              = "1016"
	//TxCodeisnotexist                    = "1017"
	//NotEnoughBalance                    = "1018"
	//NoData                              = "1019"
	//Invaliddatetimeformat               = "1024"
	//Invalidtransactionstatus            = "1025"
	//Invalidbetlimitsetting              = "1026"
	//InvalidCertificate                  = "1027"
	//Unabletoproceed                     = "1028"
	//invalidIPaddress                    = "1029"
	//invalidDevicetocallAPI              = "1030"
	//Systemisundermaintenance            = "1031"
	//Duplicatelogin                      = "1032"
	//InvalidgameCode                     = "1033"
	//Timedoesnotmeet                     = "1034"
	//InvalidAgentId                      = "1035"
	Invalidparameters = "1036"
	//Invalidcustomersetting              = "1037"
	//Duplicatetransaction                = "1038"
	//Transactionnotfound                 = "1039"
	//Requesttimeout                      = "1040"
	//HTTPStatuserror                     = "1041"
	//HTTPResponseisempty                 = "1042"
	//Bethascanceled                      = "1043"
	//Invalidbet                          = "1044"
	//Addaccountstatementfailed           = "1045"
	//TransferFailed                      = "1046"
	//Gameisundermaintenance              = "1047"
)
