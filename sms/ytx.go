package sms

/*配置文件内容：{
    "appid": "8a48b5514e5298b9014e67a3f02f1411",
    "main_account": "aaf98fda42c744c90142d505bfab0135",
    "main_account_token": "5cd9cb617def42678553fe4d93b8f291",
    "soft_version": "2013-12-26",
    "sign": "{@main_account}{@main_account_token}{@timestamp}",
    "auth":"@main_account:{@timestamp}",
    "header":"Accept:application/xml\r\nContent-type:application/xml;charset=utf-8\r\nAuthorization:@auth",
    "url": "https://app.cloopen.com:8883/{@soft_version}/Accounts/{@main_account}/SMS/TemplateSMS?sig=@sign",
    "body": "<?xml version='1.0' encoding='utf-8'?><TemplateSMS><to>{@mobile}</to><appId>{@appid}</appId><templateId>65871</templateId><datas>{@data}</datas></TemplateSMS>"
}*/

import (
	"fmt"

	"strings"

	"time"

	"github.com/qxnw/lib4go/encoding/base64"
	"github.com/qxnw/lib4go/net/http"
	"github.com/qxnw/lib4go/security/md5"
)

type eSMS struct {
	mobile  string
	data    string
	url     string
	body    string
	charset string
	header  map[string]string
}

func getYtxParams(mobile, data, content string) (sms *eSMS, err error) {
	sms = &eSMS{header: make(map[string]string)}
	datas := strings.Split(data, ";")
	for _, v := range datas {
		sms.data = fmt.Sprintf("%s<data>%s</data>", sms.data, v)
	}

	form, err := NewJSONConfWithJson(content, 0, nil)
	if err != nil {
		err = fmt.Errorf("setting[%s]配置错误，无法解析(err:%v)", content, err)
		return
	}
	form.Transform.Set("mobile", mobile)
	form.Transform.Set("data", sms.data)
	form.Transform.Set("timestamp", time.Now().Format("20060102150405"))

	raw := form.Translate(form.String("raw"))
	if raw == "" || strings.Contains(raw, "@") {
		err = fmt.Errorf("args.setting配置错误，raw配置错误(raw:%s)(%s)", raw, content)
		return
	}
	form.Transform.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	sms.url = form.Translate(form.String("url"))
	if sms.url == "" || strings.Contains(sms.url, "@") {
		err = fmt.Errorf("args.setting配置错误，url配置错误(url:%s)(%s)", sms.url, content)
		return
	}
	sms.body = form.Translate(form.String("body"))
	if sms.body == "" || strings.Contains(sms.body, "@") {
		err = fmt.Errorf("args.setting配置错误，body配置错误(body:%s)(%s)", sms.body, content)
		return
	}
	auth := form.Translate(form.String("auth"))
	if auth == "" || strings.Contains(auth, "@") {
		err = fmt.Errorf("args.setting配置错误，auth配置错误(auth:%s)(%s)", auth, content)
		return
	}
	form.Transform.Set("auth", base64.Encode(auth))

	headers, err := form.GetArray("header")
	if err != nil {
		err = fmt.Errorf("配置文件中未包含header参数")
		return
	}
	for _, header := range headers {
		v := header.(string)
		hs := strings.SplitN(v, ":", 2)
		if len(hs) != 2 {
			err = fmt.Errorf("args.setting配置错误，header配置错误,不是有效的键值对(header:%s)", header)
			return
		}
		sms.header[hs[0]] = form.Translate(hs[1])
	}
	sms.charset = form.String("charset", "utf-8")
	return

}

//SendSMS 发送短消息
func SendSMS(mobile, data, setting string) (st int, r string, err error) {
	if mobile == "" {
		err = fmt.Errorf("接收人手机号不能为空")
		return
	}
	if data == "" {
		err = fmt.Errorf("短信内容(data)不能为空")
		return
	}
	if setting == "" {
		err = fmt.Errorf("配置内容不能为空")
		return
	}
	m, err := getYtxParams(mobile, data, setting)
	if err != nil {
		return
	}
	client := http.NewHTTPClient()
	r, st, err = client.Request("post", m.url, m.body, m.charset, m.header)
	if err != nil {
		err = fmt.Errorf("%v(url:%s,body:%s,header:%s)", err, m.url, m.body, m.header)
		return
	}
	return
}
