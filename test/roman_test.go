package test

import "roman-calc/roman"
import "testing"

func TestRomanToDecimal(t *testing.T) {
	var values = []struct {
		roman   string
		decimal int
	}{
		{
			roman:   "I",
			decimal: 1,
		},
		{
			roman:   "IV",
			decimal: 4,
		},
		{
			roman:   "IX",
			decimal: 9,
		},
		{
			roman:   "CMXCIX",
			decimal: 999,
		},
		{
			roman:   "MCMLXXXIV",
			decimal: 1984,
		},
	}

	for _, value := range values {
		if v := roman.ToInteger(value.roman); v != value.decimal {
			t.Fatalf("Wrong value for %s, expected %d, was %d", value.roman, value.decimal, v)
		}
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
