package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	//panic("Please implement the Transform function")
	result := make(map[string]int)
	for key, value := range in {
		for _, s := range value {
			result[strings.ToLower(s)] = key
		}
	}
	return result
}