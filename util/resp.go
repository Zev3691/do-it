package util

import "re_new/util/errorsx"

var Nil = struct{}{}

type response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(msg string, data interface{}) *response {
	if msg == "" {
		msg = "success"
	}
	return &response{
		Code: int(errorsx.ErrSuccess),
		Msg:  msg,
		Data: data,
	}
}

func SuccessNIL() *response {
	return &response{
		Code: 0,
		Msg:  "success",
		Data: Nil,
	}
}

func Response(errCode errorsx.ErrCode, msg ...errorsx.CustomErrMsg) *response {
	errx := errorsx.New(errCode, msg...)
	return &response{
		Code: int(errx.ErrCode),
		Msg:  errx.ErrMsg,
		Data: Nil,
	}
}
