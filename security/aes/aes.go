package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/encoding/base64"
)

const (
	AesECB = "ECB"
	AesCBC = "CBC"
	AesCTR = "CTR"
	AesCFB = "CFB"
	AesOFB = "OFB"
)

const (
	PaddingNull  = "null" //不填充
	PaddingPkcs7 = "pkcs7"
	PaddingPkcs5 = "pkcs5"
	// PaddingZero  = "zero"
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

//Encrypt 加密字符串  默认解密模式CFB,数据不填充
func Encrypt(msg string, key string) (string, error) {
	return EncryptMode(key, msg, AesCFB, PaddingNull)
}

//EncryptMode 自定义加密模式和数据填充模式
func EncryptMode(key, cipherText, mode, padding string) (plainText string, err error) {
	keyBytes := getKey(key)
	blockSize := aes.BlockSize
	var iv = keyBytes[:blockSize]
	aesCipher, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("encryptData.NewCipher:%v", err)
	}

	var cipherBytes = []byte(cipherText)
	switch strings.ToLower(padding) {
	case PaddingPkcs7:
		cipherBytes = pkcs7Padding(cipherBytes, blockSize)
	case PaddingPkcs5:
		cipherBytes = pkcs5Padding(cipherBytes, blockSize)
	default:
		//CFB 可以不进行数据填充
		if mode != AesCFB {
			return "", fmt.Errorf("encryptData.不支持的padding类型:%v", padding)
		}
	}

	var entBytes = make([]byte, len(cipherBytes))
	var stream cipher.Stream
	switch strings.ToUpper(mode) {
	case AesCBC:
		blockMode := cipher.NewCBCEncrypter(aesCipher, iv)
		blockMode.CryptBlocks(entBytes, cipherBytes)
	case AesCTR:
		stream = cipher.NewCTR(aesCipher, iv)
		stream.XORKeyStream(entBytes, cipherBytes)
	case AesCFB:
		stream = cipher.NewCFBEncrypter(aesCipher, iv)
		stream.XORKeyStream(entBytes, cipherBytes)
	case AesOFB:
		stream = cipher.NewOFB(aesCipher, iv)
		stream.XORKeyStream(entBytes, cipherBytes)
	default:
		return "", fmt.Errorf("encryptData.不支持的加密模式:%v", mode)
	}

	plainText = base64.EncodeBytes(entBytes)
	return
}

//Decrypt 解密字符串 默认解密模式CFB,数据不填充
func Decrypt(src string, key string) (msg string, err error) {
	return DecryptMode(key, src, AesCFB, PaddingNull)
}

//DecryptMode 自定义加密模式和数据填充模式
func DecryptMode(key, cipherText, mode, padding string) (plainText string, err error) {
	keyBytes := getKey(key)
	blockSize := aes.BlockSize
	var iv = keyBytes[:blockSize]
	aesCipher, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("decryptData.NewCipher:%v", err)
	}

	var dstBytes []byte
	cipherBytes, err := base64.DecodeBytes(cipherText)
	dstBytes = make([]byte, len(cipherBytes))

	var stream cipher.Stream
	switch strings.ToUpper(mode) {
	case AesCBC:
		blockMode := cipher.NewCBCDecrypter(aesCipher, iv)
		blockMode.CryptBlocks(dstBytes, cipherBytes)
	case AesCTR:
		stream = cipher.NewCTR(aesCipher, iv)
		stream.XORKeyStream(dstBytes, cipherBytes)
	case AesCFB:
		stream = cipher.NewCFBDecrypter(aesCipher, iv)
		stream.XORKeyStream(dstBytes, cipherBytes)
	case AesOFB:
		stream = cipher.NewOFB(aesCipher, iv)
		stream.XORKeyStream(dstBytes, cipherBytes)
	}

	switch strings.ToLower(padding) {
	case PaddingPkcs7:
		dstBytes = pkcs7UnPadding(dstBytes)
	case PaddingPkcs5:
		dstBytes = pkcs5UnPadding(dstBytes)
	default:
		if mode != AesCFB {
			return "", fmt.Errorf("decryptData.不支持的padding类型:%v", padding)
		}
	}

	plainText = strings.TrimSpace(string(dstBytes))
	return
}

func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unPadding := int(plainText[length-1])
	return plainText[:(length - unPadding)]
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unPadding := int(plainText[length-1])
	return plainText[:(length - unPadding)]
}
