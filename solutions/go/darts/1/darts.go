package darts

func Score(x, y float64) int {
	distSq := x*x + y*y
	switch {
	case distSq <= 1:
		return 10
	case distSq <= 25:
		return 5
	case distSq <= 100:
		return 1
	default:
		return 0
	}
}
