package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/qxnw/lib4go/file"
)

type TestType struct {
	name string
	age  int
}

// TestDebug 测试记录Debug日志
func TestDebug(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")

	// 写入字符串
	log.Debug("content1")

	// 写入nil
	log.Debug(nil)

	// 写入int
	log.Debug(1)

	// 写入sliens
	log.Debug(make([]string, 2))

	// 写入数组
	log.Debug([3]int{1, 2, 3})

	// 写入结构体
	log.Debug(TestType{name: "test", age: 11})

	// 日志组件为空
	log = New("")
	log.Debug("hello world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("debug") != 6 {
		t.Errorf("test fail except : %d, actual : %d", 6, GetResult("debug"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestDebug 测试记录Debugf日志【format】
func TestDebugf(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")
	// 参数正确
	log.Debugf("%s %s", "hello", "world")

	// format 为空字符串
	log.Debugf("", "hello")

	// format 不包含格式化参数
	log.Debugf("hello", "world")

	// format 格式化参数过多
	log.Debugf("%s %s %s", "hello", "world")

	// 内容为nil
	log.Debugf("hello", nil)

	// 内容和格式化参数类型不匹配
	log.Debugf("%s %d", "hello", "world")

	// 内容为结构体
	log.Debugf("%+v", TestType{name: "test", age: 11})

	// 日志组件为空
	log = New("")
	log.Debugf("%s %s", "hello", "world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("debug") != 7 {
		t.Errorf("test fail except : %d, actual : %d", 7, GetResult("debug"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestInfo 测试记录Info日志
func TestInfo(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")

	// 写入字符串
	log.Info("content1")
	// // 每秒钟写入文件一次
	// time.Sleep(time.Second * 2)

	// 写入nil
	log.Info(nil)

	// 写入int
	log.Info(1)

	// 写入sliens
	log.Info(make([]string, 2))

	// 写入数组
	log.Info([3]int{1, 2, 3})

	// 写入结构体
	log.Info(TestType{name: "test", age: 11})

	// 日志组件为空
	log = New("")
	log.Info("hello world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("info") != 6 {
		t.Errorf("test fail except : %d, actual : %d", 6, GetResult("info"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestInfof 测试记录Info日志【format】
func TestInfof(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")
	// 参数正确
	log.Infof("%s %s", "hello", "world")

	// format 为空字符串
	log.Infof("", "hello")

	// format 不包含格式化参数
	log.Infof("hello", "world")

	// format 格式化参数过多
	log.Infof("%s %s %s", "hello", "world")

	// 内容为nil
	log.Infof("hello", nil)

	// 内容和格式化参数类型不匹配
	log.Infof("%s %d", "hello", "world")

	// 内容为结构体
	log.Infof("%+v", TestType{name: "test", age: 11})

	// 日志组件为空
	log = New("")
	log.Infof("%s %s", "hello", "world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("info") != 7 {
		t.Errorf("test fail except : %d, actual : %d", 7, GetResult("info"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestError 测试记录Error日志
func TestError(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")

	// 写入字符串
	log.Error("content1")

	// 写入nil
	log.Error(nil)

	// 写入int
	log.Error(1)

	// 写入sliens
	log.Error(make([]string, 2))

	// 写入数组
	log.Error([3]int{1, 2, 3})

	// 写入结构体
	log.Error(TestType{name: "test", age: 11})

	// 日志组件为空
	log = New("")
	log.Error("hello world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("error") != 6 {
		t.Errorf("test fail except : %d, actual : %d", 6, GetResult("error"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestErrorf 测试记录Error日志【format】
func TestErrorf(t *testing.T) {
	// 清空数据统计
	manager.factory = &testLoggerAppenderFactory{}
	ResultClear()

	log := New("key1")
	// 参数正确
	log.Errorf("%s %s", "hello", "world")

	// format 为空字符串
	log.Errorf("", "hello")

	// format 不包含格式化参数
	log.Errorf("hello", "world")

	// format 格式化参数过多
	log.Errorf("%s %s %s", "hello", "world")

	// 内容为nil
	log.Errorf("hello", nil)

	// 内容为结构体
	log.Errorf("%+v", TestType{name: "test", age: 11})

	// 内容和格式化参数类型不匹配
	log.Errorf("%s %d", "hello", "world")

	// 日志组件为空
	log = New("")
	log.Errorf("%s %s", "hello", "world")

	time.Sleep(time.Second * 2)
	// 统计数据是否和预期的一致
	if GetResult("error") != 7 {
		t.Errorf("test fail except : %d, actual : %d", 7, GetResult("error"))
	}

	Close()
	manager, _ = newLoggerManager()
}

// TestWriteToBuffer 测试写入日志的时候，是否漏掉了日志记录，通过测试的testLoggerAppenderFactory来不进行真的日志记录
func TestWriteToBuffer(t *testing.T) {
	manager.factory = &testLoggerAppenderFactory{}
	// 清空结果
	ResultClear()
	totalCount := 10000 * 1
	ch := make(chan int, totalCount)
	lk := sync.WaitGroup{}

	doWrite := func(ch chan int, lk *sync.WaitGroup) {
		log := New("abc")
	START:
		for {
			select {
			case v, ok := <-ch:
				if ok {
					log.Debug(v)
					log.Info(v)
					log.Error(v)
				} else {
					break START
				}
			}
		}
		lk.Done()
	}

	for i := 0; i < 100; i++ {
		lk.Add(1)
		go doWrite(ch, &lk)
	}

	for i := 0; i < totalCount; i++ {
		ch <- i
	}

	close(ch)
	lk.Wait()

	time.Sleep(time.Second * 2)

	Close()

	for i := 0; i < len(ACCOUNT); i++ {
		fmt.Println(ACCOUNT[i].name, " ", ACCOUNT[i].count)
		if strings.EqualFold(ACCOUNT[i].name, "debug") {
			if ACCOUNT[i].count != totalCount {
				t.Errorf("test fail, actual : %d", ACCOUNT[i].count)
			}
		}
		if strings.EqualFold(ACCOUNT[i].name, "info") {
			if ACCOUNT[i].count != totalCount {
				t.Errorf("test fail, actual : %d", ACCOUNT[i].count)
			}
		}
		// 测试不执行fatal日志记录
		if strings.EqualFold(ACCOUNT[i].name, "fatal") {
			if ACCOUNT[i].count != 0 {
				t.Errorf("test fail, actual : %d", ACCOUNT[i].count)
			}
		}
		if strings.EqualFold(ACCOUNT[i].name, "error") {
			if ACCOUNT[i].count != totalCount {
				t.Errorf("test fail, actual : %d", ACCOUNT[i].count)
			}
		}
	}

	manager, _ = newLoggerManager()
}

// TestLoggerToFile 测试输出到文件，并检验日志数量
func TestLoggerToFile(t *testing.T) {
	// 把数据写入文件
	totalAccount := 10000 * 1
	lk := sync.WaitGroup{}
	ch := make(chan int, totalAccount)
	name := "ABC"

	log := New(name)

	doWriteToFile := func(ch chan int, lk *sync.WaitGroup) {
	START:
		for {
			select {
			case l, ok := <-ch:
				if ok {
					log.Debug(l)
					log.Info(l)
					log.Error(l)
				} else {
					break START
				}
				lk.Done()
			}
		}
	}

	for i := 0; i < 100; i++ {
		go doWriteToFile(ch, &lk)
	}

	for i := 0; i < totalAccount; i++ {
		lk.Add(1)
		ch <- i
	}
	close(ch)
	lk.Wait()

	time.Sleep(time.Second * 1)

	Close()

	// 开始读取文件
	path := fmt.Sprintf("../logs/%s/%d%d%d.log", name, time.Now().Year(), time.Now().Month(), time.Now().Day())
	filePath, _ := file.GetAbs(path)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Errorf("test fail : %v", err)
	}
	count := len(strings.Split(string(data), "\n"))
	if count != totalAccount*3+6 {
		t.Errorf("test fail, actual:%d, except:%d", count, totalAccount*3+6)
	}

	// 删除日志防止下次进行测试的时候数据错误
	os.Remove(filePath)
}

// TestLoggerToFileCheckOrder 测试写入到文件，然后判断输入的顺序
func TestLoggerToFileCheckOrder(tx *testing.T) {
	// 写入日志到文件
	manager, _ = newLoggerManager()
	logger := GetSession("tofile", "12345678")
	// t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	// if err != nil {
	// 	tx.Errorf("test fail, %+v", err)
	// }

	// 构建要测试的数据和预期数据
	data := map[string]string{
		// LogEvent{Level: "Debug", Now: t, Name: "tofile", Session: "12345678", Content: "content1", Output: "output1"}: []string{"[d]", "content1"},
		// LogEvent{Level: "Debug", Now: t, Name: "tofile", Session: "12345678", Content: "content2", Output: "output2"}: []string{"[d]", "content2"},
		// LogEvent{Level: "Info", Now: t, Name: "tofile", Session: "12345678", Content: "content3", Output: "output3"}:  []string{"[i]", "content3"},
		// LogEvent{Level: "Fatal", Now: t, Name: "tofile", Session: "12345678", Content: "content4", Output: "output4"}: []string{"[f]", "content4"},
		// LogEvent{Level: "Error", Now: t, Name: "tofile", Session: "12345678", Content: "content5", Output: "output5"}: []string{"[e]", "content5"},
		// LogEvent{Level: "Error", Now: t, Name: "tofile", Session: "12345678", Content: "content6", Output: "output6"}: []string{"[e]", "content6"},
		// LogEvent{Level: "Test", Now: t, Name: "tofile", Session: "12345678", Content: "content7", Output: "output6"}:  []string{"[t]", "content7"},
		"content1": "content1",
		"content2": "content2",
		"content3": "content3",
		"content4": "content4",
		"content5": "content5",
		"content6": "content6",
		"content7": "content7",
	}

	// 获取日志文件的绝对路径
	filePath, _ := file.GetAbs(fmt.Sprintf("../logs/tofile/%d%d%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day()))

	// 删除文件，多次测试前面的测试会覆盖掉结果
	os.Remove(filePath)

	excepts := []string{}
	for event, except := range data {
		// 写内容到buffer
		// manager.Log(event)
		logger.Info(event)
		// 添加预期的结果【测试的时候map顺序不确定】
		// excepts = append(excepts, fmt.Sprintf(`[2016/11/28 16:38:27]%s[12345678] %s`, "12345678", except))
		excepts = append(excepts, fmt.Sprintf(`][i][12345678] %s`, except))
	}
	tx.Log(excepts)
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

// Account 日志记录的对象
type Account struct {
	name  string
	count int
}

// mutex 保证日志记录是原子操作
var mutex sync.Mutex

// ACCOUNT 记录日志的结果
var ACCOUNT []*Account

// SetResult 存放测试结果
func SetResult(name string, n int) {
	for i := 0; i < len(ACCOUNT); i++ {
		if strings.EqualFold(ACCOUNT[i].name, name) {
			mutex.Lock()
			ACCOUNT[i].count = ACCOUNT[i].count + n
			mutex.Unlock()
			return
		}
	}

	mutex.Lock()
	account := &Account{name: name, count: n}
	ACCOUNT = append(ACCOUNT, account)
	mutex.Unlock()
}

// GetResult 获取测试结果
func GetResult(name string) int {
	for i := 0; i < len(ACCOUNT); i++ {
		if strings.EqualFold(ACCOUNT[i].name, name) {
			return ACCOUNT[i].count
		}
	}
	return 0
}

// ResultClear 清空测试结果
func ResultClear() {
	ACCOUNT = []*Account{}
}
