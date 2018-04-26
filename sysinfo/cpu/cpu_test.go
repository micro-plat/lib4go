package cpu

import "testing"

func Test(t *testing.T) {

	avaliable := GetInfo()
	if avaliable.Total == "" || avaliable.Idle == "" || avaliable.Used == "" {
		t.Error("test fail")
	}
}
