package sms

import (
	"testing"

	"github.com/qxnw/lib4go/ut"
)

func TestSendSMS(t *testing.T) {
	c := `
	{
		"appid": "8aaf07085fe2d98c015febbe09880351",
		"account": "8a216da85fc7a0a4015fdd8871ef077c",
		"token": "28d63e155179469bba8ef25982105e88",
		"ver": "2013-12-26",
		"raw": "{@account}{@token}{@timestamp}",
		"auth": "{@account}:{@timestamp}",
		"header": [
			"Accept:application/xml",
			"Content-type:application/xml",
			"charset:utf-8",
			"Authorization:@auth"
		],
		"url": "https://app.cloopen.com:8883/{@ver}/Accounts/{@account}/SMS/TemplateSMS?sig=@sign",
		"body": "<?xml version='1.0' encoding='utf-8'?><TemplateSMS><to>{@mobile}</to><appId>{@appid}</appId><templateId>227617</templateId><datas>{@data}</datas></TemplateSMS>"
	}`

	st, _, err := SendSMS("15828680877", "52654", c)
	ut.Expect(t, st, 200)
	ut.Expect(t, err, nil)

}
