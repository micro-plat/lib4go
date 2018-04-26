package html

import (
	"html"
)

// Encode html编码
func Encode(input string) string {
	return html.EscapeString(input)
}

// Decode html解码
func Decode(input string) string {
	return html.UnescapeString(input)
}
