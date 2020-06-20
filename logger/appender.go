package logger

import (
	"fmt"
	"sync"
)

//IAppender 定义appender接口
type IAppender interface {
	Write(*Layout, *LogEvent) error
	Close() error
}

type appenderWriter struct {
	appenders map[string]IAppender
	layouts   []*Layout
	lock      sync.RWMutex
}

func newAppenderWriter() *appenderWriter {
	return &appenderWriter{
		appenders: make(map[string]IAppender),
		layouts:   make([]*Layout, 0, 2),
	}
}

func (a *appenderWriter) AddAppender(typ string, i IAppender) {
	a.lock.Lock()
	defer a.lock.Lock()
	if _, ok := a.appenders[typ]; ok {
		panic(fmt.Errorf("不能重复注册appender:%s", typ))
	}
	a.appenders[typ] = i
}
func (a *appenderWriter) AddLayout(layouts ...*Layout) {
	a.lock.Lock()
	defer a.lock.Lock()
	for _, layout := range layouts {
		if _, ok := a.appenders[layout.Type]; !ok {
			panic(fmt.Errorf("layout中配置的日志组件类型不支持:%s", layout.Type))
		}
		a.layouts = append(a.layouts, layout)
	}
}
func (a *appenderWriter) Log(event *LogEvent) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, layout := range a.layouts {
		if GetLevel(layout.Level) > GetLevel(event.Level) {
			continue
		}
		a.appenders[layout.Type].Write(layout, event.Event(layout.Layout))
	}
}

var defWriter = newAppenderWriter()
var sysLog = newSysLogger()

//AddAppender 添加appender
func AddAppender(typ string, i IAppender) {
	defWriter.AddAppender(typ, i)
}

//AddLayout 添加日志输出配置
func AddLayout(l ...*Layout) {
	defWriter.AddLayout(l...)
}

func logNow(event *LogEvent) {
	defWriter.Log(event)
}
