package main

import (
	"net/http"
	"encoding/json"
	"github.com/metooweb/gateway/xerrors"
)

type Resp struct {
	Data interface{} `json:"data"`
}

type RespError struct {
	Msg     string `json:"msg"`                //错误消息
	Code    int    `json:"code"`               //错误码
	SubMsg  string `json:"sub_msg,omitempty"`  //api错误消息
	SubCode string `json:"sub_code,omitempty"` //api错误码
}

type Response struct {
	Writer http.ResponseWriter
}

func (res *Response) outputError(err *RespError) {

	bytes, _ := json.Marshal(err)

	res.Writer.Write(bytes)

}

func (res *Response) IspError(code int, msg, subMsg string) {

	res.outputError(&RespError{
		Code:   code,
		Msg:    msg,
		SubMsg: subMsg,
	})

}

func (res *Response) IsvError(subCode, subMsg string) {

	res.outputError(&RespError{
		Code:    xerrors.ErrCodeRemoteServiceError,
		Msg:     "Remote service error",
		SubCode: subCode,
		SubMsg:  subMsg,
	})

}

func (res *Response) Output(val interface{}) {

	resp := Resp{Data: val}

	bytes, _ := json.Marshal(&resp)

	res.Writer.Write(bytes)
}
