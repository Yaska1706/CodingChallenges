package main

import (
	"reflect"
	"testing"
)

func Test_noobtwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:   []int{2, 11, 7, 9, -2},
				target: 9,
			},
			want: []int{-2, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noobtwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noobtwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
