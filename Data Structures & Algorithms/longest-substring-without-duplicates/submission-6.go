func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	cache := make(map[byte]int)
	cache[s[0]] = 0
	l := 0
	r := 1
	longest := 1

	for l < len(s) && r < len(s) && l <= r {
		ch := s[r]

		if idx, present := cache[ch]; present {
			for i := l; i < idx + 1; i++ {
				delete(cache, s[i])
			}
			l = idx+1
		}
		
		cache[ch] = r
		longest = max(longest, r-l+1)
		r++
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
