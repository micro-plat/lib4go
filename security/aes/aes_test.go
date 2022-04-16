package aes

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestAES(t *testing.T) {

	origin := "d6f3b43fc160e2d4ed05539e75509e95"
	key := "d6f3b43fc160e2d4"
	expect := "Psc6tWRgMje0IT7g1v/tyqQNPE3XuDsW2MSIJ2Mu8xoFQodfrtQfaWMTHMvNdQKC"
	e, err := Encrypt(origin, key, "CBC/PKCS5")
	assert.Equal(t, nil, err, err)
	assert.Equal(t, expect, e)
}

func Test_encryptData(t *testing.T) {

	type args struct {
		key        string
		cipherText string
		mode       []string
	}
	tests := []struct {
		name          string
		args          args
		wantPlainText string
		wantErr       bool
		wantErr1      bool
	}{
		{name: "1.1 CFB-pkcs5加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CFB/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.2 CFB-pkcs7加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CFB/PKCS7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.3 CFB-zero填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CFB/zero"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.4 CFB-null不填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CFB/NULL"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.5 CFB-不传入mode加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "1.6 CFB-mode为mil加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "2.1 CBC-pkcs5加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CBC/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.2 CBC-pkcs7加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CBC/pkcs7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.3 CBC-zero填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CBC/ZERO"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.4 CBC-null不填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CBC/NULL"}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "2.5 CBC-不传入mode加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "2.6 CBC-mode为mil加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "3.1 CTR-pkcs5加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CTR/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "3.2 CTR-pkcs7加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CTR/pkcs7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "3.3 CTR-zero填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CTR/zero"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "3.4 CTR-null不填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"CTR/null"}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "3.5 CTR-不传入mode加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "3.6 CTR-mode为mil加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "4.1 OFB-pkcs5加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"OFB/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "4.2 OFB-pkcs7加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"OFB/pkcs7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "4.3 OFB-zero填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"OFB/zero"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "4.4 OFB-null不填充加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"OFB/null"}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "4.5 OFB-不传入mode加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "4.6 OFB-mode为mil加解密", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "5.1 错误的加解密码类型", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"ERR/pkcs5"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
		{name: "5.2 错误的填充类型", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"OFB/err"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
		{name: "5.3 错误的加解密和填充类型", args: args{key: "1324568789123456", cipherText: "中国李宁", mode: []string{"err/err"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlainText, err := Encrypt(tt.args.cipherText, tt.args.key, tt.args.mode...)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			resText, err := Decrypt(gotPlainText, tt.args.key, tt.args.mode...)
			if (err != nil) != tt.wantErr {
				t.Errorf("decryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resText != tt.wantPlainText {
				t.Errorf("加解密结果不正确 = %v, want %v", gotPlainText, tt.wantPlainText)
			}
		})
	}
}
