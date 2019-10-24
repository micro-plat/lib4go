package transform

import (
	"testing"

	"github.com/micro-plat/lib4go/ut"
)

func TestParse1(t *testing.T) {

	check(t, Parse("@a=100"), []Expression{Expression{Left: "@a", Symbol: SymbolEqual, Right: "100"}})
	check(t, Parse("aes=100"), []Expression{Expression{Left: "aes", Symbol: SymbolEqual, Right: "100"}})
	check(t, Parse("@a>100"), []Expression{Expression{Left: "@a", Symbol: SymbolMore, Right: "100"}})
	check(t, Parse("@a!=100"), []Expression{Expression{Left: "@a", Symbol: SymbolNotEqual, Right: "100"}})
	check(t, Parse("@a<100"), []Expression{Expression{Left: "@a", Symbol: SymbolLess, Right: "100"}})
	check(t, Parse("@a=100&"), []Expression{Expression{Left: "@a", Symbol: SymbolEqual, Right: "100"}})
	check(t, Parse("@a=100&b!=200"), []Expression{Expression{Left: "@a", Symbol: SymbolEqual, Right: "100"}, Expression{Left: "b", Symbol: SymbolNotEqual, Right: "200"}})
	check(t, Parse("@a=100&b>=200"), []Expression{Expression{Left: "@a", Symbol: SymbolEqual, Right: "100"}, Expression{Left: "b", Symbol: SymbolMoreOrEqual, Right: "200"}})
	check(t, Parse("@a=100&b<=200"), []Expression{Expression{Left: "@a", Symbol: SymbolEqual, Right: "100"}, Expression{Left: "b", Symbol: SymbolLessOrEqual, Right: "200"}})

}

func check(t *testing.T, got []Expression, expected []Expression) {
	ut.Expect(t, len(got), len(expected))
	for i, v := range got {
		ut.Expect(t, v.Left, expected[i].Left)
		ut.Expect(t, v.Right, expected[i].Right)
		ut.Expect(t, v.Symbol, expected[i].Symbol)
	}
}
