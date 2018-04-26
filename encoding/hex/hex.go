package hex

import (
	"encoding/hex"
	"fmt"
)

// Encode 把[]byte类型通过hex编码成string
func Encode(src []byte) string {
	return hex.EncodeToString(src)
}

// Decode 把一个string类型通过hex解码成string
func Decode(src string) (r string, err error) {
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", fmt.Errorf("hex decode fail:%v", err)
	}
	r = string(data)
	return
}
