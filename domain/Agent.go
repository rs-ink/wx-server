package domain

import "wx-server/rtype/base"

//代理商
type Agent struct {
	AgentId int    `json:"agentId"`
	Name    string `json:"name"`
	OwnerBase
	base.StateInfo
}
