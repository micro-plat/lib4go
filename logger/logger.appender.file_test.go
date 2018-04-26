package logger

import (
	"testing"
	"time"

	"os"

	"bufio"

	"io"

	"strings"

	"github.com/qxnw/lib4go/file"
)

// TestNewFailAppender 测试构建一个日志输出对象appender
func TestNewFailAppender(t *testing.T) {
	// 创建日志输出对象，创建文件成功
	path := "../logs/test.log"
	layout := &Appender{Type: "file", Level: "All"}
	_, err := NewFileAppender(path, layout)
	if err != nil {
		t.Errorf("test fail:%v", err)
	}

	// 创建日志输出对象，创建文件失败
	path = "/root/test.log"
	layout = &Appender{Type: "file", Level: "All"}
	_, err = NewFileAppender(path, layout)
	if err == nil {
		t.Error("test fail")
	}
}

// TestWrite 测试写日志到buffer
func TestWrite(t *testing.T) {
	path := "../logs/test.log"
	layout := &Appender{Type: "file", Level: "All"}
	f, err := NewFileAppender(path, layout)
	if err != nil {
		t.Errorf("test fail:%v", err)
	}

	event := &LogEvent{Level: "All", Output: "output"}
	f.Write(event)

	// 不能写日志
	f.Level = getLevel("Off")
	f.Write(event)
}

// TestWriteToFileAndReadCheck 测试写入buffer的内容和最后在文件中读取到的内容是否一致
func TestWriteToFileAndReadCheck(tx *testing.T) {
	// 写入文件中
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}
	path := "../logs/20161128.log"
	// layout := &Appender{Type: "file", Level: "All", Path: path}
	layout := &Appender{Type: "file", Level: "All"}
	fa, err := NewFileAppender(path, layout)
	if err != nil {
		tx.Errorf("test fail:%v", err)
	}

	events := []*LogEvent{
		&LogEvent{Level: "Debug", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output1"},
		&LogEvent{Level: "Debug", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output2"},
		&LogEvent{Level: "Info", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output3"},
		&LogEvent{Level: "Fatal", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output4"},
		&LogEvent{Level: "Error", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output5"},
		&LogEvent{Level: "Error", Now: t, Name: "test", Session: "12345678", Content: "content", Output: "output6"},
	}
	for _, event := range events {
		fa.Write(event)
	}
	fa.Close()

	// 读取文件，进行校验
	filePath, _ := file.GetAbs(path)
	f, err := os.Open(filePath)
	if err != nil {
		tx.Errorf("test fail:%v", err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			tx.Errorf("test fail: %v", err)
			break
		}
		if strings.Contains(line, "output1output2output3output4output5output6") {
			tx.Log("test success!")
			break
		}
	}

	// 删除文件
	f.Close()
	err = os.Remove(filePath)
	if err != nil {
		tx.Errorf("test fail:%v", err)
	}
}
