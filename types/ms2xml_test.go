package types

import (
	"reflect"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

type p struct {
	Name string   `xml:"name"`
	Age  int      `xml:"age"`
	Item []string `xml:"item"`
}

type userInfo struct {
	Name   string            `xml:"name"`
	ID     int               `xml:"id"`
	His    map[string]string `xml:"his"`
	Parent p                 `xml:"parent"`
	Child  *p                `xml:"child"`
	Ages   []int             `xml:"age"`
}

func TestIsZero(t *testing.T) {

	var u = &userInfo{Name: "abc"}
	assert.Equal(t, true, !reflect.ValueOf(u.Parent).IsValid() || reflect.ValueOf(u.Parent).IsZero())

	var v string
	assert.Equal(t, true, !reflect.ValueOf(v).IsValid() || reflect.ValueOf(v).IsZero())

	var v1 string = ""
	assert.Equal(t, true, !reflect.ValueOf(v1).IsValid() || reflect.ValueOf(v1).IsZero())

	var v2 string = "1"
	assert.Equal(t, false, !reflect.ValueOf(v2).IsValid() || reflect.ValueOf(v2).IsZero())

	var b int
	assert.Equal(t, true, !reflect.ValueOf(b).IsValid() || reflect.ValueOf(b).IsZero())

	var c int = 0
	assert.Equal(t, true, !reflect.ValueOf(c).IsValid() || reflect.ValueOf(c).IsZero())

	var c1 int = 1
	assert.Equal(t, false, !reflect.ValueOf(c1).IsValid() || reflect.ValueOf(c1).IsZero())

	var d interface{}
	assert.Equal(t, true, !reflect.ValueOf(d).IsValid() || reflect.ValueOf(d).IsZero())

	var e *userInfo
	assert.Equal(t, true, !reflect.ValueOf(e).IsValid() || reflect.ValueOf(e).IsZero())

	var e1 *userInfo = &userInfo{}
	assert.Equal(t, false, !reflect.ValueOf(e1).IsValid() || reflect.ValueOf(e1).IsZero())

	var f userInfo
	assert.Equal(t, true, !reflect.ValueOf(f).IsValid() || reflect.ValueOf(f).IsZero())

	var g userInfo = userInfo{Name: "colin"}
	assert.Equal(t, false, !reflect.ValueOf(g).IsValid() || reflect.ValueOf(g).IsZero())

}

func TestSlice2XML(t *testing.T) {
	cases := []struct {
		name   string
		input  interface{}
		root   []string
		result string
	}{

		{name: "1. string数组转xml", input: []string{"a", "b"}, root: []string{"item"}, result: "<item>a</item><item>b</item>"},
		{name: "2. int数组转xml", input: []int{1, 2}, root: []string{"item"}, result: "<item>1</item><item>2</item>"},
		{name: "3. int数组转xml", input: []int{1, 2}, result: "<item>1</item><item>2</item>"},
		{name: "2. struct中数组转xml", input: p{Item: []string{"a", "b"}, Name: "colin"}, root: []string{"xml"}, result: "<xml><name>colin</name><item>a</item><item>b</item></xml>"},
	}
	for _, c := range cases {

		xml, err := Any2XML(c.input, c.root...)
		assert.Equal(t, nil, err, c.name)
		assert.Equal(t, c.result, xml, c.name)

	}
}

func TestMap2XML(t *testing.T) {
	cases := []struct {
		name   string
		input  interface{}
		root   []string
		result string
	}{
		{name: "1. 字符串转xml", input: map[string]interface{}{"name": "colin"}, result: "<name>colin</name>"},
		{name: "2. int转xml", input: map[string]interface{}{"age": 10}, result: "<age>10</age>"},
		{name: "3. float32转xml", input: map[string]interface{}{"age": float32(10.2)}, result: "<age>10.2</age>"},
		{name: "4. float64转xml", input: map[string]interface{}{"age": float64(10.2)}, result: "<age>10.2</age>"},
		// {name: "5. decimal转xml", input: map[string]interface{}{"age": NewDecimalFromFloat(10.2)}, result: "<age>10.2</age>"},
		{name: "6. struct转xml", input: map[string]interface{}{"p": p{Name: "colin0", Age: 40}}, result: "<p><name>colin0</name><age>40</age></p>"},
		{name: "7. map[string]string转xml", input: map[string]interface{}{"p": map[string]string{"name": "colin0"}}, result: "<p><name>colin0</name></p>"},
		{name: "8. map[string]interface{}转xml", input: map[string]interface{}{"p": map[string]interface{}{"name": "colin0"}}, result: "<p><name>colin0</name></p>"},

		{name: "11. 字符串转xml", input: map[string]string{"name": "colin"}, result: "<name>colin</name>"},
		{name: "12. int转xml", input: map[string]int{"age": 10}, result: "<age>10</age>"},
		{name: "13. float32转xml", input: map[string]float32{"age": float32(10.2)}, result: "<age>10.2</age>"},
		{name: "14. float64转xml", input: map[string]float64{"age": float64(10.2)}, result: "<age>10.2</age>"},
		{name: "15. []int转xml", input: map[string]interface{}{"age": []int{1, 2}}, result: "<age>1</age><age>2</age>"},
		// {name: "15. decimal转xml", input: map[string]Decimal{"age": NewDecimalFromFloat(10.2)}, result: "<age>10.2</age>"},
	}
	for _, c := range cases {
		xml, err := Any2XML(c.input, c.root...)
		assert.Equal(t, nil, err, c.name)
		assert.Equal(t, c.result, xml, c.name)
	}

}
func TestStruct2XML(t *testing.T) {
	cases := []struct {
		name   string
		input  interface{}
		root   []string
		result string
	}{
		{name: "20:struct转xml", input: &p{Name: "colin"}, result: "<name>colin</name>"},
		{name: "21:struct转xml", input: &p{Age: 10}, result: "<age>10</age>"},

		{name: "31:struct转xml", input: &userInfo{Name: "colin", ID: 100}, result: "<name>colin</name><id>100</id>"},
		{name: "32:struct转xml", input: &userInfo{Name: "colin", ID: 100}, root: []string{"xml"}, result: "<xml><name>colin</name><id>100</id></xml>"},
		{name: "33:struct转xml", input: &userInfo{Name: "colin", ID: 100, His: map[string]string{"age": "100"}}, root: []string{"xml"}, result: "<xml><name>colin</name><id>100</id><his><age>100</age></his></xml>"},
		{name: "34:struct转xml", input: &userInfo{Name: "colin", ID: 100, His: map[string]string{"age": "100"}, Parent: p{Name: "colin0", Age: 30}}, root: []string{"xml"}, result: "<xml><name>colin</name><id>100</id><his><age>100</age></his><parent><name>colin0</name><age>30</age></parent></xml>"},
		{name: "35:struct转xml", input: &userInfo{Name: "colin", ID: 100, His: map[string]string{"age": "100"}, Child: &p{Name: "colin0", Age: 30}}, root: []string{"xml"}, result: "<xml><name>colin</name><id>100</id><his><age>100</age></his><child><name>colin0</name><age>30</age></child></xml>"},
		{name: "31:struct转xml", input: &userInfo{Name: "colin", ID: 100, Ages: []int{1, 2}}, result: "<name>colin</name><id>100</id><age>1</age><age>2</age>"},
	}
	for _, c := range cases {
		xml, err := Any2XML(c.input, c.root...)
		assert.Equal(t, nil, err, c.name)
		assert.Equal(t, c.result, xml, c.name)
	}

}
