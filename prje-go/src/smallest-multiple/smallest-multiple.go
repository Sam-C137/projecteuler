package smallest_multiple

import "prje-go/helpers"

// Run 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible with no remainder by all of the numbers from
// 1 to 20?
func Run(upperbound ...int) int {
	highest := helpers.DefaultArg(upperbound, 10)

	smallest := 1

	for i := 1; i <= highest; i++ {
		smallest = lcm(smallest, i)
	}

	return smallest
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// --------------------------Less Efficient Solution----------------------------------//

func run(upperbound ...int) int {
	highest := helpers.DefaultArg(upperbound, 10)

	for i := 11; i < ((1 << 31) - 1); i++ {
		isValid := true

		for j := 1; j <= highest; j++ {
			if i%j != 0 {
				isValid = false
				break
			}
		}

		if isValid {
			return i
		}
	}

	return -1
}
