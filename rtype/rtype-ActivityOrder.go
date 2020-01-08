package rtype

import (
	"time"
	"wx-server/rlog"
	"wx-server/rtype/base"
	"wx-server/util"
)

func init() {
	err := db.CreateTables(&Order{})
	rlog.CheckShowError(err)
}

type OrderPayState int

const (
	StateOrderPayInit OrderPayState = iota
	StateOrderPayRunning
	StateOrderPaySuccess
	StateOrderPayCancel
	StateOrderPayFail
	StateOrderComplete
)

func (state OrderPayState) String() string {
	switch state {
	case StateOrderPayInit:
		return "初始化"
	case StateOrderPayRunning:
		return "支付中"
	case StateOrderPayCancel:
		return "取消支付"
	case StateOrderPaySuccess:
		return "支付成功"
	case StateOrderPayFail:
		return "支付失败"
	case StateOrderComplete:
		return "已核销"
	default:
		return "状态异常"
	}
}

const (
	_ ErrCode = OrderBaseErrCode + iota
	ErrOrderErrorSystem
	ErrOrderErrorNoCustomer
	ErrOrderErrorActivityRepeat
	ErrOrderErrorMaxQuantity
	ErrOrderIsConsume
	ErrOrderNotExist
	ErrOrderState
)

type Order struct {
	base.OrderID      `xorm:"extends"`
	BaseTime          `xorm:"extends"`
	OwnerId           int           `json:"ownerId"`
	CustomerId        int           `json:"customerId"`
	ActivityId        int           `json:"activityId"`
	ActivityName      string        `json:"activityName"`
	Price             int           `json:"price"`
	Phone             string        `json:"phone"`
	FromId            int           `json:"fromId"`
	WxFromChannel     FromChannel   `json:"wxFromChannel"`
	CarPlate          string        `json:"carPlate"`
	PayCode           string        `json:"payCode"`
	PayCreateTime     time.Time     `json:"payCreateTime"`
	PayUpdateTime     time.Time     `json:"payUpdateTime"`
	PayState          OrderPayState `json:"payState"`
	Version           int           `json:"version" xorm:"version int(11) default 0 version"`
	ConsumeBy         int           `json:"consumeBy"`
	ConsumeUpdateTime time.Time     `json:"consumeUpdateTime"`
}

type OrderPayBase struct {
	base.OrderID      `xorm:"extends"`
	CustomerId        int           `json:"customerId"`
	ActivityId        int           `json:"activityId"`
	Phone             string        `json:"phone"`
	CarPlate          string        `json:"carPlate"`
	PayCreateTime     time.Time     `json:"payCreateTime"`
	PayUpdateTime     time.Time     `json:"payUpdateTime"`
	PayState          OrderPayState `json:"payState"`
	ConsumeBy         int           `json:"consumeBy"`
	ConsumeUpdateTime time.Time     `json:"consumeUpdateTime"`
}

func (OrderPayBase) TableName() string {
	return "tb_order"
}

func (o *Order) YinCangPhone() (result string) {
	if len(o.Phone) == 11 {
		result = o.Phone[:3] + "******" + o.Phone[9:]
	} else {
		result = ""
	}
	return
}

func (o *Order) YincangCarPlate() (result string) {
	runes := []rune(o.CarPlate)
	for i := 0; i < len(runes); i++ {
		if i < 1 || i > 4 {
			result += string(runes[i])
		} else {
			result += "*"
		}
	}
	return
}

func (Order) TableName() string {
	return "tb_order"
}

func NewOrder(info CustomerInfo, activity ActivityEntity) *Order {
	return &Order{
		OrderID:       base.NewOrderId(),
		BaseTime:      NewBaseTime(),
		OwnerId:       activity.OwnerId,
		CustomerId:    info.Id,
		ActivityId:    activity.Id,
		ActivityName:  activity.ActivityName,
		Price:         activity.Price,
		PayCreateTime: time.Unix(0, 0),
		PayUpdateTime: time.Unix(0, 0),
	}
}

type OrderDaySummary struct {
	Day         string  `json:"day"`
	TotalAmount float64 `json:"totalAmount"`
	TotalCount  int64   `json:"totalCount"`
}

func (OrderDaySummary) TableName() string {
	return "tb_order"
}

func GetTotalSuccessCount(actId int) (count int64) {
	count, _ = db.Where("activity_id=? and (pay_state=? or pay_state=?)", actId, StateOrderPaySuccess, StateOrderComplete).Count(&Order{})
	return
}

func GetLastDaysOrderCount(actId, lastDays int) map[string]OrderDaySummary {
	start, end := util.GetBetweenAndTimes(lastDays)
	var amountList []OrderDaySummary
	session := db.NewSession()
	defer session.Close()
	_ = session.Where("pay_update_time between ? and ? and activity_id = ? and (pay_state=? or pay_state=?)", start, end, actId, StateOrderPaySuccess, StateOrderComplete).
		Select("sum(price)/100 total_amount, DATE_FORMAT(pay_update_time,'%Y-%m-%d') as  day ,count(1) as total_count").
		GroupBy(`DATE_FORMAT(pay_update_time,'%Y-%m-%d')`).
		Find(&amountList)
	result := make(map[string]OrderDaySummary)
	for _, amount := range amountList {
		result[string(amount.Day[5:])] = amount
	}
	return result
}
