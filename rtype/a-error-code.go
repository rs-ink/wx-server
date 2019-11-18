package rtype

type ErrCode int

const (
	_ ErrCode = -1 * iota
	ErrCodeNotLogin
	ErrCodeNotRole
	ErrCodeServiceError
	ErrCodeSystemError
	ErrCodeParamsError
	ErrCodeSystemUndefineError
)

//基础错误码定义
const orderBaseErrCode = 1000
const activityBaseErrCode = 2000
const mediaBaseErrCode = 3000
const sourceBaseErrCode = 4000
const wxBaseErrCode = 5000
const operatorBaseErrCode = 6000

func (ec ErrCode) Result() *Result {
	res := Error(ec, "")
	return res
}
