package microblog

func Truncate(phrase string) string {
	//panic("Please implement the Truncate function")
    runes := []rune(phrase)
    if len(runes) > 5 {
    return string(runes[:5])
    }
    return string(runes)
}
