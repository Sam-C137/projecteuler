package fibonacci

// Run Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first 10 terms will be: 1, 2, 3, 5, 8, 13, 21, 34, 55, 89,
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.
func Run(fib uint64) (total uint64, even uint64) {
	if fib < 2 {
		return fib, 0
	}

	prev, curr := uint64(0), uint64(1)

	for curr < fib {
		if (curr & 1) != 1 {
			even += curr
		}

		if prev+curr < prev {
			break
		}

		prev, curr = curr, prev+curr
	}

	return curr, even
}
