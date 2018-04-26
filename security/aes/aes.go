package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"

	"github.com/qxnw/lib4go/encoding/base64"
)

func getKey(key string) []byte {
	arrKey := []byte(key)
	keyLen := len(key)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//Encrypt 加密字符串
func Encrypt(msg string, key string) (string, error) {
	keyBytes := getKey(key)
	var iv = keyBytes[:aes.BlockSize]
	encrypted := make([]byte, len(msg))
	aesBlockEncrypter, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(msg))
	return base64.Encode(string(encrypted)), nil
}

//Decrypt 解密字符串
func Decrypt(src string, key string) (msg string, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	content, err := base64.Decode(src)
	if err != nil {
		return
	}
	keyBytes := getKey(key)
	var iv = keyBytes[:aes.BlockSize]
	decrypted := make([]byte, len(content))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, []byte(content))
	return strings.TrimSpace(string(decrypted)), nil
}
