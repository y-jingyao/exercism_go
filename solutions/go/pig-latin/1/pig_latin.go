package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	var result []string
	for _, w := range words {
		result = append(result, translateWord(w))
	}
	return strings.Join(result, " ")
}

func translateWord(word string) string {
	reVowel := regexp.MustCompile(`^(xr|yt|[aeiou])`)
	if reVowel.MatchString(word) {
		return word + "ay"
	}

	reQu := regexp.MustCompile(`^([^aeiou]*qu)(.+)`)
	if m := reQu.FindStringSubmatch(word); m != nil {
		return m[2] + m[1] + "ay"
	}

	reY := regexp.MustCompile(`^([^aeiou]+)(y.+)`)
	if m := reY.FindStringSubmatch(word); m != nil {
		return m[2] + m[1] + "ay"
	}

	reConsonant := regexp.MustCompile(`^([^aeiou]+)(.+)`)
	if m := reConsonant.FindStringSubmatch(word); m != nil {
		return m[2] + m[1] + "ay"
	}

	return word + "ay"
}
