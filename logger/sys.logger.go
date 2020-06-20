package logger

type sysLogger struct {
	appender IAppender
	layout   *Layout
}

func newSysLogger() *sysLogger {
	return &sysLogger{
		appender: NewStudoutAppender(),
	}
}

func (s *sysLogger) Log(event *LogEvent) {
	s.appender.Write(s.layout, event)
}
