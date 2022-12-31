package models

import "strings"

type (
	Mode        string
	ServiceName string
	DealerURL   string
	//WalletURL      string
	//MemberURL      string
	//ExternalURL    string
	//ReportURL      string
	//ConfigURL      string
	//JaegerEndpoint string
	//SettingGameURL string
	//ClientOrigin string
)

func (m Mode) String() string {
	return string(m)
}

func (s ServiceName) String() string {
	return string(s)
}

func (d DealerURL) String() string {
	return string(d)
}

//func (c ClientOrigin) String() string {
//	return string(c)
//}

//func (w WalletURL) String() string {
//	return string(w)
//}
//
//func (m MemberURL) String() string {
//	return string(m)
//}
//
//func (o ExternalURL) String() string {
//	return string(o)
//}
//
//func (r ReportURL) String() string {
//	return string(r)
//}
//
//func (r SettingGameURL) String() string {
//	return string(r)
//}
//
//func (c ConfigURL) String() string {
//	return string(c)
//}
//
//func (j JaegerEndpoint) String() string {
//	return string(j)
//}

func (m Mode) IsLocal() bool {
	return strings.ToLower(m.String()) == "local"
}
