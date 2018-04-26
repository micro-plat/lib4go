package xsrf

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

//CreateXSRFToken 生成xsrf token
func CreateXSRFToken(secret string, data string) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(data))
	encoder.Close()
	vb := buf.Bytes()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sig := getCookieSig(secret, vb, timestamp)
	return strings.Join([]string{string(vb), timestamp, sig}, "|")
}
func getCookieSig(key string, val []byte, timestamp string) string {
	hm := hmac.New(sha1.New, []byte(key))

	hm.Write(val)
	hm.Write([]byte(timestamp))

	hex := fmt.Sprintf("%02x", hm.Sum(nil))
	return hex
}

//ParseXSRFToken 转换xsrf token
func ParseXSRFToken(secret string, value string) string {
	parts := strings.SplitN(value, "|", 3)
	val, timestamp, sig := parts[0], parts[1], parts[2]
	if getCookieSig(secret, []byte(val), timestamp) != sig {
		return ""
	}

	ts, _ := strconv.ParseInt(timestamp, 0, 64)
	if time.Now().Unix()-31*86400 > ts {
		return ""
	}

	buf := bytes.NewBufferString(val)
	encoder := base64.NewDecoder(base64.StdEncoding, buf)

	res, _ := ioutil.ReadAll(encoder)
	return string(res)
}
