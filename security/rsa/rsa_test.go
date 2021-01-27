/*
	参考链接：
	1、http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Brsa/
	2、http://studygolang.com/articles/5257
*/
package rsa

import (
	"testing"

	"strings"
)

const (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJAZ3/rkUodqiNuKtGGsJvo68HzDPCMjQYjD0VpaAclwQFL0s7uPUPL
G6qnLg37wweLiamH16hxA4EJfmiK7Uh5oQIDAQABAkAMQBLUzo32Tl1CyiwECWAn
T3yCIpKwOnK54wBX5MiuMF77Rn7ktJqgDINvx37GegNsHpoS7R0EzP+WkHuP5rJ1
AiEAwVbyh23um1kzObks3aicgDj3Umq3MCssYgo3DlnfcA8CIQCJCxvcGA3CIw9j
SDqzcLGKb69D3PrCg7Y6whwYy/fLTwIhALJjfDGTMCZsPkSTZB89NPFmHmUQC+hI
3ZG0JSp7qBrnAiBA+4eGYdGEUOOnDETpeXJ2VmchItO1EIeEbS6tg2pIeQIgXIwF
UgtCj4Sr23PgYHs3VTXHfniZCOv2R4HS1epkKks=
-----END RSA PRIVATE KEY-----`
	publicKey = `-----BEGIN PUBLIC KEY-----
MFswDQYJKoZIhvcNAQEBBQADSgAwRwJAZ3/rkUodqiNuKtGGsJvo68HzDPCMjQYj
D0VpaAclwQFL0s7uPUPLG6qnLg37wweLiamH16hxA4EJfmiK7Uh5oQIDAQAB
-----END PUBLIC KEY-----`
	errPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJAZ3/rkUodqiNuKtGGsJvo68HzDPCMjQYjD0VpaAclwQFL0s7uPUPL
G6qnLg37wweLiamH16hxA4E89NPFmHmUQC+hI
3ZG0JSp7qBrnAiBA+4eGYdGEUOOnDETpeXJ2VmchItO1EIeEbS6tg2pIeQIgXIwF
UgtCj4Sr23PgYHs3VTXHfniZCOv2R4HS1epkKks=
-----END RSA PRIVATE KEY-----`
	errPublicKey = `-----BEGIN PUBLIC KEY-----
MFswDQYJKoZIhvcNAQEBBQADSgAwRwJAZ3/rkUodqiNuKtGGsJvo68HzDPCMjQYj
-----END PUBLIC KEY-----`
	errPrivateKey2 = `-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJAZ3/dbUodqiNuKtGGsJvo68HzDPCMjQYjD0VpaAclwQFL0s7uPUPL
G6qnLg37wweLiamH16hxA4EJfmiK7Uh5oQIDAQABAkAMQBLUzo32Tl1CyiwECWAn
T3yCIpKwOnK54wBX5MiuMF77Rn7ktJqgDINvx37GegNsHpoS7R0EzP+WkHuP5rJ1
AiEAwVbyh23um1kzObks3aicgDj3Umq3MCssYgo3DlnfcA8CIQCJCxvcGA3CIw9j
SDqzcLGKb69D3PrCg7Y6whwYy/12TwIhALJjfDGTMCZsPkSTZB89NPFmHmUQC+hI
3ZG0JSp7qBrnAiBA+4eGYqeEUOOnDETpeXJ2VmchItO1EIeEbS6tg2pIeQIgXIwF
UgtCj4Sr23PgYHs3VTXHfniZCOv2R4HS1epkKks=
-----END RSA PRIVATE KEY-----`
	errPublickey2 = `-----BEGIN PUBLIC KEY-----
MFswDQYJKoZIhvcNAQEBBQADSgAwRwJAV2/rkUodqiNuKtGGsJvo68HzDPCMjQYj
D0VpaAclwQFL0s7uPUPLG6qnLg37wweLiamH16hxA4EJfmiK7Uh5oQIDAQWE
-----END PUBLIC KEY-----`
)

