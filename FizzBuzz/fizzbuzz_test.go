package FizzBuzz

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		numbers int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				numbers: 15,
			},
			want: fizzBuzz(15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
