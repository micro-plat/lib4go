package types

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	nmap := NewXMaps(1)
	m := NewXMap()
	m.SetValue("name", "123")
	nmap.Append(m)
	fmt.Println("nmap", nmap)
	if nmap.Len() != 0 {
		t.Errorf("长度错误:%d", nmap.Len())
	}
}
