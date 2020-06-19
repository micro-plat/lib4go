package logger

import "github.com/micro-plat/lib4go/concurrent/cmap"

type IAppenderCreator interface {
	GetType() string
	MakeAppender(l *Layout, event *LogEvent) (IAppender, error)
	MakeUniq(l *Layout, event *LogEvent) string
}

var registedFactory cmap.ConcurrentMap

func init() {
	registedFactory = cmap.New(2)
}
func RegistryFactory(f IAppenderCreator, layout *Layout) {
	registedFactory.SetIfAbsent(f.GetType(), f)
	manager.append(layout)
}
func UnRegistryFactory(tp string) {
	registedFactory.Remove(tp)
	manager.remote(tp)
}
