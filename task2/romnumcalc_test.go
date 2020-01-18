package task2

import (
	"testing"
)

func TestRomNumeral(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"24 + 11", args{"XXIV + XI"}, "XXXV"},
		{"24 + 12", args{"XXIV+XII"}, "XXXVI"},
		{"24 ^ 2", args{"XXIV^II"}, "DLXXVI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumeralCalculate(tt.args.in0); got != tt.want {
				t.Errorf("RomNumeral(%v) = %v, want %v", tt.args.in0, got, tt.want)
			}
		})
	}
}

func TestRomanNumeralCalculate(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumeralCalculate(tt.args.in); got != tt.want {
				t.Errorf("RomanNumeralCalculate(%v) = %v, want %v", tt.args.in, got, tt.want)
			}
		})
	}
}
