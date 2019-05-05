package unicode

import (
	"fmt"
	"strings"
	"testing"
)

func TestV(t *testing.T) {
	fmt.Println(Encode("政企分公司测试"))
}
func TestEncode(t *testing.T) {
	input := "你好"
	except := "\\u4f60\\u597d"
	actual := Encode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Encode fail %s to %s", except, actual)
	}

	input = "hello world"
	except = "hello world"
	actual = Encode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Encode fail %s to %s", except, actual)
	}

	input = "!@#!"
	except = "!@#!"
	actual = Encode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Encode fail %s to %s", except, actual)
	}
}

func TestDecode(t *testing.T) {
	input := "\\u4f60\\u597d"
	except := "你好"
	actual := Decode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Decode fail %s to %s", except, actual)
	}

	input = "\\u0068\\u0065\\u006c\\u006c\\u006f\\u0020\u0077\\u006f\\u0072\\u006c\\u0064"
	except = "hello world"
	actual = Decode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Decode fail %s to %s", except, actual)
	}

	input = "!@#!"
	except = "!@#!"
	actual = Decode(input)
	if !strings.EqualFold(except, actual) {
		t.Errorf("Decode fail %s to %s", except, actual)
	}
}
