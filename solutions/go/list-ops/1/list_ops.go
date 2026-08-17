package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	//panic("Please implement the Foldl function")
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	//panic("Please implement the Foldr function")
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	//panic("Please implement the Filter function")
	filtered := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s IntList) Length() int {
	//panic("Please implement the Length function")
	res := 0
	for range s {
		res++
	}
	return res
}

func (s IntList) Map(fn func(int) int) IntList {
	//panic("Please implement the Map function")
	for i, v := range s {
		s[i] = fn(v)
	}
	return s
}

func (s IntList) Reverse() IntList {
	//panic("Please implement the Reverse function")
	for i, j := 0, s.Length()-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	//panic("Please implement the Append function")
	oldLen := s.Length()
	newLen := oldLen + lst.Length()
	result := make(IntList, newLen)
	copy(result, s)
	copy(result[oldLen:], lst)
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	//panic("Please implement the Concat function")
	for _, lst := range lists {
		s = s.Append(lst)
	}
	return s
}
