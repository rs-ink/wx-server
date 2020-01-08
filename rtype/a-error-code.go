package rtype

type ErrCode int

const (
	_                          ErrCode = -1 * iota
	ErrCodeNotLogin                    //未登录
	ErrCodeNotRole                     //无权限
	ErrCodeServiceError                //服务异常
	ErrCodeSystemError                 //系统异常
	ErrCodeParamsError                 //参数异常
	ErrCodeSystemUndefineError         //系统未定义异常
	ErrCodeInfoExist                   //信息已存在
	ErrCodeInfoNotExist                //信息未存在
)

//基础错误码定义
const OrderBaseErrCode = 1000
const ActivityBaseErrCode = 2000
const MediaBaseErrCode = 3000
const SourceBaseErrCode = 4000
const WxBaseErrCode = 5000
const OperatorBaseErrCode = 6000

func (ec ErrCode) Result() *Result {
	res := Error(ec, "")
	return res
}
