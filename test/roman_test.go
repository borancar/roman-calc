package test

import "roman-calc/roman"
import "testing"

func TestRomanSingle(t *testing.T) {
	if v := roman.ToInteger("I"); v != 1 {
		t.Fatalf("Wrong value for I, expected %d, was %d", 1, v)
	}
}

func TestYears(t *testing.T) {
	if v := roman.FromInteger(2018); v != "MMXVIII" {
		t.Fatalf("Wrong value for 2018, expected %s, was %s", "MMXVIII", v)
	}
}
