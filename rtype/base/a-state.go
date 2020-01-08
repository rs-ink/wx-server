package base

type StateBase int

const (
	StateBaseInit StateBase = iota
	StateEnable
	StateDisable
	StateDel
)

type StateInfo struct {
	State     StateBase `json:"state" info:"状态"`
	StateDesc string    `json:"stateDesc" info:"状态描述"`
}

func (state StateBase) Info(desc string) StateInfo {
	return StateInfo{
		State:     state,
		StateDesc: desc,
	}
}
