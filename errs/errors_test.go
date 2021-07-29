package errs

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestT(t *testing.T) {
	ErrOrderTimeout := errors.New("ORDER_TIME_OUT")
	err := NewError(http.StatusNoContent, fmt.Errorf("订单(%s)已超时%w", "", ErrOrderTimeout))
	assert.Equal(t, true, errors.Is(err, ErrOrderTimeout))

	nerr := NewError(904, "不存在")

	err0 := errors.New("ERR")
	err2 := New("油站%w%v", nerr, err0)
	assert.Equal(t, GetCode(err2), 904)
	assert.Equal(t, true, err2.Is(nerr))
	assert.Equal(t, "油站不存在ERR", err2.GetError().Error())

}
func TestR(t *testing.T) {
	var v error = NewResult(200, map[string]interface{}{"id": 100})
	_, ok := v.(IResult)
	assert.Equal(t, true, ok)
}
