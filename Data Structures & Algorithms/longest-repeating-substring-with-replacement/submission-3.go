func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}

	freq := [26]int{}
	l := 0
	longest := 0
	maxfreq := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]-'A']++
		maxfreq = max(maxfreq, freq[s[r]-'A'])

		if r-l+1 - maxfreq > k {
			freq[s[l]-'A']--
			l++
		}

		longest = max(longest, r-l+1)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
