package main

import "testing"

func TestIsFibCalc(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// True
		{"Three", args{3}, true},
		{"Five", args{5}, true},
		{"Eight", args{8}, true},
		{"Thirteen", args{13}, true},
		{"Twenty-one", args{21}, true},
		{"Thirty-four", args{34}, true},
		{"Fifty-five", args{55}, true},
		{"Eighty-nine", args{89}, true},
		//False
		{"Fifteen", args{15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFibCalc(tt.args.n); got != tt.want {
				t.Errorf("IsFibCalc(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
