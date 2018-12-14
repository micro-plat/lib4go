package tpl

import (
	"fmt"
	"testing"
)

func TestTplxx(t *testing.T) {
	sql, _, _ := AnalyzeTPL("where &id", map[string]interface{}{
		"id": 1,
	}, func() string {
		return "1"
	})
	fmt.Println("sql:", sql)
}

func TestORCLContext(t *testing.T) {
	context, err := GetDBContext("ORACLE")
	if err != nil || context == nil {
		t.Error("GetDBContext返回结果有误")
	}

	context, err = GetDBContext("oracle")
	if err != nil || context == nil {
		t.Error("GetDBContext返回结果有误")
	}

	context, err = GetDBContext("sqlite")
	if err != nil || context == nil {
		t.Error("GetDBContext返回结果有误")
	}

	context, err = GetDBContext("SQLITE")
	if err != nil || context == nil {
		t.Error("GetDBContext返回结果有误")
	}

	context, err = GetDBContext("mysql2")
	if err == nil || context != nil {
		t.Error("GetDBContext返回结果有误")
	}

	/*add by champly 2016年11月9日11:54:20*/
	// 输入不同的字符
	context, err = GetDBContext("#@！%￥%")
	if err == nil || context != nil {
		t.Error("GetDBContext返回结果有误")
	}

	// 输入空字符
	context, err = GetDBContext("")
	if err == nil || context != nil {
		t.Error("GetDBContext返回结果有误")
	}
	/*end*/
}
