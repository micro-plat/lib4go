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
}
