package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	test := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !Equal(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for index, num := range a {
		if num != b[index] {
			return false
		}
	}

	return true
}
