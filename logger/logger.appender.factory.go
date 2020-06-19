package logger

import (
	"fmt"
	"strings"
)

//MakeAppender 根据appender配置及日志信息生成appender对象
func MakeAppender(l *Layout, event *LogEvent) (IAppender, error) {
	switch strings.ToLower(l.Type) {
	case appender_file:
		return NewFileAppender(MakeUniq(l, event), l)
	case appender_stdout:
		return NewStudoutAppender(MakeUniq(l, event), l)
	default:
		if value, b := registedFactory.Get(l.Type); b {
			fa := value.(IAppenderCreator)
			return fa.MakeAppender(l, event)
		}
	}
	return nil, fmt.Errorf("不支持的日志类型:%s", l.Type)
}

//MakeUniq 根据appender配置及日志信息生成appender唯一编号
func MakeUniq(l *Layout, event *LogEvent) string {
	switch strings.ToLower(l.Type) {
	case appender_file:
		return transform(l.Path, event)
	case appender_stdout:
		return l.Type
	default:
		if value, b := registedFactory.Get(l.Type); b {
			fa := value.(IAppenderCreator)
			return fa.MakeUniq(l, event)
		}
	}
	return l.Type
}
