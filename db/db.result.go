package db

import (
	"github.com/micro-plat/lib4go/types"
)

//QueryRows 多行数据
type QueryRows = types.XMaps

//NewQueryRowsByJSON 根据json创建QueryRows
func NewQueryRowsByJSON(j string) (QueryRows, error) {
	return types.NewXMapsByJSON(j)
}

//NewQueryRows 构建QueryRows
func NewQueryRows(len ...int) QueryRows {
	return types.NewXMaps(len...)
}
