package db

import (
	"database/sql"
	"reflect"

	"github.com/micro-plat/lib4go/types"
)

// QueryRow 单行数据
type QueryRow = types.XMap

// NewQueryRow 构建QueryRow对象
func NewQueryRow(len ...int) QueryRow {
	return types.NewXMap(len...)
}

// QueryRows 多行数据
type QueryRows = types.XMaps

// NewQueryRows 构建QueryRows
func NewQueryRows(len ...int) QueryRows {
	return types.NewXMaps(len...)
}

func resolveFullRows(rows *sql.Rows) (dataRows QueryRows, err error) {
	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := types.XMaps{}

	// 逐行解析
	for rows.Next() {
		currentRow, err := ResolveRow(rows, columns)
		if err != nil {
			return nil, err
		}
		// 将当前行添加到结果切片
		result = append(result, currentRow)
	}

	// 检查错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func ResolveRow(rows *sql.Rows, columns []string) (map[string]interface{}, error) {
	// 动态分配目标变量
	dest := make([]interface{}, len(columns))
	for i := range columns {
		dest[i] = new(string)
	}

	// 扫描数据
	if err := rows.Scan(dest...); err != nil {
		return nil, err
	}

	// 构建当前行的map
	currentRow := make(map[string]interface{})
	for i, col := range columns {
		val := reflect.ValueOf(dest[i]).Elem().Interface()
		currentRow[col] = val
	}

	return currentRow, nil
}
