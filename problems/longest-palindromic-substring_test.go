package problems

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"babad", "bab"},
		{"abc", "a"},
		{"1abba23", "abba"},
		{"abcddd", "ddd"},
		{"abcdd", "dd"},
		{"abc1dd1", "1dd1"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := longestPalindrome(tt.arg); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
