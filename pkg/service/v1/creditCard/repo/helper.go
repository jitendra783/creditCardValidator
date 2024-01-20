package repo

// At returns the digits from the start to the given length
func (d *digits) At(i int) int {
	return d[i-1]
}

//
func isInBetween(n, min, max int) bool {
	return n >= min && n <= max
}

func matchesValue(number int, numbers []int) bool {
	for _, v := range numbers {
		if v == number {
			return true
		}
	}
	return false
}