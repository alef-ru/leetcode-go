package problems

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Example from task description",
			args{createList([]int{1, 2, 3, 4, 5}), 2, 4},
			createList([]int{1, 4, 3, 2, 5}),
		}, {
			"Whole list reversion",
			args{createList([]int{1, 2, 3, 4, 5}), 1, 5},
			createList([]int{5, 4, 3, 2, 1}),
		}, {
			"One element",
			args{createList([]int{42}), 1, 1},
			createList([]int{42}),
		},
		{
			"Two elements",
			args{createList([]int{4, 2}), 1, 2},
			createList([]int{4, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", list2slice(got), list2slice(tt.want))
			}
		})
	}
}
