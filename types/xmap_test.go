package types

import (
	"fmt"
	"testing"
)

type Input struct {
	Name  string
	Value *Input
}

func TestAppend(t *testing.T) {
	i := map[string]interface{}{
		"order": 1234567,
		"items": map[string]interface{}{
			"pid":   100,
			"price": []string{"100.2", "99.8"},
			"children": map[string]string{
				"path": "/",
			},
		},
	}
	m := NewXMap()
	m.Cascade(NewXMapByMap(i))
	fmt.Println(m)

	fmt.Println(GetCascade("name", "colin"))
	fmt.Println(GetCascade("age", 100))
	t.Error("err")
}
