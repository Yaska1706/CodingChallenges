package FizzBuzzPlus

import (
	"reflect"
	"testing"
)

func Test_checkNumber(t *testing.T) {
	type args struct {
		number int
		n      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "when a number exists",
			args: args{
				number: 358,
				n:      3,
			},
			want: true,
		},
		{
			name: "number does not exist",
			args: args{
				number: 43,
				n:      2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNumber(tt.args.number, tt.args.n); got != tt.want {
				t.Errorf("checkNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzbuzzplus(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "when numbers are 1 to 5",
			args: args{
				number: 5,
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "when numbers are 1 to 15",
			args: args{
				number: 15,
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "6", "7", "8", "9", "10", "11", "12", "13", "14", "Buzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzzplus(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzbuzzplus() = %v, want %v", got, tt.want)
			}
		})
	}
}
