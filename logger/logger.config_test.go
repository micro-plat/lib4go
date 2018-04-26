package logger

import "testing"

// TestRead 测试读取配置文件的时候可能遇到的问题
func TestRead(t *testing.T) {
	// 配置文件不存在
	loggerPath = "../conf/no_ars.logger.json"
	_, err := read()
	t.Log(err)
	if err == nil {
		t.Error("test fail")
		return
	}

	// 无法读取配置文件,文件权限为000
	loggerPath = "../conf/without_x_ars.logger.json"
	_, err = read()
	t.Log(err)
	if err == nil {
		t.Error("test fail")
		return
	}

	// 配置文件有误
	loggerPath = "../conf/err_ars.logger.json"
	_, err = read()
	t.Log(err)
	if err == nil {
		t.Error("test fail")
		return
	}

	// 正常读取配置文件
	loggerPath = "../conf/ars.logger.json"
	_, err = read()
	t.Log(err)
	if err != nil {
		t.Errorf("读取配置文件失败：%v", err)
		return
	}
}

// TestWriteToFile 测试如果配置文件不存在，创建默认配置文件
func TestWriteToFile(t *testing.T) {
	loggerPath = "../conf/ars.logger.json"
	appenders, err := read()
	if err != nil {
		t.Errorf("读取配置文件失败：%v", err)
		return
	}

	// 读取配置文件失败【没有权限】
	err = writeToFile("/root/ars.logger.json", appenders)
	if err == nil {
		t.Error("test fail")
	}

	// 配置文件不存在
	err = writeToFile("../logs/newlogger/20161123.log", appenders)
	if err != nil {
		t.Errorf("test fail：%v", err)
	}

	// 正常写配置文件
	err = writeToFile("../logs/newlogger/20161125.log", appenders)
	if err != nil {
		t.Errorf("test fail：%v", err)
	}

	// 路径包含特殊字符
	err = writeToFile("../////", appenders)
	if err == nil {
		t.Error("test fail")
	}
}
