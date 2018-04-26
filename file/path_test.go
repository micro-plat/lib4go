package file

import (
	"strings"
	"testing"
)

func TestExists(t *testing.T) {
	if !Exists("/home/champly") {
		t.Error("/home/champly exist test fail22.")

	}

	if Exists("/home/champly/adfada") {
		t.Error("/home/champly/adfada is not exist test fail")
	}

	if Exists("") {
		t.Error(" is not exist test fail")
	}

	if Exists("adsf") {
		t.Error("adsf is not exist test fail")
	}
}

func TestGetAbs(t *testing.T) {
	path := GetAbs("123")
	if strings.EqualFold(path, "") {
		t.Error("test fail")
	}

	path = GetAbs("../../")
	if !strings.EqualFold(path, "/home/champly/qxnw") {
		t.Error("test fail")
	}
}

func TestCreateFile(t *testing.T) {
	_, err := CreateFile("../../")
	if err == nil {
		t.Errorf("test fail")
	}

	_, err = CreateFile("1")
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	_, err = CreateFile("1")
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	_, err = CreateFile("")
	if err == nil {
		t.Error("test fail")
	}

	_, err = CreateFile("test.log")
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	_, err = CreateFile("/root/test")
	if err == nil {
		t.Error("test fail")
	}
}
