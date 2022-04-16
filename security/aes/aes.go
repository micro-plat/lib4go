package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/encoding/base64"
	"github.com/micro-plat/lib4go/security/padding"
	"github.com/micro-plat/lib4go/types"
)

const (
	AesECB = "ECB"
	AesCBC = "CBC"
	AesCTR = "CTR"
	AesCFB = "CFB"
	AesOFB = "OFB"
)

// Encrypt 加密字符串
// mode 加密类型/填充模式,不传默认为:CFB/NULL
// key 加密密钥[字符串长度必须是大于16,且是8的倍数]
// 加密类型:ECB,CBC,CTR,CFB,OFB
// 填充模式:PKCS7,PKCS5,ZERO,NULL(只有CFB加密模式支持NULL不填充)
// iv偏移量:默认 []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
func Encrypt(msg string, key string, mode ...string) (plainText string, err error) {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return EncryptByte(key, []byte(msg), iv, mode...)
}

// EncryptByte 加密字符串
// mode 加密类型/填充模式,不传默认为:CFB/NULL
// key 加密密钥[字符串长度必须是大于16,且是8的倍数]
// 加密类型:ECB,CBC,CTR,CFB,OFB
// 填充模式:PKCS7,PKCS5,ZERO,NULL(只有CFB加密模式支持NULL不填充)
// iv:偏移量[字节长度必须是16]
func EncryptByte(key string, cipherText, iv []byte, mode ...string) (plainText string, err error) {
	keyBytes := []byte(key)
	blockSize := aes.BlockSize
	aesCipher, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("encryptData.NewCipher:%v", err)
	}

	cmode := types.GetStringByIndex(mode, 0, fmt.Sprintf("%s/%s", AesCFB, padding.PaddingNull))
	m, p, err := padding.GetModePadding(cmode)
	if err != nil {
		return "", err
	}

	var cipherBytes = cipherText
	switch p {
	case padding.PaddingPkcs7:
		cipherBytes = padding.PKCS7Padding(cipherBytes)
	case padding.PaddingPkcs5:
		cipherBytes = padding.PKCS5Padding(cipherBytes, blockSize)
	case padding.PaddingZero:
		cipherBytes = padding.ZeroPadding(cipherBytes, blockSize)
	default:
		//CFB 可以不进行数据填充
		if m != AesCFB {
			return "", fmt.Errorf("encryptData.不支持的padding类型:%v", p)
		}
	}

	var entBytes = make([]byte, len(cipherBytes))
	var stream cipher.Stream
	switch m {
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
		return "", fmt.Errorf("encryptData.不支持的加密模式:%v", cmode)
	}

	plainText = base64.EncodeBytes(entBytes)
	return
}

// Decrypt 解密字符串
// key 解密密钥[字符串长度必须是大于16,且是8的倍数]
// mode 解密类型/填充模式,不传默认为:CFB/NULL
// 解密类型:ECB,CBC,CTR,CFB,OFB
// 填充模式:PKCS7,PKCS5,ZERO,NULL(只有CFB解密模式支持NULL不填充)
// iv偏移量:默认 []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
func Decrypt(src string, key string, mode ...string) (msg string, err error) {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return DecryptByte(key, []byte(src), iv, mode...)
}

// DecryptByte 解密字符串
// key 解密密钥[字符串长度必须是大于16,且是8的倍数]
// mode 解密类型/填充模式,不传默认为:CFB/NULL
// 解密类型:ECB,CBC,CTR,CFB,OFB
// 填充模式:PKCS7,PKCS5,ZERO,NULL(只有CFB解密模式支持NULL不填充)
// iv:偏移量[字节长度必须是16]
func DecryptByte(key string, cipherText, iv []byte, mode ...string) (plainText string, err error) {
	keyBytes := []byte(key)
	aesCipher, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("DecryptByte.NewCipher:%v", err)
	}

	cmode := types.GetStringByIndex(mode, 0, fmt.Sprintf("%s/%s", AesCFB, padding.PaddingNull))

	m, p, err := padding.GetModePadding(cmode)
	if err != nil {
		return "", fmt.Errorf("DecryptByte.GetModePadding:%s,%v", cmode, err)
	}

	var dstBytes []byte
	cipherBytes, err := base64.DecodeBytes(string(cipherText))
	if err != nil {
		return "", fmt.Errorf("DecryptByte.DecodeBytes:%v", err)
	}
	dstBytes = make([]byte, len(cipherBytes))

	var stream cipher.Stream
	switch m {
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
	default:
		return "", fmt.Errorf("DecryptByte.不支持的解密类型:%v", m)
	}

	switch p {
	case padding.PaddingPkcs7:
		dstBytes = padding.PKCS7UnPadding(dstBytes)
	case padding.PaddingPkcs5:
		dstBytes = padding.PKCS5UnPadding(dstBytes)
	case padding.PaddingZero:
		dstBytes = padding.ZeroUnPadding(dstBytes)
	default:
		if m != AesCFB {
			return "", fmt.Errorf("DecryptByte.不支持的padding类型:%v", p)
		}
	}

	plainText = strings.TrimSpace(string(dstBytes))
	return
}
