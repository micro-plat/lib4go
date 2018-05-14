package db

import (
	"testing"

	"github.com/micro-plat/lib4go/ut"
)

func TestT1(t *testing.T) {
	var q QueryRows
	ut.Expect(t, len(q), 0)
	q = nil
	ut.Expect(t, len(q), 0)
	ut.Expect(t, q.IsEmpty(), true)
}
