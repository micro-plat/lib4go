package logger

import (
	"strings"
	"testing"
	"time"
)

// TestTransform 测试日志格式转换成具体日志的模块
func TestTransform(tx *testing.T) {
	t, err := time.Parse("2006/01/02 15:04:05", "2016/11/28 16:38:27")
	if err != nil {
		tx.Errorf("test fail, %+v", err)
	}
	tpls := map[string][]interface{}{
		``:          []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, ``},
		`%session`:  []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `12345678`},
		`%date`:     []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `20161128`},
		`%datetime`: []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `2016/11/28 16:38:27`},
		`%year`:     []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `2016`},
		`%mm`:       []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `11`},
		`%dd`:       []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `28`},
		`%hh`:       []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `16`},
		`%mi`:       []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `38`},
		`%ss`:       []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `27`},
		`%level`:    []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `Info`},
		`%l`:        []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `I`},
		`%name`:     []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `name1`},
		`%content`:  []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `content1`},
		`%test`:     []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, ``},
		`test`:      []interface{}{&LogEvent{Level: "Info", Now: t, Name: "name1", Session: "12345678", Content: "content1", Output: "output1"}, `test`},
	}

	for tpl, except := range tpls {
		actual := transform(tpl, except[0].(*LogEvent))
		if !strings.EqualFold(actual, except[1].(string)) {
			tx.Errorf("test fail actual：%s\texcept:%s\t", actual, except[1].(string))
		}
	}
}
