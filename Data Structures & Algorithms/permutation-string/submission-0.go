func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	f1 := [26]int{}
	f2 := [26]int{}

	for i := 0; i < len(s1); i++ {
		f1[s1[i]-'a']++
		f2[s2[i]-'a']++
	}

	for i := 0; i < len(s2) - len(s1); i++ {
		if permutation(f1, f2) {
			return true
		}

		f2[s2[i]-'a']--
		f2[s2[i+len(s1)]-'a']++
	}

	return permutation(f1, f2)
}

func permutation(f1, f2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}
