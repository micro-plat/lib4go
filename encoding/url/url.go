package url

import (
	"net/url"

	"github.com/qxnw/lib4go/encoding"
)

// Encode 对字符串进行url编码
func Encode(input string) string {
	return url.QueryEscape(input)
}

// Decode 对字符串进行url解码
func Decode(input string) (string, error) {
	return url.QueryUnescape(input)
}

//DecodeGBK gbk解码
func DecodeGBK(input string) (string, error) {
	gbkBuff, err := Decode(input)
	if err != nil {
		return "", err
	}
	c, err := encoding.Convert([]byte(gbkBuff), "gbk")
	if err != nil {
		return "", err
	}
	return c, nil
}
