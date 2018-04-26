package logger

import (
	"testing"
)

// TestSysLoogerInfo 测试记录系统的Info日志
func TestSysLoogerInfo(t *testing.T) {
	sysLoggerInfo("content")
	sysLoggerInfo(nil)
	sysLoggerInfo([]string{"1", "2"})
	sysLoggerInfo([2]string{"1", "2"})
	sysLoggerInfo(TestType{name: "name", age: 12})
	sysLoggerInfo("content")
}

// TestLoggerError 测试记录系统的Error日志
func TestLoggerError(t *testing.T) {
	sysLoggerError("content")
	sysLoggerError(nil)
	sysLoggerError([]string{"1", "2"})
	sysLoggerError([2]string{"1", "2"})
	sysLoggerError(TestType{name: "name", age: 12})
	sysLoggerError("content")
}

// TestSysLoggerWrite 测试记录系统日志
func TestSysLoggerWrite(t *testing.T) {
	sysLoggerWrite("info", "content")
	sysLoggerWrite("test", nil)
	sysLoggerWrite("info", []string{"1", "2"})
	sysLoggerWrite("info", [2]string{"1", "2"})
	sysLoggerWrite("info", TestType{name: "name", age: 12})
	sysLoggerWrite("", "content")
}
