package problems

// submission start

func isSubsequence(s, w string) bool {
	var si, wi int //string index, word index
	for si < len(s) {
		if w[wi] == s[si] {
			wi++
			if wi == len(w) {
				return true
			}
		}
		si++
	}
	return false
}

func numMatchingSubseq(s string, words []string) (count int) {
	for _, w := range words {
		if isSubsequence(s, w) {
			count++
		}
	}

	return count
}

// submission end
