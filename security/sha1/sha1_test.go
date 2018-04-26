package sha1

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	input := "hello"
	except := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
	// except := "AAF4C61DDCC5E8A2DABEDE0F3B482CD9AEA9434D"
	actual := Encrypt(input)
	if !strings.EqualFold(actual, except) {
		t.Errorf("Encrypt fail %s to %s", except, actual)
	}
}
