package weixin

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WechatEntity struct {
	token          string
	appID          string
	encodingAESKey string
	aesKey         []byte
}

func (e WechatEntity) encodingAESKey2AESKey(encodingKey string) (data []byte, err error) {
	data, err = base64.StdEncoding.DecodeString(encodingKey + "=")
	return
}

//NewWechatEntity 设置初始参数
func NewWechatEntity(appid string, token string, encodingAESKey string) (e WechatEntity, err error) {
	e = WechatEntity{}
	e.appID = appid
	e.token = token
	e.encodingAESKey = encodingAESKey
	e.aesKey, err = e.encodingAESKey2AESKey(e.encodingAESKey)
	return
}

type TextRequestBody struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        `json:"toUserName"`
	FromUserName string        `json:"fromUserName"`
	CreateTime   time.Duration `json:"createTime"`
	MsgType      string        `json:"msgType"`
	Content      string        `json:"content"`
	MsgId        int           `json:"msgId"`
	Event        string
	EventKey     string
	Latitude     string
	Longitude    string
	Precision    string
	Ticket       string
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   string
	MsgType      CDATAText
	Content      CDATAText
}

type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

type EncryptResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATAText
	MsgSignature CDATAText
	TimeStamp    string
	Nonce        CDATAText
}

