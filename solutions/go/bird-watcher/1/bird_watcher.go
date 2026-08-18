package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	//panic("Please implement the TotalBirdCount() function")
	ans := 0
	for _, bird := range birdsPerDay {
		ans += bird
	}
	return ans
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	//panic("Please implement the BirdsInWeek() function")
	start := (week - 1) * 7
	end := start + 7
	sum := 0
	for _, bird := range birdsPerDay[start:end] {
		sum += bird
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	//panic("Please implement the FixBirdCountLog() function")
	for i:=0; i < len(birdsPerDay); i+=2{
		birdsPerDay[i]++
	}
	return birdsPerDay
}
