package unicode

import (
	"fmt"
	"strings"
	"testing"
)

func TestV(t *testing.T) {
	fmt.Println(Decode("\u8BF7\u6C42\u53C2\u6570\u4E0D\u80FD\u4E3A你好\u7A7A你好"))
	t.Error("abc")
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
