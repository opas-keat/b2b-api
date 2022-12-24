package models

import "strings"

type (
	Mode        string
	ServiceName string
)

func (m Mode) String() string {
	return string(m)
}

func (s ServiceName) String() string {
	return string(s)
}

func (m Mode) IsLocal() bool {
	return strings.ToLower(m.String()) == "local"
}
