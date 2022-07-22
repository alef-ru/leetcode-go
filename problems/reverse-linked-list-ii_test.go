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
			args{intSlice{1, 2, 3, 4, 5}.asList(), 2, 4},
			intSlice{1, 4, 3, 2, 5}.asList(),
		}, {
			"Whole list reversion",
			args{intSlice{1, 2, 3, 4, 5}.asList(), 1, 5},
			intSlice{5, 4, 3, 2, 1}.asList(),
		}, {
			"One element",
			args{intSlice{42}.asList(), 1, 1},
			intSlice{42}.asList(),
		},
		{
			"Two elements",
			args{intSlice{4, 2}.asList(), 1, 2},
			intSlice{2, 4}.asList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
