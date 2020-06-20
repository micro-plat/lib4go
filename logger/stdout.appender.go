package logger

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"sync"

	"github.com/zkfy/log"
)

//StdoutAppender 标准输出器
type StdoutAppender struct {
	name      string
	lastWrite time.Time
	output    *log.Logger
	buffer    *bytes.Buffer
	ticker    *time.Ticker
	interval  time.Duration
	lock      sync.Mutex
}

//NewStudoutAppender 构建基于文件流的日志输出对象
func NewStudoutAppender() (fa *StdoutAppender) {
	fa = &StdoutAppender{interval: time.Millisecond * 200}
	fa.buffer = bytes.NewBufferString("")
	fa.output = log.New(fa.buffer, "", log.Llongcolor)
	fa.ticker = time.NewTicker(fa.interval)
	fa.output.SetOutputLevel(log.Ldebug)
	go fa.writeTo()
	return
}

//Write 写入日志
func (f *StdoutAppender) Write(layout *Layout, event *LogEvent) error {
	current := GetLevel(event.Level)
	if GetLevel(layout.Level) > GetLevel(event.Level) {
		return nil
	}
	f.lastWrite = time.Now()
	f.lock.Lock()
	defer f.lock.Unlock()
	switch current {
	case ILevel_Debug:
		f.output.Debug(event.Output)
	case ILevel_Info:
		f.output.Info(event.Output)
	case ILevel_Warn:
		f.output.Warn(event.Output)
	case ILevel_Error:
		f.output.Error(event.Output)
	case ILevel_Fatal:
		f.output.Output("", log.Lfatal, 1, fmt.Sprintln(event.Output))
	}
	return nil
}

//writeTo 定时写入文件
func (f *StdoutAppender) writeTo() {
START:
	for {
		select {
		case _, ok := <-f.ticker.C:
			if ok {
				f.lock.Lock()
				f.buffer.WriteTo(os.Stdout)
				f.buffer.Reset()
				f.lock.Unlock()
			} else {
				f.buffer.WriteTo(os.Stdout)
				break START
			}
		}
	}
}

//Close 关闭当前appender
func (f *StdoutAppender) Close() error {
	f.ticker.Stop()
	return nil
}
