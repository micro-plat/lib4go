package logger

import (
	"bytes"
	"io"
	"sync"
	"time"

	"github.com/micro-plat/lib4go/file"
)

//writer 文件输出器
type writer struct {
	name      string
	buffer    *bytes.Buffer
	lastWrite time.Time
	layout    *Layout
	interval  time.Duration
	file      io.WriteCloser
	ticker    *time.Ticker
	locker    sync.Mutex
	Level     int
}

//newwriter 构建基于文件流的日志输出对象
func newWriter(path string, layout *Layout) (fa *writer, err error) {
	fa = &writer{layout: layout, interval: time.Second}
	fa.Level = GetLevel(layout.Level)
	fa.buffer = bytes.NewBufferString("\n--------------------begin------------------------\n\n")
	fa.ticker = time.NewTicker(fa.interval)
	fa.file, err = file.CreateFile(path)
	if err != nil {
		return
	}
	go fa.writeTo()
	return
}

//Write 写入日志
func (f *writer) Write(event *LogEvent) {
	if f.Level > GetLevel(event.Level) {
		return
	}
	f.locker.Lock()
	f.buffer.WriteString(event.Output)
	f.locker.Unlock()
	f.lastWrite = time.Now()
}

//Close 关闭当前appender
func (f *writer) Close() {
	f.Level = ILevel_OFF
	f.ticker.Stop()
	f.locker.Lock()
	f.buffer.WriteString("\n---------------------end-------------------------\n")
	f.buffer.WriteTo(f.file)
	f.file.Close()
	f.locker.Unlock()
}

//writeTo 定时写入文件
func (f *writer) writeTo() {
START:
	for {
		select {
		case _, ok := <-f.ticker.C:
			if ok {
				f.locker.Lock()
				f.buffer.WriteTo(f.file)
				f.buffer.Reset()
				f.locker.Unlock()
			} else {
				break START
			}
		}
	}
}
