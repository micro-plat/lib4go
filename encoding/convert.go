package encoding

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//UTF82GBK utf8字符串转gbk编码
func UTF82GBK(content string) (result []byte, err error) {
	reader := GetEncodeReader([]byte(content), "gbk")
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		err = fmt.Errorf("编码转换失败:content:%s, err:%+v", content, err)
		return
	}
	return d, nil
}

//GBK2UTF8 将gbk编码转换为utf-8编码
func GBK2UTF8(content string) (result []byte, err error) {
	reader := GetDecodeReader([]byte(content), "gbk")
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		err = fmt.Errorf("编码转换失败:content:%s, err:%+v", content, err)
		return
	}
	return d, nil
}

// GetDecodeReader 获取
// charset不区分大小写
func GetDecodeReader(buff []byte, charset string) io.Reader {
	charset = strings.ToLower(charset)
	if strings.EqualFold(charset, "gbk") || strings.EqualFold(charset, "gb2312") {
		return transform.NewReader(bytes.NewReader(buff), simplifiedchinese.GBK.NewDecoder())
	}
	return strings.NewReader(string(buff))
}

// GetEncodeReader 获取
func GetEncodeReader(buff []byte, charset string) io.Reader {
	charset = strings.ToLower(charset)
	if strings.EqualFold(charset, "gbk") || strings.EqualFold(charset, "gb2312") {
		return transform.NewReader(bytes.NewReader(buff), simplifiedchinese.GBK.NewEncoder())
	}
	return strings.NewReader(string(buff))
}
