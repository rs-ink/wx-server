package wxpay

import "wx-server/rtype/base"

//SENDING:发放中
//SENT:已发放待领取
//FAILED：发放失败
//RECEIVED:已领取
//RFUND_ING:退款中
//REFUND:已退款
type RedPackStatus string

func (w RedPackStatus) String() string {
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
	WxRedPackSending  RedPackStatus = "SENDING"
	WxRedPackSent                   = "SENT"
	WxRedPackFailed                 = "FAILED"
	WxRedPackReceived               = "RECEIVED"
	WxRedPackRFundIng               = "RFUND_ING"
	WxRedPackRefund                 = "REFUND"
)

type PayOrder struct {
	Body           string `json:"body"`
	OutTradeNo     string `json:"out_trade_no"`
	SpBillCreateIp string `json:"spbill_create_ip"`
	TotalFee       int64  `json:"total_fee"`
	OpenId         string `json:"openid"`
	NotifyUrl      string `json:"notify_url"`
}

type RedPack struct {
	base.OrderID
	OpenId      string `xml:"openId"`
	TotalAmount int64  `xml:"totalAmount"`
	TotalNum    int64  `xml:"totalNum"`
}

type ResultBase struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

type ResultRedPackInfo struct {
	ResultBase
	MchBillNo   string        `xml:"mch_billno"`
	MchId       string        `xml:"mch_id"`
	DetailId    string        `xml:"detail_id"`
	Status      RedPackStatus `xml:"status"`
	SendType    string        `xml:"send_type"`
	HbType      string        `xml:"hb_type"`
	TotalNum    string        `xml:"total_num"`
	TotalAmount string        `xml:"total_amount"`
	SendTime    string        `xml:"send_time"`
}

type ResultSendRedPack struct {
	ResultBase
	TotalAmount int64  `xml:"total_amount"`
	ReOpenid    int64  `xml:"re_openid"`
	WxAppId     int64  `xml:"wxappid"`
	MchId       int64  `xml:"mch_id"`
	MchBillNo   string `xml:"mch_billno"`
}

func NewRedPack() RedPack {
	return RedPack{
		OrderID: base.NewOrderId(),
	}
}
