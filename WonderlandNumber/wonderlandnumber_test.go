package WonderlandNumber

import (
	"reflect"
	"testing"
)

func Test_countDigits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "When number is less than 10",
			args: args{
				number: 9,
			},
			want: 1,
		},
		{
			name: "when number is 10 or more",
			args: args{
				number: 567,
			},
			want: 3,
		},
		{
			name: "when number is large",
			args: args{
				number: 567867,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.number); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storepresentDigits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "When digit count is less than 6",
			args: args{
				number: 4885,
			},
			want: []int{},
		},
		{
			name: "When digit count is greater than 6",
			args: args{
				number: 4885487,
			},
			want: []int{},
		},
		{
			name: "When digit count is 6",
			args: args{
				number: 488579,
			},
			want: []int{9, 7, 5, 8, 8, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storePresentDigits(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storepresentDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNumbersExist(t *testing.T) {
	type args struct {
		number int
		store  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Numbers all exist",
			args: args{
				number: 346759,
				store:  []int{6, 7, 9, 5, 3, 4},
			},
			want: true,
		},
		{
			name: "Numbers don't exist",
			args: args{
				number: 288211,
				store:  []int{6, 7, 9, 5, 3, 4},
			},
			want: false,
		},
		{
			name: "When not all Numbers exist",
			args: args{
				number: 672813,
				store:  []int{6, 7, 9, 5, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNumbersExist(tt.args.number, tt.args.store); got != tt.want {
				t.Errorf("checkNumbersExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wonderlandnumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "non-cyclic number",
			args: args{
				number: 234567,
			},
			want: false,
		},
		{
			name: "cyclic number",
			args: args{
				number: 142857,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderlandNumber(tt.args.number); got != tt.want {
				t.Errorf("wonderlandnumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatewonderlandnumber(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "cyclic number",
			args: args{
				size: 999999,
			},
			want: 142857,
		},
		{
			name: "out of range number",
			args: args{
				size: 1000,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateWonderlandNumber(tt.args.size); got != tt.want {
				t.Errorf("generatewonderlandnumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
