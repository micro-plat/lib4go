package http

import (
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	client := NewHTTPClient()
	method := ""
	url := "http://www.baidu.com"
	httpClientRequest := client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}

	method = "test"
	url = "http://www.baidu.com"
	httpClientRequest = client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}

	method = ""
	url = "http://www.baidsdf!#2u.com"
	httpClientRequest = client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}
}

func TestSet(t *testing.T) {
	client := NewHTTPClient()
	method := ""
	url := "http://www.baidu.com"
	httpClientRequest := client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}
	httpClientRequest.SetData("test=test")
	httpClientRequest.SetHeader("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0")

	// // 测试request
	// content, status, err := httpClientRequest.Request()
	// if err != nil {

	// }
}

func TestHTTPClientRequest(t *testing.T) {
	client := NewHTTPClient()
	method := ""
	url := "http://www.baidu.com"
	httpClientRequest := client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}
	content, status, err := httpClientRequest.Request()
	if err != nil {
		t.Errorf("Request fail %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("Request fail status : %d", status)
	}
	if content == "" {
		t.Error("Request fail whit not content")
	}

	client = NewHTTPClient()
	method = ""
	url = "http://www.google.com"
	httpClientRequest = client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}
	content, status, err = httpClientRequest.Request()
	if err == nil {
		t.Error("test fail")
	}
	if status != 0 {
		t.Errorf("Request fail status : %d", status)
	}
	if content != "" {
		t.Error("Request fail whit not content")
	}

	client = NewHTTPClient()
	method = ""
	url = "http://192.168.0.100"
	httpClientRequest = client.NewRequest(method, url, "UTF-8")
	if httpClientRequest == nil {
		t.Error("Create Request Client fail")
	}
	content, status, err = httpClientRequest.Request()
	if err == nil {
		t.Errorf("Request fail %v", err)
	}
	if status != 0 {
		t.Errorf("Request fail status : %d", status)
	}
	if content != "" {
		t.Error("Request fail")
	}
}
