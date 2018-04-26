package md5

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	input := "md5str"
	except := "7A3F2CFF1CD6BA4B6A23B7BFAD65719D"
	actual := Encrypt(input)
	if !strings.EqualFold(strings.ToLower(actual), strings.ToLower(except)) {
		t.Errorf("测试失败%s\t%s", strings.ToLower(actual), strings.ToLower(except))
	}
}
