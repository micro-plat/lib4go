package aes

import "testing"

func Test_encryptData(t *testing.T) {

	type args struct {
		key        string
		cipherText string
		mode       string
		padding    string
	}
	tests := []struct {
		name          string
		args          args
		wantPlainText string
		wantErr       bool
		wantErr1      bool
	}{
		{name: "1.1 CFB-pkcs5加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CFB", padding: "pkcs5"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.2 CFB-pkcs7加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CFB", padding: "pkcs7"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "1.3 CFB-null不填充加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CFB", padding: "null"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},

		{name: "2.1 CBC-pkcs5加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CBC", padding: "pkcs5"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.2 CBC-pkcs7加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CBC", padding: "pkcs7"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "2.3 CBC-null不填充加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CBC", padding: "null"}, wantPlainText: "", wantErr: true, wantErr1: true},

		{name: "3.1 CTR-pkcs5加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CTR", padding: "pkcs5"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "3.2 CTR-pkcs7加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CTR", padding: "pkcs7"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "3.3 CTR-null不填充加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "CTR", padding: "null"}, wantPlainText: "", wantErr: true, wantErr1: true},

		{name: "4.1 OFB-pkcs5加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "OFB", padding: "pkcs5"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "4.2 OFB-pkcs7加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "OFB", padding: "pkcs7"}, wantPlainText: "中国李宁", wantErr: false, wantErr1: false},
		{name: "4.3 OFB-null不填充加解密", args: args{key: "13245687891234567", cipherText: "中国李宁", mode: "OFB", padding: "null"}, wantPlainText: "", wantErr: true, wantErr1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlainText, err := EncryptMode(tt.args.key, tt.args.cipherText, tt.args.mode, tt.args.padding)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			resText, err := DecryptMode(tt.args.key, gotPlainText, tt.args.mode, tt.args.padding)
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
