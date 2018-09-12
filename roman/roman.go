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

	var falldowns = []struct {
		roman string
		value int
	}{
		{
			roman: "M",
			value: 1000,
		},
		{
			roman: "CM",
			value: 900,
		},
		{
			roman: "D",
			value: 500,
		},
		{
			roman: "CD",
			value: 400,
		},
		{
			roman: "C",
			value: 100,
		},
		{
			roman: "XC",
			value: 90,
		},
		{
			roman: "L",
			value: 50,
		},
		{
			roman: "XL",
			value: 40,
		},
		{
			roman: "X",
			value: 10,
		},
		{
			roman: "IX",
			value: 9,
		},
		{
			roman: "V",
			value: 5,
		},
		{
			roman: "IV",
			value: 4,
		},
		{
			roman: "I",
			value: 1,
		},
	}

	for i := 0; i < len(falldowns); i++ {
		for n >= falldowns[i].value {
			b.WriteString(falldowns[i].roman)
			n -= falldowns[i].value
		}
	}

	return b.String()
}
