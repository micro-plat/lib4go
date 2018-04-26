package hex

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	enInput := "www.baidu.com"
	enExcept := "7777772e62616964752e636f6d"
	enActual := Encode([]byte(enInput))
	if !strings.EqualFold(enExcept, enActual) {
		t.Errorf("Encode fail %s to %s", enExcept, enActual)
	}

	deInput := "7777772e62616964752e636f6d"
	deExcept := "www.baidu.com"
	deActual, err := Decode(deInput)
	if err != nil {
		t.Error("Decode错误")
	}
	if !strings.EqualFold(deExcept, deActual) {
		t.Errorf("Decode fail %s to %s", deExcept, deActual)
	}

	errInput := "!@#!$"
	_, err = Decode(errInput)
	if err == nil {
		t.Error("测试失败")
	}
}
