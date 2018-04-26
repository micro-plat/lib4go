package weixin

import "encoding/xml"

type Wechat struct {
	w WechatEntity
}

//NewWechat 创建微信加解密对象
func NewWechat(appid string, token string, encodingAESKey string) (w Wechat, err error) {
	w = Wechat{}
	w.w, err = NewWechatEntity(appid, token, encodingAESKey)
	return
}

//Decrypt 解密请求报文
func (w Wechat) Decrypt(content string) (r string, err error) {
	response, err := w.w.Decrypt(content)
	if err != nil {
		return
	}
	buff, err := xml.Marshal(response)
	if err != nil {
		return
	}
	r = string(buff)
	return
}

//Encrypt 加密响应报文
func (w Wechat) Encrypt(fromUserName, toUserName, content, nonce, timestamp string) (r string, err error) {
	buffer, err := w.w.makeEncryptResponseBody(fromUserName, toUserName, content, nonce, timestamp)
	if err != nil {
		return
	}
	r = string(buffer)
	return
}

//MakeSign 构建签名
func (w Wechat) MakeSign(timestamp, nonce, msgEncrypt string) string {
	return w.w.makeMsgSignature(timestamp, nonce, msgEncrypt)
}
