package logger

import (
	"time"

	"github.com/micro-plat/lib4go/concurrent/cmap"
)

//FileAppender 文件FileAppender
type FileAppender struct {
	writers cmap.ConcurrentMap
	ticker  *time.Ticker
	done    chan struct{}
}

//NewFileAppender 构建file FileAppender
func NewFileAppender() *FileAppender {
	a := &FileAppender{
		done:    make(chan struct{}),
		writers: cmap.New(6),
		ticker:  time.NewTicker(time.Minute * 10),
	}
	go a.clean()
	return a
}
func (a *FileAppender) Write(layout *Layout, event *LogEvent) error {
	p := event.Transform(layout.Path)
	_, w, err := a.writers.SetIfAbsentCb(p, func(input ...interface{}) (interface{}, error) {
		return newWriter(p, layout)
	})
	if err != nil {
		return err
	}
	w.(*writer).Write(event)
	return nil
}
func (a *FileAppender) clean() {
EXIT:
	for {
		select {
		case <-a.done:
			break EXIT
		case <-a.ticker.C:
		LOOP:
			for {
				select {
				case <-a.done:
					break LOOP
				case v, ok := <-a.writers.IterBuffered():
					if !ok {
						break LOOP
					}
					w := v.Val.(*writer)
					if time.Since(w.lastWrite) < 5*time.Minute {
						a.writers.Remove(v.Key)
					}
				}
			}
		}
	}

}

//Close 关闭组件
func (a *FileAppender) Close() error {
	close(a.done)
	a.writers.RemoveIterCb(func(key string, w interface{}) bool {
		w.(*writer).Close()
		return true
	})
	return nil
}
