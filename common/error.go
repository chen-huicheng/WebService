package common

import (
	"fmt"
)

var (
	PermissionDeny = NewHTTPError(403, "无权限")
	InvalidParam   = NewHTTPError(400, "参数错误")
	Unknown        = NewHTTPError(500, "内部错误")
	BizErr         = NewHTTPError(412, "业务异常")
)

type HTTPErr interface {
	Code() int
	Msg() string
	Error() string
}

type HTTPError struct {
	code   int
	msg    string
	srcErr error
}

func (e *HTTPError) Code() int {
	if e == nil {
		return 0
	}
	return e.code
}
func (e *HTTPError) Msg() string {
	if e == nil {
		return ""
	}
	return e.msg
}
func (e *HTTPError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[%d](%s)", e.code, e.msg)
}

func (e *HTTPError) Wrap(err error) *HTTPError {
	if e == nil || err == nil {
		return e
	}
	// 避免全局变量信息被修改
	newErr := e.clone()
	newErr.srcErr = err
	return newErr
}

func (e *HTTPError) SetMsg(msg string) *HTTPError {
	if e == nil {
		return e
	}
	// 避免全局变量信息被修改
	newErr := e.clone()
	newErr.msg = msg
	return newErr
}

func (e *HTTPError) clone() *HTTPError {
	if e == nil {
		return nil
	}

	return &HTTPError{
		code:   e.code,
		msg:    e.msg,
		srcErr: e.srcErr,
	}
}

func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{code: code, msg: msg}
}
