package md5

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	input := "S3jOVETEcihe6U8xDZCgVBNoiRTUqTjh9ogg2VdeU/A="
	except := "831e434f277052c60b325d5b9d87a831"
	actual := Encrypt(input)
	if !strings.EqualFold(strings.ToLower(actual), strings.ToLower(except)) {
		t.Errorf("测试失败%s\t%s", strings.ToLower(actual), strings.ToLower(except))
	}
}
