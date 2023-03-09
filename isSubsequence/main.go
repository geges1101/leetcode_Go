package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	curr := 0
	for _, c := range t {
		if string(c) == string(s[curr]) {
			curr++
			if curr == len(s) {
				return true
			}
		}
	}
	return false
}
