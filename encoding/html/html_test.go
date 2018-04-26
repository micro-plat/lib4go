package html

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	htmlEnInput := "<div>"
	htmlEnExcept := "&lt;div&gt;"
	htmlEnActual := Encode(htmlEnInput)
	if !strings.EqualFold(htmlEnExcept, htmlEnActual) {
		t.Errorf("Encode fail %s to %s", htmlEnInput, htmlEnActual)
	}

	htmlEnInput = "~!@! ~"
	htmlEnExcept = "~!@! ~"
	htmlEnActual = Encode(htmlEnInput)
	if !strings.EqualFold(htmlEnExcept, htmlEnActual) {
		t.Errorf("Encode fail %s to %s", htmlEnInput, htmlEnActual)
	}

	htmlEnInput = ""
	htmlEnExcept = ""
	htmlEnActual = Encode(htmlEnInput)
	if !strings.EqualFold(htmlEnExcept, htmlEnActual) {
		t.Errorf("Encode fail %s to %s", htmlEnInput, htmlEnActual)
	}

	htmlDeInput := "&lt;div&gt;"
	htmlDeExcept := "<div>"
	htmlDeActual := Decode(htmlDeInput)
	if !strings.EqualFold(htmlDeExcept, htmlDeActual) {
		t.Errorf("Decode fail %s to %s", htmlDeExcept, htmlDeActual)
	}

	htmlDeInput = "!@#!# !"
	htmlDeExcept = "!@#!# !"
	htmlDeActual = Decode(htmlDeInput)
	if !strings.EqualFold(htmlDeExcept, htmlDeActual) {
		t.Errorf("Decode fail %s to %s", htmlDeExcept, htmlDeActual)
	}

	htmlDeInput = ""
	htmlDeExcept = ""
	htmlDeActual = Decode(htmlDeInput)
	if !strings.EqualFold(htmlDeExcept, htmlDeActual) {
		t.Errorf("Decode fail %s to %s", htmlDeExcept, htmlDeActual)
	}
}
