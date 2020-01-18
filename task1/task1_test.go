package task1

import "testing"

func TestFizzBuzzPinkFlamingo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Three", args{3}, "Flamingo"},
		{"Ten", args{10}, "Buzz"},
		{"Twelve", args{12}, "Fizz"},
		{"Fifteen", args{15}, "FizzBuzz"},
		{"Six-Thousand Sixty-five", args{6765}, "Pink Flamingo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzzPinkFlamingo(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzzPinkFlamingo(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
