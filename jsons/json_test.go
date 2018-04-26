package jsons

import (
	"strings"
	"testing"
)

func TestEscape(t *testing.T) {
	input := "\\u0026123\\u003c1\\u003e"
	except := "&123<1>"
	actual := Escape(input)
	if !strings.EqualFold(actual, except) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}
}
