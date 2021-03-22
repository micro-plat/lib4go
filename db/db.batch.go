package db

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//executeBatch  批量执行SQL语句
func executeBatch(db IDBExecuter, sqls []string, input map[string]interface{}) (QueryRows, error) {
	if len(sqls) == 0 {
		return nil,
			fmt.Errorf("未传入任何SQL语句")
	}
	ninput := types.XMap(input)
	output := make(types.XMaps, 0, 0)
	for i, sql := range sqls {

		if len(sql) < 6 {
			return nil, fmt.Errorf("sql语句错误%s", sql)
		}
		prefix := strings.Trim(strings.TrimSpace(strings.TrimLeft(sql, "\n")), "\t")[:6]
		switch strings.ToUpper(prefix) {
		case "SELECT":
			output, err := db.Query(sql, ninput.ToMap())
			if err != nil {
				return nil, err
			}
			if output.Len() == 0 {
				return nil, fmt.Errorf("%s数据不存在%w input:%+v", sql, errs.ErrNotExist, input)
			}
			ninput.Merge(output.Get(0))
			if i == len(sqls)-1 && output.Len() > 1 {
				return output, nil
			}
		case "UPDATE", "INSERT":
			rows, err := db.Execute(sql, ninput.ToMap())
			if err != nil {
				return nil, err
			}
			if rows == 0 {
				return nil, fmt.Errorf("%s数据修改失败%w input:%+v", sql, errs.ErrNotExist, input)
			}
		default:
			return nil, fmt.Errorf("不支持的SQL语句，或SQL语句前包含有特殊字符:%s", sql)
		}
	}
	output.Append(ninput)
	return &output, nil
}
