package sum_of_primes

import "prje-go/helpers"

func Run(upperbound ...int) int {
	max := helpers.DefaultArg(upperbound, 10)
	var acc int

	for i := range max {
		if isPrime(i) {
			acc += i
		}
	}

	return acc
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
