package test

import "roman-calc/roman"
import "testing"

func TestRomanSingle(t *testing.T) {
	if v := roman.ToInteger("I"); v != 1 {
		t.Fatalf("Wrong value for I, expected %d, was %d", 1, v)
	}
}

func TestDecimalToRoman(t *testing.T) {
	var values = []struct {
		roman   string
		decimal int
	}{
		{
			roman:   "MMXVIII",
			decimal: 2018,
		},
		{
			roman:   "MCMLXXXVII",
			decimal: 1987,
		},
		{
			roman:   "CCXXVI",
			decimal: 226,
		},
		{
			roman:   "CM",
			decimal: 900,
		},
		{
			roman:   "MDCCXII",
			decimal: 1712,
		},
		{
			roman:   "CMXCVIII",
			decimal: 998,
		},
		{
			roman:   "LXIV",
			decimal: 64,
		},
	}

	for _, value := range values {
		if v := roman.FromInteger(value.decimal); v != value.roman {
			t.Fatalf("Wrong value for %d, expected %s, was %s", value.decimal, value.roman, v)
		}
	}
}
