package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "test 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				s: "([])",
			},
			want: true,
		},
		{
			name: "test 5",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "test 6",
			args: args{
				s: "[",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
