package roman

import "strings"

var values = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func ToInteger(n string) int {
	value := 0

	for i := 0; i < len(n); i++ {
		candidate := ""

		if i+1 < len(n) {
			candidate = string(n[i]) + string(n[i+1])
		}

		v, ok := values[candidate]
		if ok {
			value += v
			i++
			continue
		}

		candidate = string(n[i])
		v, ok = values[candidate]
		if ok {
			value += v
		}
	}

	return value
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
