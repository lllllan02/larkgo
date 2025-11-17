package core

import (
	"strconv"
	"strings"
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  *struct {
		Details              []*CodeErrorDetail              `json:"details,omitempty"`
		PermissionViolations []*CodeErrorPermissionViolation `json:"permission_violations,omitempty"`
		FieldViolations      []*CodeErrorFieldViolation      `json:"field_violations,omitempty"`
	} `json:"error"`
}

func (ce CodeError) Error() string {
	return ce.String()
}

func (ce CodeError) String() string {
	sb := strings.Builder{}
	sb.WriteString("msg:")
	sb.WriteString(ce.Msg)
	sb.WriteString(",code:")
	sb.WriteString(strconv.Itoa(ce.Code))
	return sb.String()
}

type CodeErrorDetail struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type CodeErrorPermissionViolation struct {
	Type        string `json:"type,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
}

type CodeErrorFieldViolation struct {
	Field       string `json:"field,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

// 非法参数错误
type IllegalParamError struct {
	msg string
}

func (err *IllegalParamError) Error() string {
	return err.msg
}

// 客户端超时错误
type ClientTimeoutError struct {
	msg string
}

func (err *ClientTimeoutError) Error() string {
	return err.msg
}

// 连接失败错误
type DialFailedError struct {
	msg string
}

func (err *DialFailedError) Error() string {
	return err.msg
}

// 服务器超时错误
type ServerTimeoutError struct {
	msg string
}

func (err *ServerTimeoutError) Error() string {
	return err.msg
}
