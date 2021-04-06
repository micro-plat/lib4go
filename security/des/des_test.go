package des

import (
	"fmt"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestEncrypt(t *testing.T) {
	input := []byte("hello")
	key := "12345678"
	iv := []byte("87654321")
	crypted, err := EncryptBytes(input, key, iv, "cbc/pkcs5")
	assert.Equal(t, err, nil)

	o, err := DecryptBytes(crypted, key, iv, "cbc/pkcs5")
	assert.Equal(t, err, nil)
	assert.Equal(t, string(o), string(input))

	pk := []byte("encrypt:ecb/pkcs7")
	fmt.Println(len(pk))
}

func TestDecrypt(t *testing.T) {
	type args struct {
		cipherText string
		key        string
		mode       []string
	}
	tests := []struct {
		name          string
		args          args
		wantR         string
		wantPlainText string
		wantErr       bool
		wantErr1      bool
	}{
		{name: "1.1 ECB-pkcs5加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"ECB/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.2 ECB-pkcs7加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"ECB/PKCS7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.3 ECB-zero填充加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"ECB/zero"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.4 ECB-null不填充加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"ECB/NULL"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
		{name: "1.5 ECB-不传入mode加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.6 ECB-mode为mil加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "2.1 CBC-pkcs5加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"CBC/pkcs5"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.2 CBC-pkcs7加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"CBC/pkcs7"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.3 CBC-zero填充加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"CBC/ZERO"}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.4 CBC-null不填充加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"CBC/NULL"}}, wantPlainText: "", wantErr: true, wantErr1: true},
		{name: "2.5 CBC-不传入mode加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{""}}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.6 CBC-mode为mil加解密", args: args{key: "13245687", cipherText: "中国李宁", mode: nil}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "5.1 错误的加解密码类型", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"ERR/pkcs5"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
		{name: "5.2 错误的填充类型", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"OFB/err"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
		{name: "5.3 错误的加解密和填充类型", args: args{key: "13245687", cipherText: "中国李宁", mode: []string{"err/err"}}, wantPlainText: "中国李宁", wantErr: true, wantErr1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlainText, err := Encrypt(tt.args.cipherText, tt.args.key, tt.args.mode...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			resText, err := Decrypt(gotPlainText, tt.args.key, tt.args.mode...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resText != tt.wantPlainText {
				t.Errorf("加解密结果不正确 = %v, want %v", gotPlainText, tt.wantPlainText)
			}
		})
	}
}
