package rsa

import (
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

// Encrypt RSA加密
// publicKey 加密时候用到的公钥
func Encrypt(origData string, publicKey string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509 ParsePKIXPublicKey err:%v", err)
	}
	pub := pubInterface.(*rsa.PublicKey)
	data, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		return "", fmt.Errorf("rsa EncryptPKCS1v15 err:%v", err)
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

// Decrypt RSA解密
// privateKey 解密时候用到的秘钥
func Decrypt(ciphertext string, privateKey string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509 ParsePKCS1PrivateKey err:%v", err)
	}
	input, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 URLEncoding DecodeString err:%v", err)
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, input)
	if err != nil {
		return "", fmt.Errorf("rsa DecryptPKCS1v15 err:%v", err)
	}

	return string(data), nil
	/*end*/
}

// Sign 使用RSA生成签名
// privateKey 加密时使用的秘钥	mode 加密的模式[目前只支持MD5，SHA1，不区分大小写]
func Sign(message string, privateKey string, mode string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
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
// publicKey 验证签名的公钥	mode 加密的模式[目前只支持MD5，SHA1，不区分大小写]
func Verify(src string, sign string, publicKey string, mode string) (pass bool, err error) {
	//步骤1，加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
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
