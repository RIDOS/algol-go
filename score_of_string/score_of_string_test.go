package scoreofstring

import "testing"

func Test_scoreOfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				s: "hello",
			},
			want: 13,
		},
		{
			name: "test 2",
			args: args{
				s: "zaz",
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfString(tt.args.s); got != tt.want {
				t.Errorf("scoreOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
