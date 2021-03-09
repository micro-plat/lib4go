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

func TestGenerateKey(t *testing.T) {
	priKey, PubKey, err := GenerateKey("PKCS1", 1024)
	if err != nil {
		t.Errorf("GenerateKey fail:%v", err)
	}

	t.Error("priKey:", priKey)
	t.Error("PubKey:", PubKey)
}

func TestEncrypt(t *testing.T) {
	publicKey := `-----BEGIN PUBLIC KEY-----
	MIGJAoGBAMmFSZdIaoE2j/QT93zM1g4t5H7cA19+c+H0iSUwsEYDmYsKbm6VW92O
	lHY6+cgid0iDZp4QllKMAjPiJWAYmF01sYYKlEl+9a5GO37EF+iQBJ/b0L2gRhLJ
	CBgySIF6s4KQnS2B8mk+4prhJKfG8cKMI5aWXthWpAuodttglzPbAgMBAAE=
	-----END PUBLIC KEY-----`
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
	MIICXAIBAAKBgQDJhUmXSGqBNo/0E/d8zNYOLeR+3ANffnPh9IklMLBGA5mLCm5u
	lVvdjpR2OvnIIndIg2aeEJZSjAIz4iVgGJhdNbGGCpRJfvWuRjt+xBfokASf29C9
	oEYSyQgYMkiBerOCkJ0tgfJpPuKa4SSnxvHCjCOWll7YVqQLqHbbYJcz2wIDAQAB
	AoGAA+L+OFy9MSDMRfjcnRuWRU+9SHUV25Gkyobc3krCG5eWLohU+O0IiI1nb6BT
	kPiZNFzUbdgEDjOFF1sVPXU7+wdwcWyggO+Mdl6Q6aulUaHAg/B71aIE+kMMaWbY
	Qj71s0PeXw5uRx+n+i3Ro+TEBOghzKf1a5U/NL5weGUq9gECQQD6brqZBM6s4uvV
	bh4PNnvYSo/SQZX1A6Rx/b3S/UBkf5vP+xdYNLI0NxGeul7oIJEGied75dHU7v4q
	kFcnMaErAkEAzgAwfiHL0pmphgLMVxxOHvz16TMularhX2kL8MfK923HpmCvHVGn
	QfDV25DS6i01irkei3hO4LzAMlARqiKAEQJBAIN/GOO4Ln2BOav8AjSiuyy7GgGh
	Boh8vSBNyBq9d85NYxc2FO/v25KnR808txDT6NKyHqZj6mYQh8z5tYmS+bkCQCve
	FnWFtOXQGy2Sgvk56djnfWZ/o7fzf7LVp9lKcopmMlHX3PKdZMTCCIiNOpzrq68y
	5LJGmGV7TGJqcpiMaEECQEt8v8d3ZkCO8xJ5f3JMFTrK5DROSrwYCv0oijEf/9Bo
	fOAeVAzwsfIXxBNJp9ycRdJtec7bDKvjoUdxYaXRlwc=
	-----END RSA PRIVATE KEY-----`
	input := "hello worldxxx"
	data, err := Encrypt(input, publicKey, PKCS1)
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
	data, err = Encrypt(input, publicKey, PKCS1)
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
	_, err = Encrypt(input, errPublicKey, PKCS1)
	if err == nil {
		t.Error("test fail")
	}

	input = "hello world"
	data, err = Encrypt(input, publicKey, PKCS1)
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
