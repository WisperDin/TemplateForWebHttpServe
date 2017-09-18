package controller

import (
	"net/http"
	"fmt"
	"strconv"
)

//处理排序相关的参数
func handleOrderParam(r *http.Request)string{
	var orderStr string
	orderType := r.FormValue("orderType")
	orderProp := r.FormValue("orderProp")
	if len(orderProp)<=0{
		return ""
	}
	//默认升序
	if len(orderType)<=0{
		orderType = "ASC"
	}
	orderStr = fmt.Sprintf(" ORDER BY %s %s ", orderProp, orderType)
	return orderStr
}

//处理分页相关的参数
func handlePageParam(r *http.Request)string{
	//默认参数
	skip := 0
	limit := 20

	var limitStr string
	page := r.FormValue("page")
	pageSize := r.FormValue("pageSize")

	pageInt, err := strconv.Atoi(page)
	if err == nil {
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err == nil {
			limit = pageSizeInt
			skip = limit * (pageInt - 1)
		}
	}
	limitStr = fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, skip)
	return limitStr
}
