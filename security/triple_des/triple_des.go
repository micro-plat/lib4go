package tripleDes

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"strings"

	xdes "github.com/qxnw/lib4go/security/des"
)

//Encrypt 3des加密cbs/pkcs5
func Encrypt(input string, skey string, mode string) (r string, err error) {
	secMode, padding, err := getModes(mode)
	if err != nil {
		return
	}
	origData := []byte(input)
	key := []byte(skey)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("des NewTripleDESCipher err:%v", err)
	}

	switch padding {
	case "pkcs5":
		origData = xdes.PKCS5Padding(origData, block.BlockSize())
	case "pkcs7":
		origData = xdes.PKCS7Padding(origData)
	case "zero":
		origData = xdes.ZeroPadding(origData, block.BlockSize())
	default:
		err = fmt.Errorf("不支持的填充模式:%s", padding)
		return
	}

	switch secMode {
	case "cbc":
		iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
		blockMode := cipher.NewCBCEncrypter(block, iv)
		crypted := make([]byte, len(origData))
		blockMode.CryptBlocks(crypted, origData)
		r = strings.ToUpper(hex.EncodeToString(crypted))
		return
	case "ecb":
		blockMode := xdes.NewECBEncrypter(block)
		crypted := make([]byte, len(origData))
		blockMode.CryptBlocks(crypted, origData)
		r = strings.ToUpper(hex.EncodeToString(crypted))
		return
	case "cfb":
		iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
		blockMode := cipher.NewCFBEncrypter(block, iv)
		crypted := make([]byte, len(origData))
		blockMode.XORKeyStream(crypted, origData)
		r = strings.ToUpper(hex.EncodeToString(crypted))
		return
	default:
		err = fmt.Errorf("不支持的加密模式:%s", secMode)
		return

	}
}

//Decrypt 3des解密cbs/pkcs5
func Decrypt(input string, skey string, mode string) (r string, err error) {
	secMode, padding, err := getModes(mode)
	if err != nil {
		return
	}
	crypted, err := hex.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("des DecodeString err:%v", err)
	}
	key := []byte(skey)[:24]
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("des NewTripleDESCipher err:%v", err)
	}
	origData := make([]byte, len(crypted))
	switch secMode {
	case "cbc":
		iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
		blockMode := cipher.NewCBCDecrypter(block, iv)
		origData := make([]byte, len(crypted))
		blockMode.CryptBlocks(origData, crypted)
	case "ecb":
		blockMode := xdes.NewECBDecrypter(block)
		blockMode.CryptBlocks(origData, crypted)
	case "cfb":
		iv := []byte{0, 0, 0, 0, 0, 0, 0, 0}
		blockMode := cipher.NewCFBDecrypter(block, iv)
		blockMode.XORKeyStream(origData, crypted)
	default:
		err = fmt.Errorf("不支持的加密模式:%s", secMode)
		return

	}

	switch padding {
	case "pkcs5":
		origData = xdes.PKCS5UnPadding(origData)
	case "pkcs7":
		origData = xdes.PKCS7UnPadding(origData)
	case "zero":
		origData = xdes.ZeroUnPadding(origData)
	default:
		err = fmt.Errorf("不支持的填充模式:%s", padding)
		return
	}
	r = string(origData)
	return
}
func getModes(mode string) (sec string, padding string, err error) {
	if len(mode) == 0 {
		return "cbc", "pkcs5", nil
	}
	modes := strings.SplitN(mode, "/", 2)
	if len(modes) != 2 {
		return "", "", fmt.Errorf("输入的加解密模式不正确:%s", mode)
	}
	if modes[0] != "cbc" && modes[0] != "ecb" && modes[0] != "cfb" {
		return "", "", fmt.Errorf("输入的加解密模式不正确:%s只支持cbs,ecb,cfb", modes[0])
	}
	if modes[1] != "pkcs5" && modes[1] != "pkcs7" && modes[1] != "zero" {
		return "", "", fmt.Errorf("输入的填充模式不正确:%s只支持pkcs5,pkcs7,zero", modes[0])
	}
	return modes[0], modes[1], nil
}
