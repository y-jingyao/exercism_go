package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	runes := []rune(log)
	for _, r := range runes {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	//panic("Please implement the Replace() function")
	runes := []rune(log)
	for i := 0; i < len(runes); i++ {
		if runes[i] == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	//panic("Please implement the WithinLimit() function")
	runes := []rune(log)
	return len(runes) <= limit
}
