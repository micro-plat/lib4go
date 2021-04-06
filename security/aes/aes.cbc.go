package aes

import (
	"fmt"

	"github.com/micro-plat/lib4go/security/padding"
)

// EncryptCBCPKCS7WithIV CBC模式,PKCS7填充
// key 加密密钥[字符串长度必须是大于16,且是8的倍数]
// iv:偏移量[字节长度必须是16]
func EncryptCBCPKCS7WithIV(contentStr string, keyStr string, iv []byte) (string, error) {
	return EncryptByte(keyStr, []byte(contentStr), iv, fmt.Sprintf("%s/%s", AesCBC, padding.PaddingPkcs7))
}

// DecryptCBCPKCS7WithIV CBC模式,PKCS7填充
// key 加密密钥[字符串长度必须是大于16,且是8的倍数]
// iv:偏移量[字节长度必须是16]
func DecryptCBCPKCS7WithIV(contentStr string, keyStr string, iv []byte) (string, error) {
	return DecryptByte(keyStr, []byte(contentStr), iv, fmt.Sprintf("%s/%s", AesCBC, padding.PaddingPkcs7))
}
