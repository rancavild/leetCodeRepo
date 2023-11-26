package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func LongestSubstring(s string) int {
	seenChar := map[string]int{}

	if len(s) <= 1 {
		return len(s)
	}
	left, right, longest := 0, 0, 0

	for left < len(s) && right < len(s) {
		curr := s[right]
		prev, exists := seenChar[string(curr)]
		if exists {
			left = max(left, prev+1)
		}
		seenChar[string(curr)] = right
		longest = max(longest, right-left+1)
		right++
	}
	return longest
}
