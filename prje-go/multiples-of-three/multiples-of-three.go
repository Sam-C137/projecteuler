package multiples_of_three

// Run If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and $9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
func Run(upperbound ...int) int {
	sum := 0
	max := 0

	if len(upperbound) < 1 {
		max = 10
	} else {
		max = upperbound[0]
	}

	for i := range max {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
