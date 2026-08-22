package robotname

import "math/rand"

// Define the Robot type here.

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	//panic("Please implement the Name function")
    if r.name != "" {
        return r.name, nil
    }
	num := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var letter []rune
	lt := 'A'
	for i := 0; i < 26; i++ {
		letter = append(letter, lt)
		lt++
	}
	var name string
	name += string(letter[rand.Intn(25)])
	name += string(letter[rand.Intn(25)])
	name += string(num[rand.Intn(9)])
	name += string(num[rand.Intn(9)])
	name += string(num[rand.Intn(9)])
    r.name = name
	return string(name), nil

}

func (r *Robot) Reset() {
	//panic("Please implement the Reset function")
	r.name = ""
}
