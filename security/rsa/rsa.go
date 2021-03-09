package rsa

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	PKCS1 = "PKCS1"
	PKCS8 = "PKCS8"
)

//GenerateKey 生成基于pkcs1的rsa私、公钥对
//pkcsType 密钥格式类型:PKCS1,PKCS8
//bits 密钥位数:1024,2048
func GenerateKey(pkcsType string, bits int) (prikey string, pubkey string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	switch pkcsType {
	case PKCS1:
		prikey = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
		pubkey = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	case PKCS8:
		data, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return "", "", err
		}
		prikey = base64.StdEncoding.EncodeToString(data)
		data, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			return "", "", err
		}
		pubkey = base64.StdEncoding.EncodeToString(data)
	default:
		err = fmt.Errorf("不支持的莫要生成格式[%s]", pkcsType)
	}

	prikey = FormatPrivateKey(prikey, pkcsType)
	pubkey = FormatPublicKey(pubkey)
	return
}

// Encrypt RSA加密
// origData 加密原串
// publicKey 加密时候用到的公钥
// pkcsType 密钥格式类型:PKCS1,PKCS8
func Encrypt(origData, publicKey, pkcsType string) (data string, err error) {
	pub, err := getPublicKey(publicKey, pkcsType)
	if err != nil {
		return "", fmt.Errorf("rsa Encrypt getPublicKey err:%v", err)
	}

	res, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		return "", fmt.Errorf("rsa EncryptPKCS1v15 err:%v", err)
	}
	data = base64.StdEncoding.EncodeToString(res)
	return
}

// Decrypt RSA解密
// ciphertext 解密数据原串
// privateKey 解密时候用到的秘钥
// pkcsType 密钥格式类型:PKCS1,PKCS8
func Decrypt(ciphertext, privateKey, pkcsType string) (data string, err error) {
	priv, err := getPrivateKey(privateKey, pkcsType)
	if err != nil {
		return "", fmt.Errorf("rsa Decrypt getPrivateKey err:%v", err)
	}

	input, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 StdEncoding DecodeString err:%v", err)
	}
	res, err := rsa.DecryptPKCS1v15(rand.Reader, priv, input)
	if err != nil {
		return "", fmt.Errorf("rsa DecryptPKCS1v15 err:%v", err)
	}

	data = string(res)
	return
}

// Sign 使用RSA生成签名
// message 签名数据原串
// privateKey 签名时候用到的秘钥
// mode 加密的模式[目前只支持MD5，SHA1，SHA256,不区分大小写]
// pkcsType 密钥格式类型:PKCS1,PKCS8
func Sign(message, privateKey, mode, pkcsType string) (string, error) {
	priv, err := getPrivateKey(privateKey, pkcsType)
	if err != nil {
		return "", fmt.Errorf("rsa Sign getPrivateKey err:%v", err)
	}

	switch strings.ToLower(mode) {
	case "sha256":
		t := sha256.New()
		io.WriteString(t, message)
		digest := t.Sum(nil)
		data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, digest)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(data), nil
	case "sha1":
		t := sha1.New()
		io.WriteString(t, message)
		digest := t.Sum(nil)
		data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, digest)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(data), nil

	case "md5":
		t := md5.New()
		io.WriteString(t, message)
		digest := t.Sum(nil)
		data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, digest)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(data), nil
	default:
		return "", errors.New("签名模式不支持")
	}
}

