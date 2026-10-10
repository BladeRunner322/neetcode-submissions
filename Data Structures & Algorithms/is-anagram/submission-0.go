func isAnagram(s string, t string) bool {
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, ch := range s {
		count1[ch]++
	}
	for _, ch := range t {
		count2[ch]++
	}

	if len(count1) != len(count2) {
		return false
	}
	for ch, n := range count1 {
		if count2[ch] != n {
			return false
		}
	}
	return true
}
