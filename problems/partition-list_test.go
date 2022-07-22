package problems

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Example 1",
			args{intSlice{1, 4, 3, 2, 5, 2}.asList(), 3},
			intSlice{1, 2, 2, 1, 3, 5}.asList(),
		},
		{
			"Example 2",
			args{intSlice{2, 1}.asList(), 2},
			intSlice{1, 2}.asList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got.asSlice(), tt.want.asSlice())
			}
		})
	}
}
