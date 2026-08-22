package robotname

import (
	"errors"
	"math/rand"
	"sync"
)

type Robot struct {
	name string
}

var (
	allNames []string
	nextIdx  int
	mu       sync.Mutex
)

func init() {
	total := 26 * 26 * 1000
	allNames = make([]string, 0, total)
	for a := 'A'; a <= 'Z'; a++ {
		for b := 'A'; b <= 'Z'; b++ {
			for num := 0; num < 1000; num++ {
				name := string(a) + string(b) + sprintf3(num)
				allNames = append(allNames, name)
			}
		}
	}

	rand.Shuffle(len(allNames), func(i, j int) {
		allNames[i], allNames[j] = allNames[j], allNames[i]
	})
	nextIdx = 0
}

func sprintf3(n int) string {
	return string(rune(n/100)+'0') +
		string(rune((n%100)/10)+'0') +
		string(rune(n%10)+'0')
}

func (r *Robot) Name() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if r.name != "" {
		return r.name, nil
	}

	if nextIdx >= len(allNames) {
		return "", errors.New("")
	}

	newName := allNames[nextIdx]
	nextIdx++
	r.name = newName
	return newName, nil
}

func (r *Robot) Reset() {
	mu.Lock()
	defer mu.Unlock()
	r.name = ""
}
