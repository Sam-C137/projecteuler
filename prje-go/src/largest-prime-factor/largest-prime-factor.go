package largest_prime_factor

// Run The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?
func Run(target uint64) uint64 {
	maxNumber := uint64(0)
	n := target

	for n%2 == 0 {
		maxNumber = 2
		n /= 2
	}

	for i := uint64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			maxNumber = i
			n /= i
		}
	}

	if n > 2 {
		maxNumber = n
	}

	return maxNumber
}
