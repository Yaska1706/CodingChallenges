package JosephusProblem

import "testing"

func TestJosephus(t *testing.T) {
	type args struct {
		number int
		start  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "When there are 5 people and starting position is 2",
			args: args{
				number: 5,
				start:  2,
			},
			want: 3,
		},
		{
			name: "When there are 14 people and starting position is 2",
			args: args{
				number: 14,
				start:  2,
			},
			want: 13,
		},
		{
			name: "When there are 14 people and starting position is 5",
			args: args{
				number: 14,
				start:  5,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Josephus(tt.args.number, tt.args.start); got != tt.want {

				t.Errorf("Expected %v but instead got %v", got, tt.want)
			}
		})
	}
}
