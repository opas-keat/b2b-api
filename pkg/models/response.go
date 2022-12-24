package models

import (
	"context"
	"fmt"
	"models/status_code"
	"regexp"
	"strconv"
	"strings"
)

const ServiceNameContextKey = "servicename"

var regexResponseCode, _ = regexp.Compile(`\d{3}$`)

type Response[T any] struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       *T     `json:"data,omitempty"`
}

func (r *Response[T]) GetData() T {
	return *r.Data
}

func (r *Response[T]) MetaToString() string {
	return fmt.Sprintf(
		"status_code:%d  code:%s  message:%s ", r.StatusCode, r.Code, r.Message)
}

func (Response[T]) Success(ctx context.Context, data *T) *Response[T] {
	statusCode := status_code.Success
	serviceName, _ := ctx.Value(ServiceNameContextKey).(string)
	return &Response[T]{
		Code:       fmt.Sprintf("%v_%v", strings.ToUpper(serviceName), statusCode),
		StatusCode: statusCode,
		Message:    status_code.StatusCodeMessage(statusCode),
		Data:       data,
	}
}

func (r Response[T]) IsSuccess() bool {
	return regexResponseCode.FindString(r.Code) == strconv.Itoa(status_code.Success)
}
