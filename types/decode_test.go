package types

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestBool(t *testing.T) {
	f := DecodeBool(true, true, false)
	assert.Equal(t, false, f)

	f = DecodeBool(true, false, false)
	assert.Equal(t, true, f)

	f = DecodeBool(true, true, true, false)
	assert.Equal(t, true, f)

	f = DecodeBool(false, true, true, false)
	assert.Equal(t, false, f)
}
func TestInt(t *testing.T) {
	status, vstatus, fstatus := 0, 0, 204
	v := DecodeInt(DecodeInt(status, 0, vstatus), 0, fstatus)
	assert.Equal(t, v, fstatus)

	status, vstatus, fstatus = 200, 300, 400
	v = DecodeInt(DecodeInt(status, 0, vstatus), 0, fstatus)
	assert.Equal(t, v, status)

	status, vstatus, fstatus = 0, 300, 400
	v = DecodeInt(DecodeInt(status, 0, vstatus), 0, fstatus)
	assert.Equal(t, v, vstatus)

}
