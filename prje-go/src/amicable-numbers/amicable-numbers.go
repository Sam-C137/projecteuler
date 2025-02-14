package amicable_numbers

import (
	"prje-go/helpers"
)

func Run(target int) int {
	var sum int
	lookup := make(map[int]int)
	set := helpers.NewSet[int]()

	for n := range target {
		div := d(n)
		lookup[n] = div
		if lookup[div] == n {
			if div != n {
				helpers.SetAdd(set, div)
				helpers.SetAdd(set, n)
			}
		}
	}

	for n := range helpers.SetValues(set) {
		sum += n
	}

	return sum
}

func d(num int) int {
	var sum int

	for n := 1; n < num; n++ {
		if num%n == 0 {
			sum += n
		}
	}

	return sum
}
