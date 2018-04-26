package url

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	urlEnInput := "www.baidu.com?name=tom"
	urlEnExcept := "www.baidu.com%3Fname%3Dtom"
	urlEnActual := Encode(urlEnInput)
	if !strings.EqualFold(urlEnExcept, urlEnActual) {
		t.Errorf("URLEncode fail %s to %s", urlEnInput, urlEnActual)
	}

	urlDeInput := "www.baidu.com%3Fname%3Dtom"
	urlDeExcept := "www.baidu.com?name=tom"
	urlDeActual, err := Decode(urlDeInput)
	if err != nil {
		t.Error("URLDecode fail")
		return
	}
	if !strings.EqualFold(urlDeExcept, urlDeActual) {
		t.Errorf("URLEncode fail %s to %s", urlDeInput, urlDeActual)
	}

	errDeInput := "!@#!!@"
	errDeActual, err := Decode(errDeInput)
	if err != nil {
		t.Error("URLDecode fail")
	}
	if !strings.EqualFold(errDeActual, errDeInput) {
		t.Error("URLDecode fail")
	}
}
