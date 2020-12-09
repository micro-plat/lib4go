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
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Age      float64                `json:"age"`
	Height   Decimal                `json:"height"`
	Date     time.Time              `json:"date" time_format:"20060102150405"`
	Children []string               `json:"children"`
	Works    map[string]interface{} `json:"work"`
	Clothes  map[string]string      `json:"clothes"`
	HD       *HD                    `json:"hd"`
	SubList  []*HD                  `json:"sublist"`
}

func TestToStruct(t *testing.T) {
	var input XMap = map[string]interface{}{
		"id":       "10000",
		"name":     "colin",
		"date":     "20201204114231",
		"age":      "10.2",
		"height":   "170.3",
		"children": []string{"a", "b"},
		"work":     map[string]interface{}{"id": 100},
		"clothes":  map[string]string{"name": "colin"},
		"hd":       &HD{Name: "colin"},
		"sublist":  []*HD{&HD{Name: "colin1"}, &HD{Name: "colin2"}},
	}

	var s = new(st)
	err := input.ToAnyStruct(&s)
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
	assert.Equal(t, []*HD{&HD{Name: "colin1"}, &HD{Name: "colin2"}}, s.SubList)

}
