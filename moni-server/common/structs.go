package common

import (
	"fmt"
	"strings"
)

type QueryString struct {
	Id     int64  `json:"id" form:"id"`
	Type   int    `json:"type" form:"type"`
	Page   int    `json:"page" form:"page"`
	Limit  int    `json:"limit" form:"limit"`
	Search string `json:"search" form:"search"`
}
type Handle struct {
	Id string `json:"id" form:"id"`
}

func (q *QueryString) CheckParam() bool {
	if q.Limit == 0 || q.Page == 0 {
		return false
	}
	q.Search = strings.TrimSpace(q.Search)
	return true
}
func (q QueryString) String() string {
	return fmt.Sprintf("{ type: %v,page: %v,limit: %v,search: %v }", q.Type, q.Page, q.Limit, q.Search)
}

type LayuiTable struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}
