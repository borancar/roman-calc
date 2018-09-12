package test

import "roman-calc/roman"
import "testing"

func TestRomanSingle(t *testing.T) {
	if v := roman.ToInteger("I"); v != 1 {
		t.Fatalf("Wrong value for I: %d", v)
	}
}