type EncryptResponseBody1 struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func (e WechatEntity) makeSignature(timestamp, nonce string) string {
	sl := []string{e.token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (e WechatEntity) makeMsgSignature(timestamp, nonce, msg_encrypt string) string {
	sl := []string{e.token, timestamp, nonce, msg_encrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (e WechatEntity) validateUrl(timestamp, nonce, signatureIn string) bool {
	signatureGen := e.makeSignature(timestamp, nonce)
	if signatureGen != signatureIn {
		return false
	}
	return true
}

func (e WechatEntity) validateMsg(timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
	msgSignatureGen := e.makeMsgSignature(timestamp, nonce, msgEncrypt)
	if msgSignatureGen != msgSignatureIn {
		return false
	}
	return true
}
func (e WechatEntity) Decrypt(content string) (b *TextRequestBody, err error) {
	cipherData, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return
	}

	// AES Decrypt
	plainData, err := aesDecrypt(cipherData, e.aesKey)
	if err != nil {
		return
	}
	b, err = e.parseEncryptTextRequestBody([]byte(plainData))
	return
}

func (e WechatEntity) parseEncryptRequestBody(r *http.Request) *EncryptRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &EncryptRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func (e WechatEntity) parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func (e WechatEntity) value2CDATA(v string) CDATAText {
	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func (e WechatEntity) makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = e.value2CDATA(fromUserName)
	textResponseBody.ToUserName = e.value2CDATA(toUserName)
	textResponseBody.MsgType = e.value2CDATA("text")
	textResponseBody.Content = e.value2CDATA(content)
	textResponseBody.CreateTime = strconv.Itoa(int(time.Duration(time.Now().Unix())))
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func (e WechatEntity) makeEncryptResponseBody(fromUserName, toUserName, content, nonce, timestamp string) ([]byte, error) {
	encryptBody := &EncryptResponseBody{}

	encryptXmlData, _ := e.makeEncryptXmlData(fromUserName, toUserName, timestamp, content)
	encryptBody.Encrypt = e.value2CDATA(encryptXmlData)
	encryptBody.MsgSignature = e.value2CDATA(e.makeMsgSignature(timestamp, nonce, encryptXmlData))
	encryptBody.TimeStamp = timestamp
	encryptBody.Nonce = e.value2CDATA(nonce)

	return xml.Marshal(encryptBody)
}

func (e WechatEntity) makeEncryptXmlData(fromUserName, toUserName, timestamp, body string) (string, error) {
	/*textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = e.value2CDATA(fromUserName)
	textResponseBody.ToUserName = e.value2CDATA(toUserName)
	textResponseBody.MsgType = e.value2CDATA("text")
	textResponseBody.Content = e.value2CDATA(content)
	textResponseBody.CreateTime = timestamp
	body, err := xml.MarshalIndent(textResponseBody, " ", "  ")
	if err != nil {
		return "", errors.New("xml marshal error")
	}*/

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int32(len(body)))
	if err != nil {
		fmt.Println("Binary write err:", err)
	}
	bodyLength := buf.Bytes()

	randomBytes := []byte("abcdefghijklmnop")

	plainData := bytes.Join([][]byte{randomBytes, bodyLength, []byte(body), []byte(e.appID)}, nil)
	cipherData, err := aesEncrypt(plainData, e.aesKey)
	if err != nil {
		return "", errors.New("aesEncrypt error")
	}

	return base64.StdEncoding.EncodeToString(cipherData), nil
}

// PadLength calculates padding length, from github.com/vgorin/cryptogo
func PadLength(slice_length, blocksize int) (padlen int) {
	padlen = blocksize - slice_length%blocksize
	if padlen == 0 {
		padlen = blocksize
	}
	return padlen
}

//from github.com/vgorin/cryptogo
func PKCS7Pad(message []byte, blocksize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blocksize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blocksize < 1<<8 {
		// calculate padding length
		padlen := PadLength(len(message), blocksize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded = append(message, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}

func aesEncrypt(plainData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey)
	if len(plainData)%k != 0 {
		plainData = PKCS7Pad(plainData, k)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipherData := make([]byte, len(plainData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, plainData)

	return cipherData, nil
}

func aesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}

func (e WechatEntity) validateAppId(id []byte) bool {
	if string(id) == e.appID {
		return true
	}
	return false
}

func (e WechatEntity) parseEncryptTextRequestBody(plainText []byte) (*TextRequestBody, error) {
	//fmt.Println(string(plainText), len(plainText))

	// Read length
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	//fmt.Println("lenth:", length)
	//	fmt.Println(string(plainText[20 : 20+length]))

	// appID validation
	appIDstart := 20 + length
	id := plainText[appIDstart : int(appIDstart)+len(e.appID)]
	if !e.validateAppId(id) {
		return nil, errors.New("Appid is invalid")
	}

	textRequestBody := &TextRequestBody{}
	xml.Unmarshal(plainText[20:20+length], textRequestBody)
	return textRequestBody, nil
}

func (e WechatEntity) parseEncryptResponse(responseEncryptTextBody []byte) {
	textResponseBody := &EncryptResponseBody1{}
	xml.Unmarshal(responseEncryptTextBody, textResponseBody)

	if !e.validateMsg(textResponseBody.TimeStamp, textResponseBody.Nonce, textResponseBody.Encrypt, textResponseBody.MsgSignature) {
		fmt.Println("msg signature is invalid")
		return
	}

	cipherData, err := base64.StdEncoding.DecodeString(textResponseBody.Encrypt)
	if err != nil {
		log.Println("Wechat Service: Decode base64 error:", err)
		return
	}

	plainText, err := aesDecrypt(cipherData, e.aesKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(plainText))
}

func (e WechatEntity) procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signature := strings.Join(r.Form["signature"], "")
	encryptType := strings.Join(r.Form["encrypt_type"], "")
	msgSignature := strings.Join(r.Form["msg_signature"], "")

	fmt.Println("timestamp =", timestamp)
	fmt.Println("nonce =", nonce)
	fmt.Println("signature =", signature)
	fmt.Println("msgSignature =", msgSignature)

	if !e.validateUrl(timestamp, nonce, signature) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}

	if r.Method == "POST" {
		if encryptType == "aes" {
			log.Println("Wechat Service: in safe mode")
			encryptRequestBody := e.parseEncryptRequestBody(r)

			// Validate msg signature
			if !e.validateMsg(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
				log.Println("Wechat Service: msg_signature is invalid")
				return
			}
			log.Println("Wechat Service: msg_signature validation is ok!")

			// Decode base64
			cipherData, err := base64.StdEncoding.DecodeString(encryptRequestBody.Encrypt)
			if err != nil {
				log.Println("Wechat Service: Decode base64 error:", err)
				return
			}

			// AES Decrypt
			plainData, err := aesDecrypt(cipherData, e.aesKey)
			if err != nil {
				fmt.Println(err)
				return
			}

			textRequestBody, _ := e.parseEncryptTextRequestBody(plainData)
			fmt.Println(textRequestBody)
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)

			responseEncryptTextBody, _ := e.makeEncryptResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				"Hello, "+textRequestBody.FromUserName,
				nonce,
				timestamp)
			w.Header().Set("Content-Type", "text/xml")
			fmt.Println("\n", string(responseEncryptTextBody))
			fmt.Fprintf(w, string(responseEncryptTextBody))

			e.parseEncryptResponse(responseEncryptTextBody)
		} else if encryptType == "raw" {
			log.Println("Wechat Service: in raw mode")
		}
	}
}

/*
func main() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}
*/
