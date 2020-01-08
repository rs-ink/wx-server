package rtype

import (
	"encoding/json"
	"github.com/rs-ink/rslog"
	"wx-server/rlog"
)

type Result struct {
	Code   ErrCode     `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Ext    interface{} `json:"ext,omitempty"`
	pcInfo rslog.PcInfo
}

func Success() *Result {
	return &Result{
		Msg: "SUCCESS",
	}
}
func Error(code ErrCode, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}
func (result *Result) GetPcInfo() rslog.PcInfo {
	return result.pcInfo
}

func (result *Result) SetExt(ext interface{}) *Result {
	result.Ext = ext
	return result
}

func (result *Result) SetMsg(msg string) *Result {
	result.Msg = msg
	return result
}

func (result *Result) SetData(data ...interface{}) *Result {
	if len(data) == 1 {
		result.Data = data[0]
	} else {
		result.Data = data
	}
	return result
}

func (result *Result) Error() string {
	data, _ := json.Marshal(*result)
	return string(data)
}

func (result *Result) RuntimeError() {
	rlog.Error("==============================")
	rlog.ErrorF("%+v", *result)
	rlog.Error("==============================")
}
