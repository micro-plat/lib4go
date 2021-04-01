package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestTransform(t *testing.T) {
	tmpl := `{"dt":"%datetime.%ms","level":"%l","session":"%session","content":"%content"}`
	cases := []struct {
		event  *LogEvent
		isJSON bool
	}{
		{event: NewLogEvent("a", "info", "10101010", `\t\n\b`, nil, 0), isJSON: true},
		// {event: NewLogEvent("a", "info", "10101010", `[{\"server-ip\":\"192.168.4.121\",\"time\":\"2021/04/01 10:21:47.958835\",\"level\":\"info\",\"session\":\"d01859e83\",\"content\":\"\"初始化: /qxgrs_down_debug/bestpay-flow/mqc-cron/flowserver/conf\"\"},{\"server-ip\":\"192.168.4.121\",\"time\":\"2021/04/01 10:21:47.996666\",\"level\":\"info\",\"session\":\"d01859e83\",\"content\":\"\"启动[mqc]服务...\"\"},{\"server-ip\":\"192.168.4.121\",\"time\":\"2021/04/01 10:21:48.370566\",\"level\":\"info\",\"session\":\"ff02bc5c5\",\"content\":\"\"启动成功(mqc,mqc://192.168.4.121,[2])\"\"},{\"server-ip\":\"192.168.4.121\",\"time\":\"2021/04/01 10:21:48.434205\",\"level\":\"info\",\"session\":\"d01859e83\",\"content\":\"\"启动[cron]服务...\"\"},{\"server-ip\":\"192.168.4.121\",\"time\":\"2021/04/01 10:21:48.656895\",\"level\":\"info\",\"session\":\"741c890f0\",\"content\":\"\"启动成功(cron,cron://192.168.4.121,[3])\"\"}]`, nil, 0), isJSON: true},
	}
	for _, c := range cases {
		r := c.event.Transform(tmpl)
		// assert.Equal(t, "", r)
		assert.Equal(t, c.isJSON, json.Valid([]byte(r)), r)

		var buff bytes.Buffer
		err := json.Compact(&buff, []byte("["+r+"]"))
		assert.Equal(t, nil, err, r)

	}
}
