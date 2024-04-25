package sprint

func AreAnagrams(str1, str2 string) bool {
	str1 = removeSpacesAndToLower(str1)
	str2 = removeSpacesAndToLower(str2)

	if len(str1) != len(str2) {
		return false
	}
	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	countFrequency(str1, freq1)
	countFrequency(str2, freq2)
	return compareFrequencyMaps(freq1, freq2)
}

func removeSpacesAndToLower(s string) string {
	var result []rune
	for _, char := range s {
		if char != ' ' {
			// Convert to lowercase
			if 'A' <= char && char <= 'Z' {
				char += 32
			}
			result = append(result, char)
		}
	}
	return string(result)
}

func countFrequency(s string, freq map[rune]int) {
	for _, char := range s {
		freq[char]++
	}
}

func compareFrequencyMaps(freq1, freq2 map[rune]int) bool {
	if len(freq1) != len(freq2) {
		return false
	}
	for char, count1 := range freq1 {
		count2, ok := freq2[char]
		if !ok || count1 != count2 {
			return false
		}
	}
	return true
}
