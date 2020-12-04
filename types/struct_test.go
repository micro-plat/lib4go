package types

import (
	"testing"
	"time"

	"github.com/micro-plat/lib4go/assert"
)

type st struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Age    float64   `json:"age"`
	Height Decimal   `json:"height"`
	Date   time.Time `json:"date" time_format:"20060102150405"`
}

func TestToStruct(t *testing.T) {
	input := map[string]interface{}{
		"id":     "10000",
		"name":   "colin",
		"date":   "20201204114231",
		"age":    "10.2",
		"height": "170.3",
	}

	var s = new(st)
	err := Any2Struct(&s, input)
	assert.Equal(t, nil, err)
	assert.Equal(t, 10000, s.ID)
	assert.Equal(t, 10.2, s.Age)
	assert.Equal(t, input["name"], s.Name)
	assert.Equal(t, NewDecimalFromFloat(170.3), s.Height)
	assert.Equal(t, input["date"], s.Date.Format("20060102150405"))
}
