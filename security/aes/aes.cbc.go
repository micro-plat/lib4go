package aes

// EncryptCBCPKCS7WithIV CBC模式,PKCS7填充
func EncryptCBCPKCS7WithIV(contentStr string, keyStr string, iv []byte) (string, error) {
	return EncryptMode(keyStr, contentStr, AesCBC, PaddingPkcs7)
}

// DecryptCBCPKCS7WithIV CBC模式,PKCS7填充
func DecryptCBCPKCS7WithIV(contentStr string, keyStr string, iv []byte) (string, error) {
	return DecryptMode(keyStr, contentStr, AesCBC, PaddingPkcs7)
}
