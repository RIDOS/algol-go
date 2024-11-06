package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 1",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test 2",
			args: args{
				list1: &ListNode{},
				list2: &ListNode{},
			},
			want: &ListNode{Next: &ListNode{}},
		},
		{
			name: "test 3",
			args: args{
				list1: &ListNode{},
				list2: &ListNode{Val: 0},
			},
			want: &ListNode{Next: &ListNode{Val: 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				got.PrintL()
				tt.want.PrintL()
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (head *ListNode) PrintL() {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}
