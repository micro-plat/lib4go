package base64

import (
	"bytes"
	"strings"
	"testing"
)

func TestUrlToByte(t *testing.T) {
	input := []byte("http://www.baidu.com")
	except := "aHR0cDovL3d3dy5iYWlkdS5jb20="
	actual := URLEncodeBytes(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}

	actualByte, err := URLDecodeBytes(actual)
	if err != nil {
		t.Errorf("test fail err : %v", err)
	}
	if !bytes.EqualFold(actualByte, input) {
		t.Error("test fail")
	}

	errInput := "!@#!"
	_, err = URLDecodeBytes(errInput)
	if err == nil {
		t.Error("测试错误")
		return
	}
}

func TestUrlToString(t *testing.T) {
	input := "http://www.baidu.com"
	except := "aHR0cDovL3d3dy5iYWlkdS5jb20="
	actual := URLEncode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}

	actual, err := URLDecode(actual)
	if err != nil {
		t.Errorf("test fail err : %v", err)
	}
	if !strings.EqualFold(actual, input) {
		t.Error("test fail")
	}

	errInput := "!@#!"
	_, err = URLDecode(errInput)
	if err == nil {
		t.Error("测试错误")
		return
	}
}
