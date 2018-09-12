package roman

import "strings"

var values = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ToInteger(n string) int {
	return values[string(n[0])]
}

func FromInteger(n int) string {
	var b strings.Builder

	for n >= 1000 {
		b.WriteString("M")
		n -= 1000
	}

	if n >= 900 {
		b.WriteString("CM")
		n -= 900
	}

	if n >= 500 {
		b.WriteString("D")
		n -= 500
	}

	if n >= 400 {
		b.WriteString("CD")
		n -= 400
	}

	for n >= 100 {
		b.WriteString("C")
		n -= 100
	}

	if n >= 90 {
		b.WriteString("XC")
		n -= 90
	}

	if n >= 50 {
		b.WriteString("L")
		n -= 50
	}

	if n >= 40 {
		b.WriteString("XL")
		n -= 40
	}

	for n >= 10 {
		b.WriteString("X")
		n -= 10
	}

	if n >= 9 {
		b.WriteString("IX")
		n -= 9
	}

	if n >= 5 {
		b.WriteString("V")
		n -= 5
	}

	for n >= 1 {
		b.WriteString("I")
		n -= 1
	}

	return b.String()
}
