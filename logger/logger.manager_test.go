package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/qxnw/lib4go/file"
)

// TestLog 测试manager的Log方法可能遇到的情况
func TestLog(tx *testing.T) {
	manager.isClose = true
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}
	manager.Log(&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"})

	// 测试完成打开appender，否则影响其他测试
	manager.isClose = false

	// 写入一个类型不存在的日志，进入记录系统日志的方法
	testCallBack = func(err error) {
		tx.Logf("进入到回调函数:%v", err)
		if !strings.EqualFold("不支持的日志类型:test", err.Error()) {
			tx.Errorf("test fail:%v", err)
		}
	}
	manager.Log(&LogEvent{Level: "test", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"})
	testCallBack = nil
}

// testclearUp 单独测试clearUp中的关键代码
func (a *loggerManager) testclearUp() {
	count := a.appenders.RemoveIterCb(func(key string, v interface{}) bool {
		apd := v.(*appenderEntity)
		if time.Now().Sub(apd.last).Seconds() > 5 {
			apd.appender.Close()
			return true
		}
		return false
	})
	if count > 0 {
		sysLoggerInfo("已移除:", count)
	}
}

// TestClearUp 测试testclearUp代码，就是manager中的clearUp， 只是时间减少了
func TestClearUp(tx *testing.T) {
	// 保证至少有一个appender
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}
	manager.Log(&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"})
	count := len(manager.appenders.Keys())

	// 休眠6秒
	time.Sleep(time.Second * 6)

	// 调用方法，判断是否被清理
	manager.testclearUp()
	if len(manager.appenders.Keys()) != 0 {
		tx.Errorf("test fail before count:%d, now:%d", count, len(manager.appenders.Keys()))
	}
}

// TestManagerClose 测试manager的Close方法
func TestManagerClose(tx *testing.T) {
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}
	manager.Log(&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"})
	if len(manager.appenders.Keys()) == 0 {
		tx.Error("test fail")
	}

	manager.Close()
	if len(manager.appenders.Keys()) != 0 {
		tx.Errorf("test fail:manager appenders have:%d", len(manager.appenders.Keys()))
	}
}

// TestLogToFile 测试通过Log方法写入到文件之后的顺序是否和预期的一致，使用map，可以使顺序不一致，最后读取文件进行校验
func TestLogToFile(tx *testing.T) {
	// 写入日志到文件
	manager, _ = newLoggerManager()
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}

	// 构建要测试的数据和预期数据
	data := map[*LogEvent][]string{
		&LogEvent{Level: "Debug", Now: t, Name: "tofile", Session: "12345678", Content: "content1", Output: "output1"}: []string{"[d]", "content1"},
		&LogEvent{Level: "Debug", Now: t, Name: "tofile", Session: "12345678", Content: "content2", Output: "output2"}: []string{"[d]", "content2"},
		&LogEvent{Level: "Info", Now: t, Name: "tofile", Session: "12345678", Content: "content3", Output: "output3"}:  []string{"[i]", "content3"},
		&LogEvent{Level: "Fatal", Now: t, Name: "tofile", Session: "12345678", Content: "content4", Output: "output4"}: []string{"[f]", "content4"},
		&LogEvent{Level: "Error", Now: t, Name: "tofile", Session: "12345678", Content: "content5", Output: "output5"}: []string{"[e]", "content5"},
		&LogEvent{Level: "Error", Now: t, Name: "tofile", Session: "12345678", Content: "content6", Output: "output6"}: []string{"[e]", "content6"},
		&LogEvent{Level: "Test", Now: t, Name: "tofile", Session: "12345678", Content: "content7", Output: "output6"}:  []string{"[t]", "content7"},
	}

	// 获取日志文件的绝对路径
	filePath, _ := file.GetAbs("../logs/tofile/20161128.log")

	// 删除文件，多次测试前面的测试会覆盖掉结果
	os.Remove(filePath)

	excepts := []string{}
	for event, except := range data {
		// 写内容到buffer
		manager.Log(event)
		// 添加预期的结果【测试的时候map顺序不确定】
		excepts = append(excepts, fmt.Sprintf(`[2016/11/28 16:38:27]%s[12345678] %s`, except[0], except[1]))
	}
	time.Sleep(time.Second * 11)

	// 读取文件中的类容

	// 当前读取文件的行数
	lineNow := 0

	// 记录匹配的行数
	actual := []int{}

	// 循环预期结果
	for _, except := range excepts {
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			tx.Errorf("test fail : %v", err)
		}
		for line, lineData := range strings.Split(string(fileData), "\n") {
			// 记录开始位置
			if strings.Contains(lineData, "begin") && lineNow == 0 {
				lineNow = line
			}

			// 有开始位置，开始匹配
			if line > lineNow {
				if strings.Contains(lineData, except) {
					lineNow = line
					actual = append(actual, line)
				}
			}
		}
	}

	if len(actual) != len(excepts) {
		tx.Errorf("test fail except: %d, actual: %d", len(excepts), len(actual))
	}

	tx.Log(actual)

	// 判断预期结果是否是连续的
	for i := 0; i < len(actual); i++ {
		if i != 0 {
			if actual[i-1]+1 != actual[i] {
				tx.Errorf("test fail, %+v", actual)
				return
			}
		}
	}
}
