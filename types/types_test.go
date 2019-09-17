package types

import "testing"
import "github.com/micro-plat/lib4go/ut"

func TestCopy1(t *testing.T) {
	m := Copy(map[string]interface{}{}, "k", "v")
	if len(m) != 1 {
		t.Error("err1:", m)
	}
	for k, v := range m {
		if k != "k" || v != "v" {
			t.Error("err2:", m)
		}
	}
}
func TestCopy2(t *testing.T) {
	m := Copy(map[string]interface{}{
		"k1": "v1",
	}, "k", "v")
	if len(m) != 2 {
		t.Error("err1:", m)
	}
	for k, v := range m {
		switch {
		case k == "k1":
			if v != "v1" {
				t.Error("err2:", m)
			}

		case k == "k":
			if v != "v" {
				t.Error("err3:", m)
			}
		}
	}
}
func BenchmarkTest(b *testing.B) {
	ut.Expect(b, DecodeString("3", "2", "3", "3", "2", "4"), "2")

}

func TestDecode1(t *testing.T) {
	ut.Expect(t, DecodeString(1, 2, 3), "")
	ut.Expect(t, DecodeString("1", 2, 3), "1")
	ut.Expect(t, DecodeString(2, 2, 3), "3")
	ut.Expect(t, DecodeString(1, 2, 3, 4), "4")
	ut.Expect(t, DecodeString(3, 2, 3, 4), "4")
	ut.Expect(t, DecodeString(3, 2, 3, 3, 2, 4), "2")

}
func TestDecode2(t *testing.T) {
	ut.Expect(t, DecodeInt(1, 2, 3, 1), 1)
	ut.Expect(t, DecodeInt(2, 2, 3, 2), 3)
	ut.Expect(t, DecodeInt(1, 2, 3, 4), 4)
	ut.Expect(t, DecodeInt(3, 2, 3, 4), 4)
	ut.Expect(t, DecodeInt(3, 2, 3, 3, 2, 4), 2)

}
func TestDecode3(t *testing.T) {
	ut.Expect(t, DecodeInt(1, 2, "3", 1), 1)
	ut.Expect(t, DecodeInt(2, 2, "3", 3), 3)
	ut.Expect(t, DecodeInt(1, 2, "3", "4"), 4)
	ut.Expect(t, DecodeInt(3, 2, "3", "4"), 4)
	ut.Expect(t, DecodeInt(3, 2, "3", 3, "2", "4"), 2)
	ut.Expect(t, DecodeInt(3, 2, "3", "3", "2", "4"), 2)
}
func TestDecode4(t *testing.T) {
	ut.Expect(t, DecodeInt(1, 1, 0, 1), 0)
	ut.Expect(t, DecodeInt(0, 1, 0, 1), 1)
	ut.Expect(t, DecodeInt("0", "1", 0, 1), 1)
	ut.Expect(t, DecodeInt("1", "1", 0, 1), 0)

	ut.Expect(t, DecodeInt("0", 1, 0, 1), 1)
	ut.Expect(t, DecodeInt("1", 1, 0, 1), 0)
	ut.Expect(t, DecodeInt(float64(1), 1, 0, 1), 0)
}
