package encoding

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetReader(t *testing.T) {
	input := "你好"
	charset := "utf-8"
	data, err := ioutil.ReadAll(GetReader(input, charset))
	if err != nil {
		t.Errorf("测试失败：%v", err)
	}
	if !strings.EqualFold(string(data), input) {
		t.Errorf("GetReader fail %s to %s", input, string(data))
	}

	charset = ""
	data, err = ioutil.ReadAll(GetReader(input, charset))
	if err != nil {
		t.Errorf("测试失败：%v", err)
	}
	if !bytes.EqualFold(data, []byte(input)) {
		t.Error("GetReader fail")
	}

	input = "你好"
	charset = "gbk"
	data, err = ioutil.ReadAll(GetReader(input, charset))
	if err != nil {
		t.Errorf("测试失败：%v", err)
	}
	if strings.EqualFold(string(data), input) {
		t.Errorf("GetReader fail %s to %s", input, string(data))
	}

	input = "你好"
	charset = "gb2312"
	data, err = ioutil.ReadAll(GetReader(input, charset))
	if err != nil {
		t.Errorf("测试失败：%v", err)
	}
	if strings.EqualFold(string(data), input) {
		t.Errorf("GetReader fail %s to %s", input, string(data))
	}
}

func TestCovert(t *testing.T) {
	input := "你好"
	charset := "utf-8"
	except := "你好"
	actual, err := Convert([]byte(input), charset)
	if err != nil {
		t.Error("Covert fail")
	}
	if !strings.EqualFold(actual, except) {
		t.Errorf("GetReader fail %s to %s", input, actual)
	}

	charset = ""
	_, err = ioutil.ReadAll(GetReader(input, charset))
	if err != nil {
		t.Errorf("测试失败:%v", err)
	}
	actual, err = Convert([]byte(input), charset)
	if err == nil {
		t.Error("Covert fail")
	}

	input = "你好"
	charset = "gbk"
	except = "浣犲ソ"
	actual, err = Convert([]byte(input), charset)
	if err != nil {
		t.Error("Covert fail")
	}
	if !strings.EqualFold(actual, except) {
		t.Errorf("GetReader fail %s to %s", input, actual)
	}

	input = "你好"
	charset = "gb2312"
	except = "浣犲ソ"
	actual, err = Convert([]byte(input), charset)
	if err != nil {
		t.Error("Covert fail")
	}
	if !strings.EqualFold(actual, except) {
		t.Errorf("GetReader fail %s to %s", input, actual)
	}
}
