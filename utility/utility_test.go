package utility

import (
	"testing"

	"github.com/qxnw/lib4go/ut"
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
	ut.Expect(t, err, nil)
	ut.Expect(t, len(m), 1)
	ut.Expect(t, m["a"], "a")
}
func TestMapWithQuery2(t *testing.T) {
	m, err := GetMapWithQuery("a=a,b")
	ut.Expect(t, err, nil)
	ut.Expect(t, len(m), 1)
	ut.Expect(t, m["a"], "a,b")
}
