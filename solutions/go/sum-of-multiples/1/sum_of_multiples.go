package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	//panic("Please implement the SumMultiples function")
	set := make(map[int]struct{})
	sum := 0
	for _, d := range divisors {
        if d == 0{
            continue
        }
		for n := d; n < limit; n+=d {
			set[n] = struct{}{}
		}
	}
	for k := range set {
		sum += k
	}
	return sum
}