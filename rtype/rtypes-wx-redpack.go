package rtype

import (
	"wx-server/rlog"
	"wx-server/rtype/base"
)

func init() {
	err := db.CreateTables(&WxRedPack{})
	rlog.CheckShowError(err)
}

type WxRedPack struct {
	ID           `xorm:"extends"`
	BaseTime     `xorm:"extends"`
	ActivityId   int             `json:"activityId" xorm:"default 0" info:"活动ID"`
	MchBillNo    base.IDuInt64   `json:"mch_billno" info:"原始单号"`
	MchId        string          `json:"mch_id" info:"商户号"`
	SendListId   string          `json:"send_listid"`
	TotalAmount  int64           `json:"total_amount"`
	TotalNum     int64           `json:"total_num"`
	ReOpenId     string          `json:"re_openid"`
	ReCustomerId int             `json:"re_customer_id"`
	WxAppId      string          `json:"wxappid" info:"公众账号appid"`
	ErrCode      string          `json:"err_code" info:"错误代码"`
	ErrCodeDes   string          `json:"err_code_des" info:"错误代码描述"`
	ResultCode   string          `json:"return_code" info:"业务结果"`
	ReturnCode   string          `json:"result_code" info:"返回状态码"`
	ReturnMsg    string          `json:"return_msg" info:"返回信息"`
	Version      int             `json:"version" xorm:"version int(11) default 0 version"`
	Status       WxRedPackStatus `json:"status"`
}

//SENDING:发放中
//SENT:已发放待领取
//FAILED：发放失败
//RECEIVED:已领取
//RFUND_ING:退款中
//REFUND:已退款
type WxRedPackStatus string

func (w WxRedPackStatus) String() string {
	switch w {
	case WxRedPackSending:
		return "发放中"
	case WxRedPackSent:
		return "已发放待领取"
	case WxRedPackFailed:
		return "发放失败"
	case WxRedPackReceived:
		return "已领取"
	case WxRedPackRFundIng:
		return "退款中"
	case WxRedPackRefund:
		return "已退款"
	default:
		return string(w)
	}
}

const (
	WxRedPackSending  WxRedPackStatus = "SENDING"
	WxRedPackSent                     = "SENT"
	WxRedPackFailed                   = "FAILED"
	WxRedPackReceived                 = "RECEIVED"
	WxRedPackRFundIng                 = "RFUND_ING"
	WxRedPackRefund                   = "REFUND"
)

func (WxRedPack) TableName() string {
	return "tb_wx_red_pack"
}

func (red WxRedPack) StatusString() string {
	if red.Status != "" {
		return red.Status.String()
	}
	if red.ErrCode == "SUCCESS" {
		return "发放成功"
	} else {
		return red.ErrCodeDes
	}
}