func TestEncrypt(t *testing.T) {
	input := "hello worldxxx"
	data, err := Encrypt(input, publicKey, PKCS8)
	if err != nil {
		t.Errorf("Encrypt fail:%v", err)
	}
	actual, err := Decrypt(data, privateKey, PKCS1)
	if err != nil {
		t.Errorf("Decrypt fail:%v", err)
	}
	if !strings.EqualFold(actual, input) {
		t.Errorf("RSA test fail %s to %s", input, actual)
	}

	input = ""
	data, err = Encrypt(input, publicKey, PKCS8)
	if err != nil {
		t.Errorf("Encrypt fail:%v", err)
	}
	actual, err = Decrypt(data, privateKey, PKCS1)
	if err != nil {
		t.Errorf("Decrypt fail:%v", err)
	}
	if !strings.EqualFold(actual, input) {
		t.Errorf("RSA test fail %s to %s", input, actual)
	}

	input = "hello world"
	data = "AxL4pO+oVtK1XlQU0GpfgWui1WfWHPX37urOKh9oeg0naQqMqV1dEPteSlNxWRqp9y1eu9YOmGroZCcx5p/rSA=="
	actual, err = Decrypt(data, privateKey, PKCS1)
	if err != nil {
		t.Errorf("Decrypt fail:%v", err)
	}
	if strings.EqualFold(actual, input) {
		t.Errorf("RSA test fail %s to %s", input, actual)
	}

	input = "hello world"
	_, err = Encrypt(input, errPublicKey, PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello world"
	data, err = Encrypt(input, publicKey, PKCS8)
	if err != nil {
		t.Errorf("Encrypt fail:%v", err)
	}
	_, err = Decrypt(data, errPrivateKey, PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello world"
	data, err = Encrypt(input, errPublickey2, PKCS8)
	if err != nil {
		t.Errorf("Encrypt fail:%v", err)
	}
	actual, err = Decrypt(data, errPrivateKey2, PKCS1)
	if err == nil {
		t.Error("test fail")
	}
	if strings.EqualFold(actual, input) {
		t.Errorf("RSA test fail %s to %s", input, actual)
	}
}

func TestSign(t *testing.T) {
	input := "hello"
	mode := "md5"
	data, err := Sign(input, privateKey, mode, PKCS1)
	if err != nil {
		t.Errorf("Sign fail %v", err)
	}
	actual, err := Verify(input, data, publicKey, mode, PKCS8)
	if err != nil {
		t.Errorf("Sign test fail %v", err)
	}
	if !actual {
		t.Error("Sign test fail")
	}

	input = "hello"
	mode = "sha1"
	data, err = Sign(input, privateKey, mode, PKCS1)
	if err != nil {
		t.Errorf("Sign fail %v", err)
	}
	actual, err = Verify(input, data, publicKey, mode, PKCS8)
	if err != nil {
		t.Errorf("Sign test fail %v", err)
	}
	if !actual {
		t.Error("Sign test fail")
	}

	input = ""
	mode = "sha1"
	data, err = Sign(input, privateKey, mode, PKCS1)
	if err != nil {
		t.Errorf("Sign fail %v", err)
	}
	actual, err = Verify(input, data, publicKey, mode, PKCS8)
	if err != nil {
		t.Errorf("Sign test fail %v", err)
	}
	if !actual {
		t.Error("Sign test fail")
	}

	input = "hello"
	mode = "base64"
	_, err = Sign(input, privateKey, mode, PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello"
	mode = "sha1"
	data, err = Sign(input, privateKey, mode, PKCS1)
	if err != nil {
		t.Errorf("Sign fail %v", err)
	}
	_, err = Verify(input, data, publicKey, "base64", PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello"
	mode = "sha1"
	_, err = Sign(input, errPrivateKey, mode, PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello"
	mode = "sha1"
	data, err = Sign(input, privateKey, mode, PKCS1)
	if err != nil {
		t.Errorf("Sign fail %v", err)
	}
	_, err = Verify(input, data, errPublicKey, mode, PKCS1)
	if err == nil {
		t.Error("Sign test fail")
	}

	input = "hello"
	mode = "sha1"
	data, err = Sign(input, errPrivateKey2, mode, PKCS1)
	if err == nil {
		t.Error("Sign test fail")
	}
	actual, err = Verify(input, data, errPublickey2, mode, PKCS1)
	if err == nil {
		t.Error("Sign test fail")
	}
	if actual {
		t.Error("Sign test fail")
	}
}
