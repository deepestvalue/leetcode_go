package problem3longestsubstring

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	left, right, maxLen := 0, 1, 0
	lookup := make(map[byte]int)
	lookup[s[left]] = left

	for right < len(s) {
		// if we found a repeated character, calculate the current max
		if _, ok := lookup[s[right]]; ok {
			maxLen = max(right-left, maxLen)

			// if current char is part of the current window, advance left past it
			if lookup[s[right]] >= left {
				left = lookup[s[right]] + 1
			}
		}

		lookup[s[right]] = right
		right++
	}

	return max(right-left, maxLen)

}
