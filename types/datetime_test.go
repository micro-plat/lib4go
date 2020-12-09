package types

import (
	"encoding/json"
	"encoding/xml"

	"testing"
	"time"

	"github.com/micro-plat/lib4go/assert"
)

type Time struct {
	NowTime *DateTime `json:"nowTime" xml:"nowTime"`
}

func TestDatetime(t *testing.T) {
	nowtime := time.Now()
	timefmt := "2006-01-02 15:04:05"
	expect := nowtime.Format(timefmt)
	var input = Time{
		NowTime: NewDateTime(nowtime),
	}

	resultJson := &Time{}

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		assert.Errorf(t, err, "json.Marshal出错", err)
	}
	err = json.Unmarshal(jsonBytes, resultJson)
	if err != nil {
		assert.Errorf(t, err, "json.Unmarshal出错", err)
	}
	t.Log("jsonBytes:", string(jsonBytes))

	resultXml := &Time{}

	xmlBytes, err := xml.Marshal(input)
	if err != nil {
		assert.Errorf(t, err, "json.Marshal出错", err)
	}
	err = xml.Unmarshal(xmlBytes, resultXml)
	if err != nil {
		assert.Errorf(t, err, "json.Unmarshal出错", err)
	}

	t.Log("xmlBytes:", string(xmlBytes))
	t.Log("resultJson:", resultJson.NowTime.String())
	t.Log("resultXml:", resultXml.NowTime.String())

	assert.Equal(t, expect, resultJson.NowTime.String(), "json")
	assert.Equal(t, expect, resultXml.NowTime.String(), "xml")

}
