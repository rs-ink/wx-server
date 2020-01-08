package base

import "wx-server/config"

type PageRequest struct {
	Start int `json:"start" info:"起始位置"`
	Limit int `json:"limit" info:"每页条数"`
}

type PageResponse struct {
	PageRequest
	Total int `json:"total"`
}

func (page *PageRequest) CheckMaxPageRequest(maxCurrents ...int) {
	var max int
	if len(maxCurrents) > 0 {
		max = maxCurrents[0]
	} else {
		max = config.Cfg().Sql.MaxPageSize
	}
	if page.Start < 0 {
		page.Start = 0
	}
	if page.Limit > max || page.Limit <= 0 {
		page.Limit = max
	}
}
