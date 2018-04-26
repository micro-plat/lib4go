package des

import (
	"fmt"
	"testing"

	"github.com/qxnw/lib4go/ut"
)

func TestEncrypt(t *testing.T) {
	input := []byte("hello")
	key := "12345678"
	iv := []byte("87654321")
	crypted, err := EncryptBytes(input, key, iv, "cbc/pkcs5")
	ut.Expect(t, err, nil)

	o, err := DecryptBytes(crypted, key, iv, "cbc/pkcs5")
	ut.Expect(t, err, nil)
	ut.Expect(t, string(o), string(input))

	pk := []byte("encrypt:ecb/pkcs7")
	fmt.Println(len(pk))
}
