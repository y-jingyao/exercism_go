package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	rePrefix := regexp.MustCompile(`^What is (.+)\?$`)
	match := rePrefix.FindStringSubmatch(question)
	if match == nil {
		return 0, false
	}
	content := match[1]

	replacer := strings.NewReplacer(
		"plus", "+",
		"minus", "-",
		"multiplied by", "*",
		"divided by", "/",
	)
	expr := replacer.Replace(content)

	reParse := regexp.MustCompile(`(-?\d+)([+\-*/](-?\d+))*`)
	if !reParse.MatchString(expr) {
		return 0, false
	}

	tokens := regexp.MustCompile(`\s+`).Split(expr, -1)

	if len(tokens)%2 == 0 {
		return 0, false
	}

	val, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	for i := 1; i < len(tokens); i += 2 {
		op := tokens[i]
		numStr := tokens[i+1]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, false
		}
		switch op {
		case "+":
			val += num
		case "-":
			val -= num
		case "*":
			val *= num
		case "/":
			if num == 0 {
				return 0, false
			}
			val /= num
		default:
			return 0, false
		}
	}

	cleaned := regexp.MustCompile(`-?\d+|plus|minus|multiplied by|divided by`).ReplaceAllString(content, "")
	cleaned = strings.TrimSpace(cleaned)
	if cleaned != "" {
		return 0, false
	}

	return val, true
}
