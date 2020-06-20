package logger

import "fmt"

type sysLogger struct {
	appender IAppender
	layout   *Layout
}

func newSysLogger() *sysLogger {
	return &sysLogger{
		appender: NewStudoutAppender(),
		layout:   &Layout{Layout: "[%datetime.%ms][%l][%session]%content", Level: SLevel_ALL},
	}
}
func (s *sysLogger) Errorf(f string, content ...interface{}) {
	s.Log(NewLogEvent("sys", SLevel_Error, CreateSession(), fmt.Sprintf(f, content...), nil, 0))
}
func (s *sysLogger) Log(event *LogEvent) {
	s.appender.Write(s.layout, event)
}
