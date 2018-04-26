package base64

import (
	"bytes"
	"strings"
	"testing"
)

func TestStdToByte(t *testing.T) {
	input := []byte("你好")
	except := "5L2g5aW9"
	actual := EncodeBytes(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}

	actualByte, err := DecodeBytes(actual)
	if err != nil {
		t.Errorf("test fail err : %v", err)
	}
	if !bytes.EqualFold(actualByte, input) {
		t.Error("test fail")
	}

	errInput := "!@#!"
	_, err = DecodeBytes(errInput)
	if err == nil {
		t.Error("测试错误")
		return
	}
}

func TestStdToString(t *testing.T) {
	input := "你好"
	except := "5L2g5aW9"
	actual := Encode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}

	actual, err := Decode(actual)
	if err != nil {
		t.Errorf("test fail err : %v", err)
	}
	if !strings.EqualFold(actual, input) {
		t.Error("test fail")
	}

	errInput := "!@#!"
	_, err = Decode(errInput)
	if err == nil {
		t.Error("测试错误")
		return
	}
}
