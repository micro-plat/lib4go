package http

import (
	"net/http"
	"testing"
)

func TestNewHTTPClientCert(t *testing.T) {
	certFile := "/home/champly/openssl/cert.pem"
	keyFile := "/home/champly/openssl/key.pem"
	caFile := "/home/champly/openssl/ca.pem"
	client, err := NewHTTPClientCert(certFile, keyFile, caFile)
	if err != nil {
		t.Error(err)
	}
	if client == nil {
		t.Error("test fail")
	}

	certFile = "/home/champly/openssl/err_cert.pem"
	keyFile = "/home/champly/openssl/err_key.pem"
	caFile = "/home/champly/openssl/err_ca.pem"
	_, err = NewHTTPClientCert(certFile, keyFile, caFile)
	if err == nil {
		t.Error("test fail")
	}
}

func TestGet(t *testing.T) {
	client := NewHTTPClient()
	content, status, err := client.Get("http://www.baidu.com", nil...)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Get error status:%d", status)
	}
	if content == "" {
		t.Error("Get error with not content")
	}

	content, status, err = client.Get("http://www.google.com", nil...)
	if err == nil {
		t.Errorf("Get error: %v", err)
	}
	if status != 0 {
		t.Errorf("Get error status:%d", status)
	}
	if content != "" {
		t.Error("Get error with not content")
	}

	content, status, err = client.Get("https://github.com/testhttp", nil...)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	if status != http.StatusNotFound {
		t.Errorf("Get error status:%d", status)
	}
	if content == "" {
		t.Error("Get error with not content")
	}
}

func TestPost(t *testing.T) {
	client := NewHTTPClient()
	content, status, err := client.Post("http://www.baidu.com", "name=bob", "UTF-8")
	if err != nil {
		t.Errorf("Post error: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Post error status:%d", status)
	}
	if content == "" {
		t.Error("Post error with not content")
	}

	content, status, err = client.Post("https://github.com/testhttp", "name=bob", "UTF-8")
	if err != nil {
		t.Errorf("Post error: %v", err)
	}
	if status != http.StatusNotFound {
		t.Errorf("Post error status:%d", status)
	}
	if content == "" {
		t.Error("Post error with not content")
	}

	content, status, err = client.Post("http://www.google.com", "name=bob", "UTF-8")
	if err == nil {
		t.Errorf("Post error: %v", err)
	}
	if status != 0 {
		t.Errorf("Post error status:%d", status)
	}
	if content != "" {
		t.Error("Post error with not content")
	}

	content, status, err = client.Post("https://github.com/testhttp", "", "UTF-8")
	if err != nil {
		t.Errorf("Post error: %v", err)
	}
	if status != http.StatusNotFound {
		t.Errorf("Post error status:%d", status)
	}
	if content == "" {
		t.Error("Post error with not content")
	}
}

func TestRequest(t *testing.T) {
	client := NewHTTPClient()
	url := "https://github.com"
	method := ""
	params := ""
	charset := "UTF-8"
	header := make(map[string]string)
	content, status, err := client.Request(method, url, params, charset, header)
	if err != nil {
		t.Errorf("Request error: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Request error status:%d", status)
	}
	if content == "" {
		t.Error("Request error with not content")
	}

	url = "http://www.google.com"
	method = ""
	params = ""
	charset = "UTF-8"
	header = make(map[string]string)
	content, status, err = client.Request(method, url, params, charset, header)
	if err == nil {
		t.Errorf("Request error: %v", err)
	}
	if status != 0 {
		t.Errorf("Request error status:%d", status)
	}
	if content != "" {
		t.Error("Request error with not content")
	}

	url = "https://github.com/testhttp"
	method = ""
	params = ""
	charset = "UTF-8"
	header = make(map[string]string)
	content, status, err = client.Request(method, url, params, charset, header)
	if err != nil {
		t.Error("test error")
	}
	if status != http.StatusNotFound {
		t.Errorf("Request error status:%d", status)
	}
	if content == "" {
		t.Error("Request error with not content")
	}

	url = "https://github.com"
	method = "qxnw/lib4go"
	params = ""
	charset = "UTF-8"
	header = make(map[string]string)
	content, status, err = client.Request(method, url, params, charset, header)
	if err == nil {
		t.Error("test error")
	}
	if status != 0 {
		t.Errorf("Request error status:%d", status)
	}
	if content != "" {
		t.Error("Request error with not content")
	}
}

func TestNewHTTPClientProxy(t *testing.T) {
	proxy := "192.168.0.100:6666"
	client := NewHTTPClientProxy(proxy)
	content, status, err := client.Get("http://www.baidu.com", "UTF-8")
	if err == nil {
		t.Errorf("Post error: %v", err)
	}

	proxy = "58.222.254.11:3128"
	client = NewHTTPClientProxy(proxy)
	content, status, err = client.Get("http://www.baidu.com", "UTF-8")
	if err != nil {
		t.Errorf("Post error: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Post error status:%d", status)
	}
	if content == "" {
		t.Error("Post error with not content")
	}
}

func TestDownload(t *testing.T) {
	client := NewHTTPClient()
	method := ""
	url := "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params := ""
	header := make(map[string]string)
	body, status, err := client.Download(method, url, params, header)
	if err != nil {
		t.Errorf("Download fail %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Download fail status:%d", status)
	}
	if body == nil {
		t.Error("Download fail")
	}

	method = ""
	url = "https://ss0.bdstatic.com/5aV1bCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params = ""
	header = make(map[string]string)
	body, status, err = client.Download(method, url, params, header)
	if err != nil {
		t.Errorf("Download fail %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Download fail status:%d", status)
	}
	if body == nil {
		t.Error("Download fail")
	}

	method = "download"
	url = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params = ""
	header = make(map[string]string)
	body, status, err = client.Download(method, url, params, header)
	if err != nil {
		t.Errorf("Download fail %v", err)
	}
	if status != 405 {
		t.Errorf("Download fail status:%d", status)
	}
	if body == nil {
		t.Error("Download fail")
	}
}

func TestSave(t *testing.T) {
	client := NewHTTPClient()
	method := ""
	url := "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params := ""
	header := make(map[string]string)
	path := "/home/champly/pictue.png"
	status, err := client.Save(method, url, params, header, path)
	if err != nil {
		t.Errorf("Save fail %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Save fail status:%d", status)
	}

	method = ""
	url = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params = ""
	header = make(map[string]string)
	path = "/home/champly"
	status, err = client.Save(method, url, params, header, path)
	if err == nil {
		t.Errorf("Save fail %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Save fail status:%d", status)
	}

	method = "save"
	url = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
	params = ""
	header = make(map[string]string)
	path = "/home/champly/pictue.png"
	status, err = client.Save(method, url, params, header, path)
	if err != nil {
		t.Errorf("Save fail %v", err)
	}
	if status != 405 {
		t.Errorf("Save fail status:%d", status)
	}
}
