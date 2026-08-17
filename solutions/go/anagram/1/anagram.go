package anagram

import "strings"

func getFreq(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	return freq
}

func Detect(subject string, candidates []string) []string {
	//panic("Please implement the Detect function")
	var result []string
	freq := getFreq(strings.ToLower(subject))
	for _, candidate := range candidates {
		if len(candidate) != len(subject) {
			continue
		} else if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		} else {
			tFreq := getFreq(strings.ToLower(candidate))
			equal := true
			for k, v := range freq {
				if tFreq[k] != v {
					equal = false
					break
				}
			}
			if equal {
				result = append(result, candidate)
			}
		}
	}
	return result
}
