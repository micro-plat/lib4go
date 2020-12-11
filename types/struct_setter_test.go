package types

import (
	"testing"
	"time"

	"github.com/micro-plat/lib4go/assert"
)

type HD struct {
	Name string `json:"name"`
}
type st struct {
	ID          int                      `json:"id"`
	Name        string                   `json:"name"`
	Age         float64                  `json:"age"`
	Height      Decimal                  `json:"height"`
	Date        time.Time                `json:"date" time_format:"20060102150405"`
	Children    []string                 `json:"children"`
	Works       map[string]interface{}   `json:"work"`
	Clothes     map[string]string        `json:"clothes"`
	HD          *HD                      `json:"hd"`
	SubList1    []*HD                    `json:"sublist1"`
	SubList2    []HD                     `json:"sublist2"`
	MapList1    []HD                     `json:"maplist1"`
	MapList2    []*HD                    `json:"maplist2"`
	MapList3    []map[string]interface{} `json:"maplist3"`
	IntList     []int                    `json:"intlist"`
	Float32List []float32                `json:"floatList"`
	Float64List []float64                `json:"float64List"`
	StrList     []string                 `json:"strList"`
}

func TestToStruct(t *testing.T) {
	var input = map[string]interface{}{
		"id":          "10000",
		"name":        "colin",
		"date":        "20201204114231",
		"age":         "10.2",
		"height":      "170.3",
		"children":    []string{"a", "b"},
		"work":        map[string]interface{}{"id": 100},
		"clothes":     map[string]string{"name": "colin"},
		"hd":          &HD{Name: "colin"},
		"sublist1":    []*HD{&HD{Name: "colin1"}, &HD{Name: "colin2"}},
		"sublist2":    []HD{HD{Name: "colin3"}, HD{Name: "colin4"}},
		"maplist1":    []map[string]interface{}{map[string]interface{}{"name": "colin5"}},
		"maplist2":    []map[string]interface{}{map[string]interface{}{"name": "colin6"}},
		"maplist3":    []map[string]interface{}{map[string]interface{}{"name": "colin7"}},
		"intlist":     []string{"1", "2"},
		"floatList":   []string{"1", "2"},
		"float64List": []string{"1", "2"},
		"strList":     "abced",
	}

	var s = new(st)
	err := Map2Struct(s, input, "json")
	assert.Equal(t, nil, err)
	assert.Equal(t, 10000, s.ID)
	assert.Equal(t, 10.2, s.Age)
	assert.Equal(t, input["name"], s.Name)
	assert.Equal(t, NewDecimalFromFloat(170.3), s.Height)
	assert.Equal(t, input["date"], s.Date.Format("20060102150405"))
	assert.Equal(t, []string{"a", "b"}, s.Children)
	assert.Equal(t, map[string]interface{}{"id": float64(100)}, s.Works)
	assert.Equal(t, map[string]string{"name": "colin"}, s.Clothes)
	assert.Equal(t, &HD{Name: "colin"}, s.HD)
	assert.Equal(t, "colin", s.HD.Name)
	assert.Equal(t, []*HD{&HD{Name: "colin1"}, &HD{Name: "colin2"}}, s.SubList1)
	assert.Equal(t, []HD{HD{Name: "colin3"}, HD{Name: "colin4"}}, s.SubList2)
	assert.Equal(t, []map[string]interface{}{map[string]interface{}{"name": "colin7"}}, s.MapList3)
	assert.Equal(t, []float32{1, 2}, s.Float32List)
	assert.Equal(t, []float64{1, 2}, s.Float64List)
	assert.Equal(t, []string{"abced"}, s.StrList)
}
