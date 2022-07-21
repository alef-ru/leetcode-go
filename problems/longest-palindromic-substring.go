package problems

//Runtime: 4 ms, faster than 95.33% of Go online submissions for Longest Palindromic Substring.
//Memory Usage: 2.7 MB, less than 79.39% of Go online submissions for Longest Palindromic Substring.

// submission start

func getPalindrom(s string, left, right int) string {
	if s[left] != s[right] {
		return ""
	}
	for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
		left--
		right++
	}
	return s[left : right+1]
}

func longestPalindrome(s string) string {
	//the longest palindrome
	tlp := ""
	l, r := 0, 0
	for l < len(s) && r < len(s) {
		p := getPalindrom(s, l, r)
		if len(p) > len(tlp) {
			tlp = p
		}
		if l == r {
			r++
		} else {
			l++
		}
	}
	return tlp
}

// submission end
