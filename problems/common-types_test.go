package problems

import "testing"

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want string
	}{
		{
			"e2e test: slice -> list -> slice -> string",
			intSlice{1, 2, 3, 4, 5}.asList(),
			"[1 2 3 4 5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.head.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
