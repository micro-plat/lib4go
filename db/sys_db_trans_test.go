package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/qxnw/lib4go/sysinfo/memory"
)

// const dbConnectStr = "oc_test/123456@orcl136"

func TestDBTEST1(t *testing.T) {
	obj, err := NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	// 正常测试
	for i := 0; i < 500000; i++ {
		sql := "select * from test where id = :1"
		args := []interface{}{"1"}
		dataRows, colus, err := obj.Query(sql, args...)
		if err != nil {
			t.Errorf("执行%s失败：%v", sql, err)
		}
		if dataRows == nil {
			t.Errorf("执行%s失败", sql)
			t.FailNow()
		}
		if dataRows[0][colus[0]] != "1" {
			t.Errorf("执行%s失败", sql)
		}
		//	}
		if i%1000 == 0 {
			fmt.Printf("%d:%+v BenchmarkEngine1:%+v\n", i, time.Now(), memory.GetInfo().Used)
		}
	}
	fmt.Println("执行完成")
	for {
		select {
		case <-time.After(time.Second * 10):
			fmt.Printf("%+v BenchmarkEngine1:%+v\n", time.Now(), memory.GetInfo().Used)
		}
	}
	time.Sleep(time.Hour)

}
func TestDBTRansQuery(t *testing.T) {
	obj, err := NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	dbTrans, err := obj.Begin()
	if dbTrans == nil || err != nil {
		t.Errorf("创建数据库事务失败:%v", err)
	}

	// 正常测试
	sql := "select * from test where id = :1"
	args := []interface{}{"1"}
	dataRows, colus, err := dbTrans.Query(sql, args...)
	if err != nil {
		t.Errorf("执行%s失败：%v", sql, err)
	}
	if dataRows == nil {
		t.Errorf("执行%s失败", sql)
	}
	if dataRows[0][colus[0]] != "1" {
		t.Errorf("执行%s失败", sql)
	}

	// 数据库连接串错误测试
	obj, err = NewSysDB("oracle", "", 2)
	if obj != nil || err == nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "select * from test where id = :1"
		args = []interface{}{"1"}
		dataRows, colus, err = dbTrans.Query(sql, args...)
		if err != nil {
			t.Errorf("执行%s失败：%v", sql, err)
		}
		if dataRows == nil {
			t.Errorf("执行%s失败", sql)
		}
		if dataRows[0][colus[0]] != "1" {
			t.Errorf("执行%s失败", sql)
		}
	}

	// 数据库连接串错误测试
	obj, err = NewSysDB("", dbConnectStr, 2)
	if obj != nil || err == nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "select * from test where id = :1"
		args = []interface{}{"1"}
		dataRows, colus, err = dbTrans.Query(sql, args...)
		if err != nil {
			t.Errorf("执行%s失败：%v", sql, err)
		}
		if dataRows == nil {
			t.Errorf("执行%s失败", sql)
		}
		if dataRows[0][colus[0]] != "1" {
			t.Errorf("执行%s失败", sql)
		}
	}

	// sql错误
	obj, err = NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	sql = "selects * from dual where 1 = :1"
	args = []interface{}{"1"}
	dataRows, colus, err = dbTrans.Query(sql, args...)
	if err == nil {
		t.Errorf("执行%s失败", sql)
	}

	// sql错误
	obj, err = NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	sql = "select * from id where 1 = :1"
	args = []interface{}{"1"}
	dataRows, colus, err = dbTrans.Query(sql, args...)
	if err == nil {
		t.Errorf("执行%s失败", sql)
	}
}

func TestDBTRansExecute(t *testing.T) {
	obj, err := NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	dbTrans, err := obj.Begin()
	if dbTrans == nil || err != nil {
		t.Errorf("创建数据库事务失败:%v", err)
	}

	// 正常测试
	sql := "update test t set t.money = t.money + 0 where t.id = :1"
	args := []interface{}{"1"}
	row, err := dbTrans.Execute(sql, args...)
	if err != nil {
		t.Errorf("执行%s失败：%v", sql, err)
	}
	if int(row) != 1 {
		t.Errorf("执行%s失败", sql)
	}

	// 数据库连接串错误测试
	obj, err = NewSysDB("oracle", "", 2)
	if obj != nil || err == nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "update test t set t.money = t.money + 0 where t.id = :1"
		args = []interface{}{"1"}
		row, err = dbTrans.Execute(sql, args...)
		if err != nil {
			t.Errorf("执行%s失败：%v", sql, err)
		}
		if int(row) == 1 {
			t.Errorf("执行%s失败", sql)
		}
	}

	// 数据库连接串错误测试
	obj, err = NewSysDB("", dbConnectStr, 2)
	if obj != nil || err == nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "update test t set t.money = t.money + 0 where t.id = :1"
		args = []interface{}{"1"}
		row, err = dbTrans.Execute(sql, args...)
		if err != nil {
			t.Errorf("执行%s失败：%v", sql, err)
		}
		if int(row) == 1 {
			t.Errorf("执行%s失败", sql)
		}
	}

	// sql错误
	obj, err = NewSysDB("oracle", dbConnectStr, 2)
	if err != nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "updates test t set t.money = t.money + 0 where t.id = :1"
		args = []interface{}{"1"}
		row, err = dbTrans.Execute(sql, args...)
		if err == nil {
			t.Errorf("测试失败")
		}
	}

	// sql错误
	obj, err = NewSysDB("oracle", dbConnectStr, 2)
	if err != nil {
		t.Error("创建数据库连接失败:", err)
	}
	if obj != nil {
		sql = "update tests t set t.money = t.money + 0 where t.id = :1"
		args = []interface{}{"1"}
		row, err = dbTrans.Execute(sql, args...)
		if err == nil {
			t.Errorf("测试失败")
		}
	}
}

func TestDBTransRollback(t *testing.T) {
	// 正常测试
	obj, err := NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	dbTrans, err := obj.Begin()
	if dbTrans == nil || err != nil {
		t.Errorf("创建数据库事务失败:%v", err)
	}
	err = dbTrans.Rollback()
	if err != nil {
		t.Error("回滚数据库事务失败")
	}

	// // 数据库连接串错误
	// obj, err = NewSysDB("oracle", "", 2, 2)
	// if obj != nil || err == nil {
	// 	t.Error("创建数据库连接失败:", err)
	// }

	// err = dbTrans.Rollback()
	// if err != nil {
	// 	t.Error("回滚数据库事务失败")
	// }
}

func TestDBTransCommit(t *testing.T) {
	obj, err := NewSysDB("oracle", dbConnectStr, 2)
	if obj == nil || err != nil {
		t.Error("创建数据库连接失败:", err)
	}

	dbTrans, err := obj.Begin()
	if dbTrans == nil || err != nil {
		t.Errorf("创建数据库事务失败:%v", err)
	}

	err = dbTrans.Commit()
	if err != nil {
		t.Error("提交数据库事务失败")
	}

	// // 数据库连接串错误
	// obj, err = NewSysDB("oracle", "",2)
	// if obj != nil || err == nil {
	// 	t.Error("创建数据库连接失败:", err)
	// }

	// err = dbTrans.Commit()
	// if err != nil {
	// 	t.Error("回滚数据库事务失败")
	// }
}
