package types

import (
	"fmt"
	"testing"
)

type Input struct {
	Name  string `m2s:"name"`
	Value int    `m2s:"value"`
}

func TestT(t *testing.T) {
	i := NewXMapByMap(
		map[string]interface{}{
			"name":  "abcef",
			"value": 100,
		})

	var input Input
	err := i.ToStruct(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", input)

	t.Error("abc")

}
func TestAppend(t *testing.T) {
	i := NewXMaps()
	i.Append(map[string]interface{}{
		"name":  "abcef",
		"value": 100,
	})

	var input []*Input
	err := i.ToStructs(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("1. %+v\n", i)
	for _, v := range input {
		fmt.Printf("2.%+v\n", v)
	}

	t.Error("abc")

}
