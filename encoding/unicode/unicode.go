package unicode

import (
	"bytes"
	"strconv"
)

// Encode 对字符串进行unicode编码
func Encode(str string) string {
	rs := []rune(str)
	json := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
			html += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

// Decode 对字符串进行unicode解码
func Decode(unicodeStr string) string {
	buf := bytes.NewBuffer(nil)
	i, j := 0, len(unicodeStr)
	for i < j {
		x := i + 6
		if x > j {
			buf.WriteString(unicodeStr[i:])
			break
		}
		if unicodeStr[i] == '\\' && unicodeStr[i+1] == 'u' {
			hex := unicodeStr[i+2 : x]
			r, err := strconv.ParseUint(hex, 16, 64)
			if err == nil {
				buf.WriteRune(rune(r))
			} else {
				buf.WriteString(unicodeStr[i:x])
			}
			i = x
		} else {
			buf.WriteByte(unicodeStr[i])
			i++
		}
	}
	return buf.String()
}