// Verify 校验签名
// src 签名认证数据原串
// sign 签名串
// publicKey 验证签名的公钥
// mode 加密的模式[目前只支持MD5，SHA1，SHA256,不区分大小写]
// pkcsType 密钥格式类型:PKCS1,PKCS8
func Verify(src, sign, publicKey, mode, pkcsType string) (pass bool, err error) {
	//步骤1，加载RSA的公钥
	rsaPub, err := getPublicKey(publicKey, pkcsType)
	if err != nil {
		return false, fmt.Errorf("rsa Encrypt getPublicKey err:%v", err)
	}

	data, _ := base64.StdEncoding.DecodeString(sign)
	switch strings.ToLower(mode) {
	case "sha256":
		t := sha256.New()
		io.WriteString(t, src)
		digest := t.Sum(nil)
		err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, digest, data)
	case "sha1":
		t := sha1.New()
		io.WriteString(t, src)
		digest := t.Sum(nil)
		err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	case "md5":
		t := md5.New()
		io.WriteString(t, src)
		digest := t.Sum(nil)
		err = rsa.VerifyPKCS1v15(rsaPub, crypto.MD5, digest, data)
	default:
		err = errors.New("验签模式不支持")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func getPrivateKey(privateKey, pkcsType string) (priv *rsa.PrivateKey, err error) {
	privateKey = FormatPrivateKey(privateKey, pkcsType)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = errors.New("private key error")
		return
	}

	switch pkcsType {
	case PKCS1:
		priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			err = fmt.Errorf("x509 ParsePKCS1PrivateKey err: %v", err)
			return
		}
	case PKCS8:
		priInterface, err1 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err1 != nil {
			err = fmt.Errorf("x509 ParsePKCS8PrivateKey:err %v", err1)
			return
		}
		priv = priInterface.(*rsa.PrivateKey)
	default:
		err = fmt.Errorf("ras解密不支持的密钥格式[%s]", pkcsType)
		return
	}

	return
}

func getPublicKey(publicKey, pkcsType string) (pub *rsa.PublicKey, err error) {
	publicKey = FormatPublicKey(publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		err = errors.New("public key error")
		return
	}

	switch pkcsType {
	case PKCS1:
		pub, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			err = fmt.Errorf("x509 ParsePKCS1PublicKey err: %v", err)
			return
		}
	case PKCS8:
		pubInterface, err1 := x509.ParsePKIXPublicKey(block.Bytes)
		if err1 != nil {
			err = fmt.Errorf("x509 ParsePKIXPublicKey err: %v", err1)
			return
		}
		pub = pubInterface.(*rsa.PublicKey)
	default:
		err = fmt.Errorf("ras加密不支持的密钥格式[%s]", pkcsType)
		return
	}
	return
}

const (
	kPublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	kPublicKeySuffix = "-----END PUBLIC KEY-----"

	kPKCS1Prefix = "-----BEGIN RSA PRIVATE KEY-----"
	KPKCS1Suffix = "-----END RSA PRIVATE KEY-----"

	kPKCS8Prefix = "-----BEGIN PRIVATE KEY-----"
	KPKCS8Suffix = "-----END PRIVATE KEY-----"
)

func FormatPrivateKey(privateKey, pkcsType string) string {
	switch pkcsType {
	case PKCS1:
		privateKey = strings.Replace(privateKey, kPKCS8Prefix, "", 1)
		privateKey = strings.Replace(privateKey, KPKCS8Suffix, "", 1)
		return formatKey(privateKey, kPKCS1Prefix, KPKCS1Suffix, 64)
	case PKCS8:
		privateKey = strings.Replace(privateKey, kPKCS1Prefix, "", 1)
		privateKey = strings.Replace(privateKey, KPKCS1Suffix, "", 1)
		return formatKey(privateKey, kPKCS8Prefix, KPKCS8Suffix, 64)
	default:
		return ""
	}
}

func FormatPublicKey(raw string) string {
	return formatKey(raw, kPublicKeyPrefix, kPublicKeySuffix, 64)
}

func RemovePriKeyFix(privateKey, pkcsType string) string {
	switch pkcsType {
	case PKCS1:
		privateKey = strings.Replace(privateKey, kPKCS8Prefix, "", 1)
		privateKey = strings.Replace(privateKey, KPKCS8Suffix, "", 1)
		return removeKeyFix(privateKey, kPKCS1Prefix, KPKCS1Suffix)
	case PKCS8:
		privateKey = strings.Replace(privateKey, kPKCS1Prefix, "", 1)
		privateKey = strings.Replace(privateKey, KPKCS1Suffix, "", 1)
		return removeKeyFix(privateKey, kPKCS8Prefix, KPKCS8Suffix)
	default:
		return ""
	}
}

func RemovePubKeyFix(raw string) string {
	return removeKeyFix(raw, kPublicKeyPrefix, kPublicKeySuffix)
}

func formatKey(raw, prefix, suffix string, lineCount int) string {
	if raw == "" {
		return ""
	}

	raw = removeKeyFix(raw, prefix, suffix)
	var sl = len(raw)
	var c = sl / lineCount
	if sl%lineCount > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * lineCount
		var e = b + lineCount
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.String()
}

func removeKeyFix(raw, prefix, suffix string) string {
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)
	return raw
}
