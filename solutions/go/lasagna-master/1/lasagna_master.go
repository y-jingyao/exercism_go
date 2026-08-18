package lasagnamaster

func PreparationTime(layers []string, minute int) int {
	if minute == 0 {
		minute = 2
	}
	return len(layers) * minute
}

func Quantities(layers []string) (int, float64) {
	grams := 0
	liters := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			grams += 50
		} else if layer == "sauce" {
			liters += 0.2
		}
	}
	return grams, liters
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, n int) []float64 {
	scaled := make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = q * float64(n) / 2.0
	}
	return scaled
}
