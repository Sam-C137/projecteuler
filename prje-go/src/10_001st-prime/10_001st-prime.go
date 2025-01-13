package _0_001st_prime

import "prje-go/helpers"

// Run By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10,001st prime number?
func Run(upperbound ...int) int {
	target := helpers.DefaultArg(upperbound, 6)

	count := 0
	num := 2

	for count < target {
		if isPrime(num) {
			count++
			if count == target {
				return num
			}
		}
		if num == 2 {
			num = 3
		} else {
			num += 2
		}
	}

	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
