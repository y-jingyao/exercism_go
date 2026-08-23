package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd].*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var result []string
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			uname := match[1]
			newLine := "[USR] " + uname + " " + line
			result = append(result, newLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}
