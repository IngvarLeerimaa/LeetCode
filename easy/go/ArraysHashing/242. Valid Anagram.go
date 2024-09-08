package ArraysHashing

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	runeMap := make(map[rune]int)

	for _, v := range s {
		runeMap[v]++

	}
	for _, v := range t {
		runeMap[v]--
	}

	for _, v := range runeMap {
		if v != 0 {
			return false
		}
	}
	return true
}
