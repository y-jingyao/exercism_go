package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	//panic("Please implement the WordCount function")
	freq := make(Frequency)
	re := regexp.MustCompile(`[a-zA-Z0-9]+(?:'[a-zA-Z0-9]+)*`)
	for _, word := range re.FindAllString(phrase, -1) {
		freq[strings.ToLower(word)]++
	}
	return freq
}
