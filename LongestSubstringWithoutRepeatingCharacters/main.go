package main

import "math"

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	start, max := 0, 0
	dict := make(map[uint8]int)
	for end := 0; end < length; end++ {
		idx, isDuplicate := dict[s[end]]
		if isDuplicate {
			max = int(math.Max(float64(end-start), float64(max)))
			for i := start; i <= idx; i++ {
				delete(dict, s[i])
			}
			start = idx + 1
		}
		dict[s[end]] = end
	}
	return int(math.Max(float64(max), float64(length-start)))
}
