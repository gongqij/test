package main

import "fmt"

func main() {
	err1 := NewRespError(ServiceParasError, "abc")
	err2 := NewRespError(ServiceParasError, "abc")
	if err1 == err2 {
		fmt.Println("两个错误相同")
	}
}

func (p *ErrorResp) Error() string {
	return p.ErrorMsg
}

func NewRespError(code ErrCode, msg string) error {
	return &ErrorResp{
		ErrorID:  code,
		ErrorMsg: msg,
	}
}

// ErrorResp ...
type ErrorResp struct {
	ErrorID  ErrCode `json:"errorid"`
	ErrorMsg string  `json:"errormsg"`
}

type ErrCode string

const (
	ServiceParasError ErrCode = "100000" // HTTP请求参数不正确
)
