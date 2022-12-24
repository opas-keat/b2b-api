package shareerrors

import (
	"fmt"
	"models/status_code"
)

// Error ...
type Error struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// NewError ...
func NewError(statusCode int, message ...string) *Error {
	err := &Error{
		StatusCode: statusCode,
		Message:    status_code.StatusCodeMessage(statusCode),
	}
	if len(message) > 0 {
		err.Message = message[0]
	}
	return err
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v:%v", e.StatusCode, e.Message)
}

func (e *Error) SetData(data interface{}) *Error {
	e.Data = data
	return e
}

type FieldError struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

func IsCode(err error, code int) bool {
	switch v := err.(type) {
	case *Error:
		return v.StatusCode == code
	default:
		return false
	}
}
