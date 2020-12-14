package utility

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

// TestGetGUID 测试生成的Guid是否重复
func TestGetGUID(t *testing.T) {
	totalAccount := 10000 * 1000
	data := map[string]int{}

	for i := 0; i < totalAccount; i++ {
		key := GetGUID()
		data[key] = i
	}

	if len(data) != totalAccount {
		t.Errorf("test fail, totalAccount:%d, actual:%d", totalAccount, len(data))
	}
}
func TestMapWithQuery1(t *testing.T) {
	m, err := GetMapWithQuery("a=a")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(m), 1)
	assert.Equal(t, m["a"], "a")
}
func TestMapWithQuery2(t *testing.T) {
	m, err := GetMapWithQuery("a=a,b")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(m), 1)
	assert.Equal(t, m["a"], "a,b")
}
