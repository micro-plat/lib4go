package logger

import (
	"fmt"
	"strings"
)

type testLoggerAppenderFactory struct {
}

func (f *testLoggerAppenderFactory) MakeAppender(l *Appender, event *LogEvent) (IAppender, error) {
	switch strings.ToLower(l.Type) {
	case appender_file:
		appender, _ := NewFileAppender(f.MakeUniq(l, event), l)
		appender.file = &TestWriteCloser{}
		return appender, nil
	case appender_stdout:
		return NewStudoutAppender(f.MakeUniq(l, event), l)
	}
	return nil, fmt.Errorf("不支持的日志类型:%s", l.Type)
}

func (f *testLoggerAppenderFactory) MakeUniq(l *Appender, event *LogEvent) string {
	switch strings.ToLower(l.Type) {
	case appender_file:
		return transform(l.Path, event)
	}
	return l.Type
}

// 重写io.WriteCloser里面的方法，以方便测试
type TestWriteCloser struct {
}

func (f *TestWriteCloser) Write(p []byte) (n int, err error) {
	content := string(p)
	if strings.Contains(content, "[d]") {
		count := strings.Count(content, "[d]")
		SetResult("debug", count)
	}
	if strings.Contains(content, "[i]") {
		count := strings.Count(content, "[i]")
		SetResult("info", count)
	}
	if strings.Contains(content, "[f]") {
		count := strings.Count(content, "[f]")
		SetResult("fatal", count)
	}
	if strings.Contains(content, "[e]") {
		count := strings.Count(content, "[e]")
		SetResult("error", count)
	}

	return len(p), nil
}

func (f *TestWriteCloser) Close() error {
	return nil
}
